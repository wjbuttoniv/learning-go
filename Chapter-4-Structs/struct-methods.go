package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// Create a method that returns a string
// in the format of
// USERNAME:PASSWORD
func (authInfo authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
}

// dont touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("=================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
