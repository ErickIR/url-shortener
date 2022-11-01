package config

import (
	"log"
	"sync"
	"time"

	"github.com/erickir/tinyurl/pkg/env"
	"github.com/joho/godotenv"
)

var (
	loadEnv = godotenv.Load

	loadEnvOnce sync.Once
)

func init() {
	if err := loadEnvConfig(); err != nil {
		log.Fatal("ERROR LOADING ENVIRONMENT VARIABLES")
	}
}

func loadEnvConfig() error {
	var err error
	loadEnvOnce.Do(func() {
		err = loadEnv(".env")
	})

	return err
}

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
			Address:  env.GetString("DATABASE_ADDRESS", ""),
			Port:     env.GetString("DATABASE_PORT", ""),
			User:     env.GetString("DATABASE_USER", ""),
			Password: env.GetString("PASSWORD", ""),
			Database: env.GetString("DATABASE_NAME", ""),
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
