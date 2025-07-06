package handlers

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Database string `json:"database"`
}

func (ctx *Context) HandleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ctx.db.Ping()
		dbStatus := "ok"
		if err != nil {
			dbStatus = err.Error()
		}

		response := healthResponse{
			Database: dbStatus,
		}

		responseData, err := json.Marshal(response)
		_, err = w.Write(responseData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}
