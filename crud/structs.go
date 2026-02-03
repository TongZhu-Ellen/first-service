package main

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (infp *UserInfo) makeUser(userID string) *User {

	return &User{
		UserID:   userID,
		Username: infp.Username,
		Password: infp.Password,
	}
}





