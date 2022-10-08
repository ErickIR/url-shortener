package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Mux struct {
	*chi.Mux
}

type Handler interface {
	Path() string
	Routes() http.Handler
}

func NewMux() *Mux {
	return &Mux{
		chi.NewMux(),
	}
}

func (m *Mux) MountHandler(handler Handler) {
	m.Mount(handler.Path(), handler.Routes())
}
