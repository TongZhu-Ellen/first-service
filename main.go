package main

import (
	"net/http"
	"log"
	"github.com/go-chi/chi"
)

func main() {

	initDB()

	r := chi.NewRouter()

	r.Post("/api/user", createHandler)
	r.Get("/api/user/{id}", readHandler)
	r.Put("/api/user/{id}", updateHandler)



	log.Fatal(http.ListenAndServe(":8080", r))

	
}