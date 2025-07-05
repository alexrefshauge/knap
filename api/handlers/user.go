package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/alexrefshauge/knap/auth"
	"net/http"
)

func (ctx *Context) HandleNewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		code := r.FormValue("code")

		if name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		if code == "" {
			http.Error(w, "code is required", http.StatusBadRequest)
			return
		}

		hash := auth.CodeHash(code)
		var userCount int
		err := ctx.db.QueryRow("SELECT 1 from users where hash = ?", hash).Scan(&userCount)
		collision := err == nil && userCount > 0
		if collision {
			http.Error(w, "Invalid code", http.StatusConflict)
			return
		}

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		var userId int
		err = ctx.db.QueryRow("INSERT INTO users (name, hash) VALUES (?, ?) returning users.id", name, hash).Scan(&userId)
		if err != nil {
			http.Error(w, "Failed to register user in database", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		sessionToken, sessionExpire, err := auth.CreateSession(ctx.db, userId)
		if err != nil {
			http.Error(w, "Failed to create session", http.StatusNoContent)
			fmt.Println(err)
		}

		auth.SetSessionCookie(w, sessionToken, sessionExpire)
		w.WriteHeader(http.StatusCreated)
	}
}
