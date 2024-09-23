package main

import "fmt"

func bulkSend(numMessages int) {
	// ?
}

// dont edit below this line

func test(numMessages int) float64 {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulek send complete! Cost = %.2f\n", cost)
	fmt.Println("===================================================================")
}
