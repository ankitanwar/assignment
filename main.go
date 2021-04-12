package main

import (
	"os"
	"os/signal"

	"github.com/ankitanwar/assignment/client/application"
	server "github.com/ankitanwar/assignment/server/startServer"
)

func main() {
	go func() {
		application.StartApplication()
	}()

	go func() {
		server.StartServer()
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
