package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

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

	r.Get("/movies", listMovies)
	r.Get("/movie/{id}", getMovie)
	r.Post("/movie", createMovie)
	r.Put("/movie/{id}", updateMovie)
	r.Delete("/movie/{id}", deleteMovie)

	fmt.Printf("Starting server on 8080 \n")
	http.ListenAndServe(":8080", r)
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	for _, movie := range movies {
		if movie.ID == chi.URLParam(r, "id") {
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	for index, item := range movies {
		if item.ID == chi.URLParam(r, "id") {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = chi.URLParam(r, "id")
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	for index, movie := range movies {
		if movie.ID == chi.URLParam(r, "id") {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
