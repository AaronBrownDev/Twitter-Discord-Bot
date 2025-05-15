package api

import (
	"github.com/bwmarrin/discordgo"
)

func (a *DiscordAPI) setChannelWithPrefix(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!set" {
		a.logger.Println("detected set channel")
		err := a.cr.SetChannel(a.ctx, m.GuildID, m.ChannelID)
		if err != nil {
			a.logger.Printf("could not set channel: %v", err)
		}
	}
}
