package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite/discord_bot.db")

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
