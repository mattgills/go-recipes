package db

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

func Init() {
	var err error

	db, err = sql.Open("sqlite", "./db/recipes.db")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS recipes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)`,
	)

	if err != nil {
		log.Fatal(err)
	}

}

func GetDb() *sql.DB {
	return db
}
