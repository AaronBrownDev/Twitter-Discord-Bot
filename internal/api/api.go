package api

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/domain"
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/repository"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

type DiscordAPI struct {
	ctx    context.Context
	cr     domain.ChannelRepository
	dg     *discordgo.Session
	db     *sql.DB
	logger *log.Logger
}

func NewDiscordAPI(ctx context.Context, db *sql.DB) *DiscordAPI {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	return &DiscordAPI{
		ctx:    ctx,
		cr:     repository.NewSqliteChannelRepository(db),
		dg:     dg,
		db:     db,
		logger: log.New(os.Stdout, "discord: ", log.LstdFlags),
	}
}

func (a *DiscordAPI) Start() error {
	a.addHandlers()

	a.dg.Identify.Intents = discordgo.IntentsGuildMessages

	err := a.dg.Open()
	if err != nil {
		return fmt.Errorf("could not open Discord session: %v", err)
	}

	// Waits for context cancellation (ctrl+c)
	<-a.ctx.Done()

	a.logger.Println("Shutting down Discord connection...")
	err = a.dg.Close()
	if err != nil {
		return fmt.Errorf("could not close Discord session: %v", err)
	}

	return nil
}

// addHandlers is a helper function that adds all the handlers to the Discord session.
func (a *DiscordAPI) addHandlers() {
	a.dg.AddHandler(a.pingpong)
	a.dg.AddHandler(a.setChannelWithPrefix)
}
