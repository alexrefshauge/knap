package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func DefaultExpire() time.Time {
	return time.Now().Add(time.Hour * 24 * 365)
}

// Authenticate authenticates user by code, and returns session token if successful
func Authenticate(db *sql.DB, code string) (bool, int) {
	hash := CodeHash(code)
	var userId int
	err := db.QueryRow("SELECT id FROM users WHERE hash = ?", hash).Scan(&userId)
	if errors.Is(err, sql.ErrNoRows) {
		return false, -1
	}
	if err != nil {
		fmt.Println(err)
		return false, -1
	}
	return true, userId
}

func AuthenticateSession(db *sql.DB, token string) (bool, int) {
	var userId int
	err := db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", token).Scan(&userId)
	if errors.Is(err, sql.ErrNoRows) {
		return false, -1
	}
	if err != nil {
		fmt.Println(err)
		return false, -1
	}
	return true, userId
}
