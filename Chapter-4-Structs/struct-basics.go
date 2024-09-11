package main

import "fmt"

// Complete the below struct
// neogit test commit v2
type messageToSend struct {
	phoneNumber int
	message     string
}

// dont edit below this line

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to:%v\n", m.message, m.phoneNumber)
	fmt.Println("===============================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "And now a message from our sponser...",
	})
}
