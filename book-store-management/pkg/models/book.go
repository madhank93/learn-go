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
	createTable()
}

func createTable() {
	rows, err := db.Query(`CREATE TABLE IF NOT EXISTS public.books(id serial NOT NULL, name text NOT NULL,author text NOT NULL, publication text NOT NULL,isbn bigint NOT NULL, PRIMARY KEY (id));`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
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
		err := rows.Scan(&book.Name, &book.Author, &book.Publication, &book.ISBN)
		if err != nil {
			log.Fatalln(err)
		}
		books = append(books, book)
	}
	return books
}

func GetBookByID(bookID int) Book {
	var book Book
	rows, err := db.Query("SELECT name, author, publication, isbn FROM public.books WHERE id = $1;", bookID)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.Name, &book.Author, &book.Publication, &book.ISBN)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return book
}

func (b *Book) CreateBook() *Book {
	rows, err := db.Query("INSERT INTO public.books(name, author, publication, isbn) VALUES ($1, $2, $3, $4);", b.Name, b.Author, b.Publication, b.ISBN)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	return b
}

func (b *Book) UpdateBook(bookID int) *Book {
	rows, err := db.Query("UPDATE public.books SET name=$1, author=$2, publication=$3, isbn=$4 WHERE id = $5;", b.Name, b.Author, b.Publication, b.ISBN, bookID)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	return b
}

func DeleteBook(bookID int) string {
	rows, err := db.Query("DELETE FROM public.books WHERE id=$1;", bookID)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	return "Book has been deleted"
}
