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

func createRepo(up *User) (bool, error) {
    res, err := db.Exec("INSERT IGNORE INTO users (id, password) VALUES (?, ?)", up.Id, up.Password)
    if err != nil {
        return false, err
    }
    rows, _ := res.RowsAffected()
    if rows == 0 {
        return false, nil // 已存在
    }
    return true, nil // 插入成功
}


func readRepo(id string) (*User, error) {
	var user User
    err := db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&user.Id)

    if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}


    return &user, nil

}


