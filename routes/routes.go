package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shawon325/go-crud/src/controllers"
)

func Routes() http.Handler {
	routes := chi.NewRouter()

	routes.Get("/", func(response http.ResponseWriter, request *http.Request) {
		_, err := response.Write([]byte("Welcome to Go CRUD API, Shawon"))
		if err != nil {
			return
		}
	})

	routes.Get("/users", controllers.Get)
	routes.Post("/users", controllers.Create)
	routes.Get("/users/{id}", controllers.Show)
	routes.Put("/users/{id}", controllers.Update)
	routes.Delete("/users/{id}", controllers.Delete)

	return routes
}
