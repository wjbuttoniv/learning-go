package main

import (
	"fmt"
)

// Assignment
// complete the below function
// it should send 2 messages
//   - 1st to customer, then to spouse
//
// use sentSMS to send the msgToCustomer
//   - if err, return 0.0 and the Error
//
// do the same for msgToSpouse
// if both are successful, return total cost of messages added
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cstCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	spouseCost, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return cstCost + spouseCost, nil
}

// don't edit below this line

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}
