package main

import (
	"fmt"
	"go_service/internal/interface/app/http"
	"go_service/internal/interface/container"
	"go_service/internal/shared/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cont := container.NewContainer(config.NewConfig("./resources"))

	fmt.Println("--------------------------------------------")
	fmt.Printf("API SERVICE: %s (%s) \n", cont.Config.Name, cont.Config.Version)
	fmt.Println("--------------------------------------------")

	initGracefulShutdown()
	err := http.StartHTTPServer(cont)
	if err != nil {
		panic(err)
	}
}

func initGracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Shutting down gracefully.")
		// clean up here
		os.Exit(1)
	}()
}
