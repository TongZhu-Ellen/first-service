package main


import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `gorm:"primaryKey"`
	Username string
	Password string
}
