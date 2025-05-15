package main

import (
	"context"
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/cmd"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	exitCode := cmd.Execute(ctx)

	os.Exit(exitCode)
}
