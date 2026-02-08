package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/swaggo/http-swagger"
)

func initChi() {
	r = chi.NewRouter()

	// User CRUD
	r.Post("/user", Create)
	r.Get("/user/{id}", Read)
	r.Put("/user/{id}", Update)
	r.Delete("/user/{id}", Delete) 

	// swagger!
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
