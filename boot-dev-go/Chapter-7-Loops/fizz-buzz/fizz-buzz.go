package main

import "fmt"

// Assignment
// Complete the below function
// print number 1-100
// if
// -	multiple of 3, print fizz
// -	multiple of 5, print buzz
// -	both, print fizzbuzz
func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 && i%3 != 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
