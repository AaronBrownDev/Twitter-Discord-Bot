package api

import "github.com/bwmarrin/discordgo"

func (a *DiscordAPI) pingpong(s *discordgo.Session, m *discordgo.MessageCreate) {

	// if message is sent by the bot then return immediately
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	} else if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

}
