package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/erickir/tinyurl/internal/config"
)

const (
	maxHeaderBytes = 1 << 20
)

// Server handles the server configuration
type Server struct {
	httpServer *http.Server
	config     *config.Config
}

func New(h http.Handler, cfg *config.Config) *Server {
	httpServer := &http.Server{
		Addr:           fmt.Sprint(cfg.Port),
		Handler:        h,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	return &Server{
		httpServer: httpServer,
		config:     cfg,
	}
}

func (s *Server) StartHTTP() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
