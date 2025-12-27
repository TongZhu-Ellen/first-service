package main

func createService(up *User) (int, error) {

	return createRepo(up)
}

func readService(id string) (*User, error) {
	return readRepo(id)
}


