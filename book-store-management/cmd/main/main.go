package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/madhank93/learn-go/book-store-management/pkg/routes"
)

func main() {
	r := chi.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
