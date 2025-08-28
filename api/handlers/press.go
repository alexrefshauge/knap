package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexrefshauge/knap/model"
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
		err := ctx.db.QueryRow("SELECT id FROM button_pushes WHERE user_id = ? ORDER BY pushed_at DESC LIMIT 1", userId).Scan(&lastId)
		if err != nil {
			http.Error(w, "Failed to find last button press", http.StatusInternalServerError)
			return
		}
		_, err = ctx.db.Exec("DELETE FROM button_pushes WHERE id = ?", lastId)
		if err != nil {
			http.Error(w, "Failed to find last button press", http.StatusInternalServerError)
			return
		}
	}
}

const dateLayout = "2-1-2006"

type countResponse struct {
	Count int `json:"count"`
}

func (ctx *Context) HandlePressGetToday() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		dateParam := query.Get("date")
		countParam := query.Get("count")
		if dateParam == "" {
			http.Error(w, "Missing parameter: date", http.StatusBadRequest)
			return
		}
		if countParam == "" {
			http.Error(w, "Missing parameter: count", http.StatusBadRequest)
			return
		}

		date, err := time.Parse(dateLayout, dateParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid parameter value for date: %s", dateParam), http.StatusBadRequest)
			fmt.Println(err.Error())
			return
		}
		dateDayStart := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
		dateDayEnd := dateDayStart.Add(24 * time.Hour)

		var presses []time.Time
		rows, err := ctx.db.Query("SELECT pushed_at FROM button_pushes WHERE pushed_at > ? AND pushed_at < ?", dateDayStart, dateDayEnd)
		if err != nil {
			http.Error(w, "Failed to find button presses", http.StatusInternalServerError)
		}
		for rows.Next() {
			var t time.Time
			rows.Scan(&t)
			presses = append(presses, t)
		}

		count := len(presses)
		if countParam == "t" {
			responseBytes, err := json.Marshal(countResponse{Count: count})
			if err != nil {
				http.Error(w, "Failed to format count response", http.StatusInternalServerError)
				fmt.Println(err.Error())
				return
			}

			_, err = w.Write(responseBytes)
			if err != nil {
				http.Error(w, "Failed to write respone", http.StatusInternalServerError)
				return
			}
			return
		}

		responseBytes, err := json.Marshal(presses)
		if err != nil {
			http.Error(w, "Failed to count presses", http.StatusInternalServerError)
			return
		}

		w.Write(responseBytes)
	}
}

func (ctx *Context) HandlePressGetWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		dateParam := query.Get("date")
		if dateParam == "" {
			http.Error(w, "Missing parameter: date", http.StatusBadRequest)
			return
		}
		date, err := time.Parse(dateLayout, dateParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid parameter value for date: %s", dateParam), http.StatusBadRequest)
			return
		}
		weekday := int(date.Weekday() + 1)
		dateStart := time.Date(date.Year(), date.Month(), date.Day()-weekday, 0, 0, 0, 0, date.Location())
		dateEnd := dateStart.Add(24 * time.Hour * 7)
		var presses []time.Time
		rows, err := ctx.db.Query("SELECT pushed_at FROM button_pushes WHERE pushed_at > ? AND pushed_at < ?", dateStart, dateEnd)
		if err != nil {
			http.Error(w, "Failed to find button presses", http.StatusInternalServerError)
		}
		for rows.Next() {
			var t time.Time
			rows.Scan(&t)
			presses = append(presses, t)
		}

		pressesGrouped := model.GroupByWeekday(presses)

		responseBytes, err := json.Marshal(pressesGrouped)
		if err != nil {
			http.Error(w, "Failed to count presses", http.StatusInternalServerError)
			return
		}

		w.Write(responseBytes)
	}
}
