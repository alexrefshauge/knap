package auth

import (
	"database/sql"
	"errors"
)

// Authenticate authenticates user by code, and returns session token if successful
func Authenticate(db *sql.DB, code string) (bool, string) {
	hash := CodeHash(code)
	err := db.QueryRow("SELECT id FROM users WHERE hash = ?", hash).Scan()
	if errors.Is(err, sql.ErrNoRows) {
		return false, ""
	}
	return false, ""
}

func AuthenticateSession(db *sql.DB, token string) bool {
	db.QueryRow("SELECT id FROM sessions WHERE token = ?", token).Scan()
	return false
}
