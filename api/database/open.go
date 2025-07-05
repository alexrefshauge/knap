package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		panic(err)
	}

	initializeDatabase(db)

	return db
}

func initializeDatabase(db *sql.DB) {
	db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		hash TEXT NOT NULL UNIQUE
	);
	
	CREATE TABLE IF NOT EXISTS sessions (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    user_id INTEGER NOT NULL REFERENCES users(id),
	    token TEXT NOT NULL,
	    expires_at DATETIME NOT NULL,
	    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS button_pushes (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    user_id INTEGER NOT NULL REFERENCES users(id),
	    pushed_at DATETIME default current_timestamp,
	    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`)
}
