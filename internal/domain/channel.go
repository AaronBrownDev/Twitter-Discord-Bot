package domain

import "time"

type Channel struct {
	guildID    string
	channelID  string
	assignedAt time.Time
}

type ChannelRepository interface {
	GetAll() ([]Channel, error)
	SetChannel(guildID, channelID string) error
}
