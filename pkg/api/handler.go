package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Mux struct {
	*chi.Mux
}

func NewMux() *Mux {
	return &Mux{
		chi.NewMux(),
	}
}

func (m *Mux) MountRoutes(path string, handler http.Handler) {
	m.Mount(path, handler)
}
