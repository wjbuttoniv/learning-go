package main

import (
	"fmt"
)

// Assignment
// Update the below so that divideError implements the error interface
// Error should return in the following way:
//
//	can not divide DIVIDEND by zero
//
// where DIVIDEND is the actual dividend of the divideError
type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("cannot divide %v by zero", d.dividend)
}

// don't edit below this line

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// We convest the `divideError` struct to an `error` type by returning it
		// as an error. As an error type, when it's printed its default value
		// will be the result of the Error() method
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
