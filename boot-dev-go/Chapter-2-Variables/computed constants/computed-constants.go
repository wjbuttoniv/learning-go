package main

import "fmt"

func main() {
	const secondsInMinutes = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinutes * minutesInHour

		fmt.Println("the number of seconds in an hour:", secondsInHour)
}
