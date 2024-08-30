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

// Corrected bug from commit c5b7d157 
  // incorrect operation of addition instead of multiplication
totalCost := costPerMessage * numMessages

fmt.Printf("Doris spent %.2f on messages today\n", totalCost)

}
