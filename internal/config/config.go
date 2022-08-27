package config

import (
	"time"
)

var (
	port           = ":8080"
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	defaultTimeout = 10 * time.Second
	debug          = false
)

type Config struct {
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	DefaultTimeout time.Duration
	Debug          bool
}

func New() *Config {
	return &Config{
		Port:           port,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		DefaultTimeout: defaultTimeout,
		Debug:          debug,
	}
}
