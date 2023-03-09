package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string
	LastName  string
}

var movies []Movie

func main() {
	r := chi.NewRouter()

	movies = append(movies, Movie{ID: "001", Title: "Fight club", Director: &Director{FirstName: "David", LastName: "Fincher"}})
	movies = append(movies, Movie{ID: "002", Title: "Gone Girl", Director: &Director{FirstName: "David", LastName: "Fincher"}})

	r.Get("/movies", getMovies)

	fmt.Printf("Starting server on 8080 \n")
	http.ListenAndServe(":8080", r)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
