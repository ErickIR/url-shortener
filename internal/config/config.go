package config

import (
	"time"

	"github.com/erickir/tinyurl/pkg/env"
)

func New() *Config {
	return &Config{
		ServerConfig: &Server{
			Port:           env.GetString("PORT", ":8080"),
			ReadTimeout:    env.GetDuration("READ_TIMEOUT", 10*time.Second),
			WriteTimeout:   env.GetDuration("WRITE_TIMEOUT", 10*time.Second),
			DefaultTimeout: env.GetDuration("DEFAULT_TIMEOUT", 10*time.Second),
			Debug:          env.GetBool("DEBUG", false),
		},
		SqlServer: &Database{
			Address:  env.GetString("DATABASE_ADDRESS", "localhost"),
			Port:     env.GetString("SQL_SERVER_PORT", "1433"),
			User:     env.GetString("DATABASE_USER", "sa"),
			Password: env.GetString("PASSWORD", ""),
			Database: env.GetString("DATABASE_NAME", "TINY_URL_DATABASE"),
		},
	}
}

type Server struct {
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	DefaultTimeout time.Duration
	Debug          bool
}

type Database struct {
	Address  string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	SqlServer    *Database
	ServerConfig *Server
}
