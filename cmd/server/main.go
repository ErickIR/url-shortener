package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	urlModule "github.com/erickir/tinyurl/internal/app/url"
	"github.com/erickir/tinyurl/internal/server"
	"github.com/erickir/tinyurl/pkg/api"
	"github.com/erickir/tinyurl/pkg/config"
	"github.com/erickir/tinyurl/pkg/mssql"
)

const (
	shutdownTimeout = 15 * time.Second
)

func startServerWithGracefullShutdown(ctx context.Context, server *server.Server) error {
	go func() {
		if err := server.StartHTTP(); err != nil {
			log.Fatal("ERROR RUNNING SERVER: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	signal.Notify(quit, os.Kill, syscall.SIGTERM)

	<-quit

	shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	return server.Shutdown(shutdownCtx)
}

func main() {
	ctx := context.Background()
	mux := api.NewMux()

	config := config.New()

	sqlConfig := config.SqlServer

	db, err := mssql.NewClient(ctx, sqlConfig.Address, sqlConfig.User, sqlConfig.Password, sqlConfig.Port, sqlConfig.Database)
	if err != nil {
		log.Fatal("ERROR CONNECTING TO DATABASE: ", err.Error())
	}

	urlHandlers := urlModule.Setup(db)

	mux.MountRoutes("/url", urlHandlers.Routes())

	s := server.New(mux, config)

	err = startServerWithGracefullShutdown(ctx, s)
	if err != nil {
		log.Fatal("ERROR SHUTTING DOWN THE SERVER: ", err.Error())
	}
}
