package main



import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (infp *UserInfo) makeUser(userID uuid.UUID) *User {
    hash, err := bcrypt.GenerateFromPassword([]byte(infp.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    return &User{
        UserID:   userID,
        Username: infp.Username,
        Password: string(hash), // 原地变 hash
    }
}






