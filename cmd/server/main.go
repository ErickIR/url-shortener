package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	urlModule "github.com/erickir/tinyurl/internal/app/url"
	"github.com/erickir/tinyurl/internal/server"
	"github.com/erickir/tinyurl/pkg/api"
	"github.com/erickir/tinyurl/pkg/config"
	"github.com/erickir/tinyurl/pkg/mssql"
)

func startServer(ctx context.Context, server *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	serverCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		signal := <-quit
		log.Println("shutdown_signal_received: ", signal.String())
		cancel()
	}()

	server.StartHTTP(serverCtx)
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

	mux.MountHandler(urlHandlers)

	s := server.New(mux, config)

	startServer(ctx, s)

	s.Shutdown(ctx)
}
