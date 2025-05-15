package api

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (a *DiscordAPI) setChannelWithPrefix(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!set" {
		fmt.Println("detected set channel")
		err := a.cr.SetChannel(a.ctx, m.GuildID, m.ChannelID)
		if err != nil {
			log.Printf("could not set channel: %v", err)
		}
	}
}
