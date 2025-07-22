package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexrefshauge/knap/auth"
	"github.com/alexrefshauge/knap/handlers"
	"github.com/rs/cors"
)

var (
	port        int
	environment string
	dbPath      string
)

const (
	PORT = iota + 1
	DB_PATH
	ENV 
	CERT_PATH
	KEY_PATH

)

func main() {
	var err error
	fmt.Println(os.Args[ENV])
	if len(os.Args) < 5 && os.Args[ENV] != "dev" {
		panic("not enough arguments")
	}

	port, err = strconv.Atoi(os.Args[PORT])
	if err != nil {
		panic(err)
	}

	dbPath = os.Args[DB_PATH]
	environment = os.Args[ENV]
	certificatePath := os.Args[CERT_PATH]
	keyPath := os.Args[KEY_PATH]
	var origins []string

	switch environment {
	case "dev":
		origins = []string{"http://localhost:5173"}
		auth.Origin = "http://localhost:5173"
	case "prod":
		origins = []string{"https://drknap.org", "https://api.drknap.org"}
	}

	var allowFunc func(string) bool
	if environment == "dev" {
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
	ctx := handlers.NewContext(dbPath)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", ctx.HandleHealth())

	mux.HandleFunc("POST /api/user/auth", ctx.HandleAuthenticate())
	mux.HandleFunc("GET /api/user/auth", ctx.HandleAuthenticate())
	mux.HandleFunc("POST /api/user/new", ctx.HandleNewUser())

	mux.HandleFunc("PUT /api/press", ctx.SessionAuthMiddleware(ctx.HandlePress()))
	mux.HandleFunc("DELETE /api/press", ctx.SessionAuthMiddleware(ctx.HandlePressUndo()))

	mux.HandleFunc("GET /api/press/today", ctx.SessionAuthMiddleware(ctx.HandlePressGetToday()))
	//mux.HandleFunc("GET /api/press/week", nil)

	fmt.Printf("[%s] Listening on port %d\n", environment, port)
	fmt.Println(time.Now().String())
	if environment == "dev" {
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), corsMiddleware.Handler(mux))
	} else {
		err = http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certificatePath, keyPath, corsMiddleware.Handler(mux))
	}
	if err != nil {
		panic(err)
	}
}
