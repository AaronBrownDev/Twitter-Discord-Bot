package api

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (a *DiscordAPI) setChannelWithPrefix(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!set" {
		err := a.cr.SetChannel(a.ctx, m.GuildID, m.ChannelID)
		if err != nil {
			log.Printf("could not set channel: %v", err)
		}
	}
}
