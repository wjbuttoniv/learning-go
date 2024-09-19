package main

import "fmt"

// Assignment
// Complete the below function
// it should return a string in this format
// SMS that costs $COST to be sent to 'RECIPIENT' cannot be sent
// cost is the cost of the sms, always formatted %.2f
// recipient is stringified representation of recipient phone number
func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' cannot be sent", cost, recipient)
}

// don't edit below this line

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(32.1, "+1 (801) 555 7456")
	test(14.4, "+1 (234) 555 6545")
}
