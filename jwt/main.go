package main

import (
	"fmt"
)

func main() {

	test()

}

func test() {

	testID := "test_id"

	token, err := IssueToken(testID)
	if err != nil {
		fmt.Println("Test not passed!")
	}

	parsedID, ok := ParseID(token)
	if !ok {
		fmt.Println("Test not passed!")
	}

	fmt.Println("testID:", testID)
	fmt.Println("token:", token)
	fmt.Println("parsedID:", parsedID)

}
