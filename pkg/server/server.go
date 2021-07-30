package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := InitHandler()

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := &Server{server: s}

	return server, nil
}

func (s *Server) Close() error {
	log.Printf("Gracefull shutdown...")
	return nil
}

func (s *Server) Start() {
	log.Printf("Server running on http://localhost%s", s.server.Addr)
	log.Fatal(s.server.ListenAndServe())
}
