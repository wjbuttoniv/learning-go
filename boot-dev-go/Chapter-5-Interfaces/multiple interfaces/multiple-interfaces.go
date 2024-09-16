package main

import (
	"fmt"
)

// Assignment
// add the required methods so that the email type
// implements both the expense and printer interfaces

// if the email is not Subscribed, cost is 0.05 per char
// if it is Subscribed, cost is 0.01 per char
const (
	subCost = 0.01
	stdCost = 0.05
)

func (e email) cost() float64 {
	eLength := float64(len(e.body))
	if e.isSubscribed == true {
		return subCost * eLength
	}
	return stdCost * eLength
}

// print method should print emails body text
// to standard out
func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("======================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}
