package main

import (
	"net/http"
	"log"
)

func main() {

	initDB()

	http.HandleFunc("/api/user", createHandler) 
	log.Fatal(http.ListenAndServe(":8080", nil))

	
}