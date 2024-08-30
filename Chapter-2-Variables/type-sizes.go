package main

import "fmt"

func main() {
	accountAge := 2.6

	// create new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer
	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
