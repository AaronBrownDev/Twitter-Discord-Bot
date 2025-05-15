package domain

import (
	"context"
	"time"
)

type Channel struct {
	GuildID    string
	ChannelID  string
	AssignedAt time.Time
}

type ChannelRepository interface {
	GetAll(ctx context.Context) ([]Channel, error)
	SetChannel(ctx context.Context, guildID, channelID string) error
}
