package database

import "errors"

// TODO logic can be changed if database plans to be extended. For now it only creates a table if does not exist.
func RunMigrations() error {
	if db == nil {
		return errors.New("database not initialized")
	}

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS channel (guildID TEXT PRIMARY KEY, channelID TEXT NOT NULL, createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`)
	return err
}
