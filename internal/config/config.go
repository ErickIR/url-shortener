package config

import (
	"time"
)

var (
	// Server
	port           = ":8080"
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	defaultTimeout = 10 * time.Second
	debug          = false

	// SqlServer
	sqlServerPort = "1433"
	address       = "localhost"
	user          = "sa"
	password      = "p@ssw0rd"
	database      = "TinyUrlDatabase"
)

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

func New() *Config {
	return &Config{
		ServerConfig: &Server{
			Port:           port,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			DefaultTimeout: defaultTimeout,
			Debug:          debug,
		},
		SqlServer: &Database{
			Address:  address,
			Port:     sqlServerPort,
			User:     user,
			Password: password,
			Database: database,
		},
	}
}
