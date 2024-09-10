package main

import "fmt"

// only return the customers firstName
func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio", firstName)
}

// dont edit below this line
// below func syntax was incorrect in init commit
func getNames() (string, string) {
	return "John", "Doe"
}
