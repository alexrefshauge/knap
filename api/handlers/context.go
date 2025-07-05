package handlers

import (
	"database/sql"

	"github.com/alexrefshauge/knap/database"
)

type Context struct {
	db *sql.DB
}

func NewContext(path string) *Context {
	return &Context{
		db: database.Open(path),
	}
}
