package auth

import (
	"database/sql"
	"errors"
	"github.com/alexrefshauge/knap/config"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var Origin = ".drknap.org"
var SameSiteMode http.SameSite = http.SameSiteNoneMode

func init() {
	if config.Environment == "dev" {
		SameSiteMode = http.SameSiteLaxMode
		Origin = "localhost"
	}
}

func CreateSession(db *sql.DB, userId int) (string, time.Time, error) {
	sessionId, err := uuid.NewUUID()
	if err != nil {
		return "", time.Now(), err
	}
	token := sessionId.String()

	expireDate := time.Now().Add(time.Hour * 24 * 360)

	result, err := db.Exec("INSERT INTO sessions (user_id, token, expires_at) VALUES (?, ?, ?)", userId, token, expireDate)
	if err != nil {
		return "", time.Now(), err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return "", time.Now(), err
	}

	if rows == 0 {
		return "", time.Now(), errors.New("failed to create session")
	}

	return token, expireDate, nil
}

func SetSessionCookie(w http.ResponseWriter, token string, expireDate time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Expires:  expireDate,
		Path:     "/",
		Domain:   Origin,
		HttpOnly: config.Environment != "dev",
		Secure:   config.Environment != "dev",
		SameSite: SameSiteMode,
	})
}
