package main

func createService(up *User) (bool, error) {

	return createRepo(up)
}

func readService(id string) (*User, error) {
	return readRepo(id)
}


