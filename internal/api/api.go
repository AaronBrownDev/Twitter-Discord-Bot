package api

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
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
