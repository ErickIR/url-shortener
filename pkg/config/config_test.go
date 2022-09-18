package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	c := require.New(t)

	cfg := New()

	c.Equal(":8080", cfg.Port)
}
