package models

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func SetDB(dbConnect *sqlx.DB) { db = dbConnect }

func MakeMigrations() {
	// Db migrations there
}
