package cmd

import (
	//"Twitter-Discord-Bot/internal/api"
	"Twitter-Discord-Bot/internal/database"
	"log"
)

func APICmd() error {
	if err := database.InitializeDB(); err != nil {
		return err
	}
	defer database.CloseDB()

	//db := database.GetDB()

	//discordAPI := api.NewDiscordAPI(db)

	log.Println("Starting Discord API")

	return nil
}
