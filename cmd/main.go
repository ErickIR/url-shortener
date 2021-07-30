package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/erickir/tinyurl/pkg/data"
	"github.com/erickir/tinyurl/pkg/server"
)

func main() {
	port := "8000"

	server, err := server.New(port)

	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}

	if err := data.GetMongoClient(); err != nil {
		log.Fatal("Error on DB: ", err.Error())
	}

	go server.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.Close()
}
