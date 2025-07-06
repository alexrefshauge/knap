package main

import (
	"fmt"
	"github.com/alexrefshauge/knap/handlers"
	"github.com/rs/cors"
	"net/http"
	"os"
	"strconv"
)

var (
	port        int
	environment string
	dbPath      string
)

func main() {
	var err error
	if len(os.Args) < 5 {
		panic("not enough arguments")
	}

	port, err = strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	dbPath = os.Args[2]
	environment = os.Args[3]
	certificatePath := os.Args[4]
	keyPath := os.Args[5]
	var origins []string

	switch environment {
	case "dev":
		origins = []string{"http://localhost:5173"}
	case "prod":
		origins = []string{"drknap.org"}
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
	mux.HandleFunc("POST /api/user/new", ctx.HandleNewUser())

	mux.HandleFunc("PUT /api/press", ctx.SessionAuthMiddleware(ctx.HandlePress()))

	err = http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certificatePath, keyPath, corsMiddleware.Handler(mux))
	if err != nil {
		panic(err)
	}
}
