package main

import "fmt"

// Alter the below to capture the return value of incrementSends
// and overwrite previous sendsSoFar value
func main() {
  sendsSoFar := 430
  const sendsToAdd = 25
  sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
  fmt.Println("you've sent", sendsSoFar, "messages")
}

// Alter the below to increment sendsSoFar before returning
func incrementSends(sendsSoFar, sendsToAdd int) int {
  sendsSoFar++
  sendsSoFar = sendsSoFar + sendsToAdd
  return sendsSoFar
}
