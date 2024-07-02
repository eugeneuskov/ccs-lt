package main

import (
	"client-ccs/app"
	"client-ccs/app/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	appConfig, err := new(config.Config).Init()
	if err != nil {
		log.Fatalf("Failed initializing config: %s\n", err.Error())
		return
	}

	application := app.NewApplication(appConfig)
	application.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
