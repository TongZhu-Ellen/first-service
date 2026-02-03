package main

type User struct {
	UserID   string `gorm:"primaryKey"`
	Username string
	Password string
}
