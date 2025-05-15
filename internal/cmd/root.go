package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Execute() {

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

}
