package app

import (
	"context"
	"errors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-project-template-v5/pkg/httputils"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	slogLogger "go-project-template-v5/pkg/logger"

	"go-project-template-v5/internal/server/middlewares"

	"go-project-template-v5/internal/api"

	"go-project-template-v5/config"
	httpserver "go-project-template-v5/internal/server/http"
	"go-project-template-v5/pkg/storage/postgres"
)

func Run(configFilePath string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// load config
	config.LoadConfigFromFile(configFilePath)
	cfg := config.Cfg()

	// init logger
	logger := slogLogger.InitLogger(cfg.Logger.Format, cfg.Logger.Level)
	slog.SetDefault(logger)

	// init pgxpool
	pg, err := postgres.New(logger, cfg.Postgres.URL, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		logger.Error("pg init error", slog.String("err", err.Error()))
	} else {
		logger.Info("pg connected")
	}
	defer pg.Close()

	// routes for all module (docs, health-checks, etc...)
	defaultRoutes := defaultRoutes(ctx)
	repositories := api.NewRepositories(ctx, pg)
	services := api.NewServices(ctx, api.Deps{
		Repos: repositories,
	})
	handler := api.NewHandler(services)
	handler.Mount(defaultRoutes)

	// HTTP server
	middlewareChain := middlewares.MiddlewareChain(
		// TODO: other middlewares here (oauth2, etc...)
		middlewares.LoggingMiddleware,
		// middlewares.AuthorizeMiddleware,
	)
	srv := httpserver.NewServer(middlewareChain(defaultRoutes))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("error occurred while running http server", slog.String("err", err.Error()))
		}
	}()

	logger.Info("server started", slog.String("port", cfg.Server.Port))

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Error("failed to stop server", slog.String("err", err.Error()))
	} else {
		logger.Info("application stopped gracefully")
	}
}

func defaultRoutes(_ context.Context) *http.ServeMux {
	router := http.NewServeMux()

	// docs

	router.Handle("/swagger-ui/", httpSwagger.WrapHandler)
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	// internal

	router.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		httputils.WriteJSON(w, http.StatusOK, map[string]string{
			"status": "UP",
		})
	})

	return router
}
