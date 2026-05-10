package app

import (
	"context"
	"errors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-project-template-v7/pkg/httputils"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	slogLogger "go-project-template-v7/pkg/logger"

	"go-project-template-v7/internal/server/middlewares"

	"go-project-template-v7/internal/api"

	"go-project-template-v7/config"
	httpserver "go-project-template-v7/internal/server/http"
	"go-project-template-v7/pkg/storage/postgres"
)

func Run(configFilePath string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// load config
	cfg := config.FromFile(configFilePath)

	// init logger
	logger := slogLogger.InitLogger(cfg.Logger.Format, cfg.Logger.Level)
	slog.SetDefault(logger)

	// init pgxpool
	pg, err := postgres.New(logger, cfg.Postgres.ConnStr, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		logger.Error("pg init error", slog.String("err", err.Error()))
		os.Exit(1)
	}
	logger.Info("pg connected")
	defer pg.Close()

	// init routing, setup API
	router := http.NewServeMux()
	mountInfraRoutes(router)
	apiHandler := api.NewHandler(
		api.NewServices(ctx, api.Deps{
			Repos: api.NewRepositories(ctx, pg),
		}))
	apiHandler.Mount(router)

	// HTTP server
	middlewareChain := middlewares.MiddlewareChain(
		// TODO: other middlewares here (oauth2, etc...)
		middlewares.LoggingMiddleware,
		middlewares.CorsMiddleware,
		// middlewares.AuthorizeMiddleware,
	)
	srv := httpserver.NewServer(middlewareChain(router))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("error occurred while running http server", slog.String("err", err.Error()))
		}
	}()

	logger.Info("server started", slog.String("port", cfg.Server.Port))

	<-ctx.Done()

	const timeout = 5 * time.Second

	shutdownCtx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(shutdownCtx); err != nil {
		logger.Error("failed to stop server", slog.String("err", err.Error()))
	} else {
		logger.Info("application stopped gracefully")
	}
}

func mountInfraRoutes(router *http.ServeMux) {
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
}
