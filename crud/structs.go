package main



import (
	"github.com/google/uuid"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (infp *UserInfo) makeUser(userID uuid.UUID) *User {

	return &User{
		UserID:   userID,
		Username: infp.Username,
		Password: infp.Password,
	}
}





