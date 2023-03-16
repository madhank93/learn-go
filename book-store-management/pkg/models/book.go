package models

import (
	"database/sql"

	"github.com/madhank93/learn-go/book-store-management/pkg/config"
)

var db *sql.DB

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}
