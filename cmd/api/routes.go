package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowCredentials: true,
	}))
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/products", app.CreateProduct)
	mux.Get("/products", app.GetAllProducts)
	mux.Get("/products/{id}", app.GetProduct)
	mux.Put("/products/{id}", app.UpdateProduct)
	mux.Delete("/products/{id}", app.DeleteProduct)
	// undelete
	mux.Post("/products/{id}/undelete", app.UndeleteProduct)
	return mux
}
