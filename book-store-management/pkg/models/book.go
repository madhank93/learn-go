package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/madhank93/learn-go/book-store-management/pkg/config"
)

var db *sql.DB

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	ISBN        int    `json:"isbn"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllBooks() []Book {
	var books []Book
	var book Book
	rows, err := db.Query("SELECT name, author, publication, isbn FROM public.books;")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&book.Name, &book.Author, &book.Publication, &book.ISBN)
		books = append(books, book)
	}
	return books
}

func GetBookByID(bookID int) Book {
	var book Book
	rows, err := db.Query("SELECT name, author, publication, isbn FROM public.books WHERE isbn = $1;", bookID)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&book.Name, &book.Author, &book.Publication, &book.ISBN)
	}
	return book
}
