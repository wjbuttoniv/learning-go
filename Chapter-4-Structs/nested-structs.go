package main

import (
	"fmt"
)

// Assignment: Need to verify required fields
// have non-zero values
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

// Need to eliminate
// raw comparisons
func isValidUser(userValid user) bool {
	return userValid.name != "" && userValid.number != 0
}

func canSendMessage(mToSend messageToSend) bool {
	return isValidUser(mToSend.sender) && isValidUser(mToSend.recipient)
}

// dont touch below this line

func test(mToSend messageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient.name,
		mToSend.recipient.number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
	test(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name:   "Njorn Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})
}
