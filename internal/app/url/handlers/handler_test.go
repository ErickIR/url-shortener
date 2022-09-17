package handlers

import (
	"testing"

	"github.com/erickir/tinyurl/internal/app/url/domain"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	c := require.New(t)

	service := domain.URLService{}

	handler := New(service)
	c.NotNil(handler)
}

func TestRoutes(t *testing.T) {
	c := require.New(t)

	service := domain.URLService{}

	handler := New(service)
	c.NotNil(handler)

	routes := handler.Routes()
	c.NotNil(routes)
}

func TestGetLongUrlHandler(t *testing.T) {
	c := require.New(t)

	service := domain.URLService{}

	handler := New(service)
	c.NotNil(handler)
}
