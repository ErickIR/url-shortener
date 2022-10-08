package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/erickir/tinyurl/pkg/config"
)

const (
	maxHeaderBytes  = 1 << 20
	shutdownTimeout = 15 * time.Second
)

// Server handles the server configuration
type Server struct {
	httpServer *http.Server
	config     *config.Config
}

func New(h http.Handler, cfg *config.Config) *Server {
	httpServer := &http.Server{
		Addr:           fmt.Sprint(cfg.ServerConfig.Port),
		Handler:        h,
		ReadTimeout:    cfg.ServerConfig.ReadTimeout,
		WriteTimeout:   cfg.ServerConfig.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	return &Server{
		httpServer: httpServer,
		config:     cfg,
	}
}

func (s *Server) StartHTTP(ctx context.Context) {
	go s.startHttpServer()

	<-ctx.Done()
	log.Println("server_stopped")
}

func (s *Server) startHttpServer() {
	log.Println("server_started")
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Println("server_error")
		panic(err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("server_shutdown")

	shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()
	return s.httpServer.Shutdown(shutdownCtx)
}
