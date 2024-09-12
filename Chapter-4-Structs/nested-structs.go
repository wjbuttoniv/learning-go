package main

// Assignment: Need to verify required fields
// have non-zero values
type messageToSend struct {
	message  string
	sender   user
	recipent user
}

type user struct {
	name   string
	number int
}

// Complete the below function
// should only return true if
// sender and recipent each contain
// a name and number
// else return false
func canSendMessage(mToSend messageToSend) bool {
	return true
}

// dont touch below this line

func test(mToSend messageToSend) {
}
