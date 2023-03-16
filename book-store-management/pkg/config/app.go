package config

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Connect() {
	connStr := "postgres://postgres:postgres@localhost/book-store?sslmode=require"
	d, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
