package cmd

import (
	"context"
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/api"
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/database"
	"log"
)

func APICmd(ctx context.Context) error {
	if err := database.InitializeDB(); err != nil {
		return err
	}
	defer database.CloseDB()

	if err := database.RunMigrations(); err != nil {
		return err
	}

	db := database.GetDB()

	discordAPI := api.NewDiscordAPI(ctx, db)

	log.Println("Starting Discord API")

	err := discordAPI.Start()
	if err != nil {
		return err
	}

	return nil
}
