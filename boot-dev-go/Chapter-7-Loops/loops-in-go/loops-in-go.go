package main

import "fmt"

// Assignment
// complete the below function
// return total cost as float64
// each messages costs 1.0 plus a fee
// 1st- 0.0, 2nd, 0.1, etc
// use a loop to calculate total cost and return it
func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

// dont edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulek send complete! Cost = %.2f\n", cost)
	fmt.Println("===================================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
