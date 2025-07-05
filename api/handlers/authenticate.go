package handlers

import "net/http"

func (ctx *Context) HandleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (ctx *Context) HandleAuthenticateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}
