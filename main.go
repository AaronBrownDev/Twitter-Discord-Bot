package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// TODO CONVERT THIS INTO PROJECT STRUCTURE

var (
	discordToken string
)

func init() {
	discordToken = os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		discordToken = "INSERT_TOKEN_HERE"
	}
}

func main() {
	session, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal(err)
	}

	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	fmt.Println("Bot is now running")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
