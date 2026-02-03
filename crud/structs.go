package main

type UserCreation struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type UserUpdating struct {
	
	Password    string `json:"password"`
	NewUsername string `json:"new_username"`
	NewPassword string `json:"new_password"`
}

type UserDeletion struct {
	
	Password string `json:"password"`
}
