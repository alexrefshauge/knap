package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/alexrefshauge/knap/auth"
	"net/http"
)

func (ctx *Context) HandleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			c, err := r.Cookie("session")
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(w, "No session cookie", http.StatusUnauthorized)
				return
			}
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if valid, _ := auth.AuthenticateSession(ctx.db, c.Value); !valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			return
		}

		code := r.FormValue("code")
		if code == "" {
			http.Error(w, "code is required", http.StatusBadRequest)
			return
		}

		var authentic bool
		var userId int
		if authentic, userId = auth.Authenticate(ctx.db, code); !authentic {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionToken, sessionExpire, err := auth.CreateSession(ctx.db, userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		auth.SetSessionCookie(w, sessionToken, sessionExpire)
		w.WriteHeader(http.StatusCreated)
	}
}

func (ctx *Context) SessionAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session")
		if errors.Is(err, http.ErrNoCookie) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Unable to read session cookie", http.StatusInternalServerError)
		}

		sessionToken := c.Value
		var valid bool
		var userId int
		if valid, userId = auth.AuthenticateSession(ctx.db, sessionToken); !valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rWithUserId := r.WithContext(context.WithValue(r.Context(), "user", userId))
		next.ServeHTTP(w, rWithUserId)
	}
}
