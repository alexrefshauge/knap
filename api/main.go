package main

import (
	"fmt"
	"github.com/alexrefshauge/knap/config"
	_ "github.com/alexrefshauge/knap/config"
	"net/http"
	"time"

	"github.com/alexrefshauge/knap/auth"
	"github.com/alexrefshauge/knap/handlers"
	"github.com/rs/cors"
)

func main() {
	var err error
	var origins []string

	switch config.Environment {
	case "dev":
		origins = []string{"http://localhost:5173"}
		auth.Origin = "http://localhost:5173"
	case "prod":
		origins = []string{"https://drknap.org", "https://api.drknap.org"}
	}

	var allowFunc func(string) bool
	if config.Environment == "dev" {
		allowFunc = func(path string) bool {
			return true
		}
	}

	corsMiddleware := cors.New(cors.Options{
		AllowOriginFunc:  allowFunc,
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowCredentials: true,
	})
	ctx := handlers.NewContext(config.DbPath)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", ctx.HandleHealth())

	mux.HandleFunc("POST /api/user/auth", ctx.HandleAuthenticate())
	mux.HandleFunc("GET /api/user/auth", ctx.HandleAuthenticate())
	mux.HandleFunc("POST /api/user/new", ctx.HandleNewUser())
	mux.HandleFunc("DELETE /api/user/invalidate", ctx.SessionAuthMiddleware(ctx.HandleInvalidate()))

	mux.HandleFunc("PUT /api/press", ctx.SessionAuthMiddleware(ctx.HandlePress()))
	mux.HandleFunc("DELETE /api/press", ctx.SessionAuthMiddleware(ctx.HandlePressUndo()))

	mux.HandleFunc("GET /api/press/today", ctx.SessionAuthMiddleware(ctx.HandlePressGetToday()))
	mux.HandleFunc("GET /api/press/week", ctx.SessionAuthMiddleware(ctx.HandlePressGetWeek()))
	//mux.HandleFunc("GET /api/press/week", nil)

	fmt.Printf("[%s] Listening on port %d\n", config.Environment, config.Port)
	fmt.Println(time.Now().String())
	if config.Environment == "dev" {
		err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), corsMiddleware.Handler(mux))
	} else {
		err = http.ListenAndServeTLS(fmt.Sprintf(":%d", config.Port), config.CertificatePath, config.KeyPath, corsMiddleware.Handler(mux))
	}
	if err != nil {
		panic(err)
	}
}
