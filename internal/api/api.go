package api

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type DiscordAPI struct {
	dg *discordgo.Session
	db *sql.DB
}

func NewDiscordAPI(db *sql.DB) *DiscordAPI {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	return &DiscordAPI{
		dg: dg,
		db: db,
	}
}

func (a *DiscordAPI) Start() error {
	a.addHandlers()

	a.dg.Identify.Intents = discordgo.IntentsGuildMessages

	err := a.dg.Open()
	if err != nil {
		return fmt.Errorf("could not open Discord session: %v", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return a.dg.Close()
}

// addHandlers is a helper function that adds all the handlers to the Discord session.
func (a *DiscordAPI) addHandlers() {
	a.dg.AddHandler(a.pingpong)
}
