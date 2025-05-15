package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/Twitter-Discord-Bot/internal/domain"
)

type sqliteChannelRepository struct {
	db *sql.DB
}

func NewSqliteChannelRepository(db *sql.DB) domain.ChannelRepository {
	return &sqliteChannelRepository{db: db}
}

// fetch is a helper function for when a sql query is expected to return multiple rows of channels
func (r *sqliteChannelRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.Channel, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []domain.Channel
	for rows.Next() {
		var channel domain.Channel
		err = rows.Scan(
			&channel.GuildID,
			&channel.ChannelID,
			&channel.AssignedAt,
		)
		if err != nil {
			return nil, err
		}

		channels = append(channels, channel)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return channels, nil
}

func (r *sqliteChannelRepository) GetAll(ctx context.Context) ([]domain.Channel, error) {
	query := `SELECT guildID, channelID, assignedAt FROM channel`

	return r.fetch(ctx, query)
}

func (r *sqliteChannelRepository) SetChannel(ctx context.Context, guildID, channelID string) error {
	query := `INSERT OR REPLACE INTO channel (guildID, channelID, assignedAt) VALUES (?, ?, CURRENT_TIMESTAMP)`

	result, err := r.db.ExecContext(ctx, query, guildID, channelID)
	if err != nil {
		return fmt.Errorf("could not set channel: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}
