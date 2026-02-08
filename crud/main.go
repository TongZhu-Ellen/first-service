package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	
    _ "crud-service/docs"  // swag init 生成的 docs 包
)

var (
	db *gorm.DB
	r  *chi.Mux
)

func main() {
	initDB()
	initChi()
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}