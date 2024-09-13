package chapter5interfaces_1

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	// ?
}

type message interface {
	// ?
}

// dont edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on $s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberofSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberofSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("=================================")
}

func chapter5interfaces_1() {
	test(sendingReport{
		reportName:    "First Report",
		numberofSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberofSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer ",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, time.UTC),
	})
}
