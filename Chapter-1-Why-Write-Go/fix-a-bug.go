package main

import "fmt"

func main() {
  messageFromDoris := []string{
    "You doing anything later?",
    "Did you get my last message?",
    "Don't leave me hanging...",
    "Please respond I'm lonely",
}

numMessages := float64(len(messageFromDoris))
  costPerMessage := 0.02

// do not make any changes above

// Line Below is the bug
  // Customer States:
  // "We are being charged inaccurately"
  // "The cost per message is 0.02"
totalCost := costPerMessage + numMessages

// do not make any changes below

fmt.Printf("Doris spent %.2f on messages today\n", totalCost)

}
