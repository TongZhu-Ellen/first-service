package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:888888@tcp(localhost:3306)/first_db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func createRepo(up *User) error {

	_, err := db.Exec(
		"INSERT INTO users (id, password) VALUES (?, ?)",
		up.Id, up.Password,
	)

	return err

}
