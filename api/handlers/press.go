package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (ctx *Context) HandlePress() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value("user").(int)
		if !ok {
			http.Error(w, "Missing user id", http.StatusInternalServerError)
			return
		}

		var pressId int
		err := ctx.db.QueryRow("INSERT INTO button_pushes (user_id, pushed_at) VALUES (?, ?) RETURNING id",
			userId, time.Now()).Scan(&pressId)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No button press was registered", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, "Failed to register button press", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Button press registered id:%d\n", pressId)
	}
}

func (ctx *Context) HandlePressUndo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value("user").(int)
		if !ok {
			http.Error(w, "Missing user id", http.StatusInternalServerError)
			return
		}
		
		var lastId int
		ctx.db.QueryRow("SELECT id FROM button_pushes WHERE user_id = ? ORDER BY pushed_at DESC LIMIT 1", userId).Scan(&lastId)
	}
}
