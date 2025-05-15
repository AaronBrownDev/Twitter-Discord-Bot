package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InitializeDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS channel (guildID TEXT PRIMARY KEY, channelID TEXT NOT NULL, createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`)
	return err
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	if db == nil {
		return fmt.Errorf("no existing database connection")
	}
	err := db.Close()
	if err != nil {
		return fmt.Errorf("err closing database connection: %v", err)
	}
	return nil
}
