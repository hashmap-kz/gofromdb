package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"go-project-template-v5/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server {
	cfg := config.Cfg()

	return &Server{
		httpServer: &http.Server{
			// TODO: other server properties here
			Addr:         ":" + cfg.Server.Port,
			Handler:      handler,
			ReadTimeout:  mustParseDuration(cfg.Server.ReadTimeout),
			WriteTimeout: mustParseDuration(cfg.Server.WriteTimeout),
		},
	}
}

func mustParseDuration(d string) time.Duration {
	parsed, err := time.ParseDuration(d)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
