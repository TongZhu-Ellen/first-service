package main

func createService(up *User) (int64, error) {

	return createRepo(up)
}

func readService(id string) (*User, error) {
	return readRepo(id)
}

func updateService(id string, reqp *UpdateUserRequest) (int64, error) {
	return updateRepo(id, reqp)
}


