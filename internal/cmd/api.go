package cmd

import (
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/api"
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/database"
	"log"
)

func APICmd() error {
	if err := database.InitializeDB(); err != nil {
		return err
	}
	defer database.CloseDB()

	db := database.GetDB()

	discordAPI := api.NewDiscordAPI(db)

	log.Println("Starting Discord API")

	err := discordAPI.Start()
	if err != nil {
		return err
	}

	return nil
}
