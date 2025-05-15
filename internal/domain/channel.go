package domain

import "time"

type Channel struct {
	GuildID    string
	ChannelID  string
	AssignedAt time.Time
}

type ChannelRepository interface {
	GetAll() ([]Channel, error)
	SetChannel(guildID, channelID string) error
}
