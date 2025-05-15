package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Execute(ctx context.Context) int {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	go func() {
		if err := APICmd(); err != nil {
			log.Printf("Error starting API: %v", err)
			quit <- syscall.SIGTERM
		}
	}()

	<-quit
	log.Println("Shutting down...")

	return 0
}
