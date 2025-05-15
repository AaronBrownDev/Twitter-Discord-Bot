package api

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (a *DiscordAPI) pingpong(s *discordgo.Session, m *discordgo.MessageCreate) {

	// if message is sent by the bot then return immediately
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Println(m.Content)

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	} else if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	} else if m.Content == "!viewAll" {
		a.logger.Println("detected")
		var prettyChannels string
		channels, err := a.cr.GetAll(a.ctx)
		if err != nil {
			log.Println(err)
		}
		a.logger.Println("channels:", channels)
		for _, channel := range channels {
			prettyChannels += fmt.Sprintf("ChannelID: %s\n", channel.ChannelID)
		}
		s.ChannelMessageSend(m.ChannelID, prettyChannels)
	}

}
