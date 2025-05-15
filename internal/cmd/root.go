package cmd

import (
	"context"
	"log"
)

func Execute(ctx context.Context) int {

	if err := APICmd(ctx); err != nil {
		log.Printf("Error starting API: %v", err)
		return 1
	}

	log.Println("Shutting down...")

	return 0
}
