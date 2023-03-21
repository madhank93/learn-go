package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/madhank93/learn-go/book-store-management/pkg/controller"
)

var RegisterBookStoreRoutes = func(router *chi.Mux) {
	router.Get("/book", controller.GetBooks)
	router.Get("/book/{id}", controller.GetBooksByID)
	router.Post("/book", controller.CreateBook)
	router.Put("/book/{id}", controller.UpdateBook)
	router.Delete("/book/{id}", controller.DeleteBook)
}
