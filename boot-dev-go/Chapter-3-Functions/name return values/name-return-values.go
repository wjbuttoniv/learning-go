package main

import "fmt"

// fix bug by using named
// return values in func sig
func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

// Output looked really ugly with
// "You are x in 0 years"
// third attempt to make Output
// more presentable

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	if yearsUntilCarRental == 0 {
		fmt.Println("You are an adult", "\nYou can drink", "\nYou can rent a car")
	} else if yearsUntilCarRental != 0 && yearsUntilDrinking == 0 {
		fmt.Println("You are an adult", "\nYou can drink", "\nYou can rent a car in", yearsUntilCarRental, "years")
	} else if yearsUntilDrinking != 0 && yearsUntilAdult == 0 {
		fmt.Println("You are an adult", "\nYou can drink in", yearsUntilDrinking, "years", "\nYou can rent a car in", yearsUntilCarRental, "years")
	} else {
		fmt.Println("You are an adult in", yearsUntilAdult, "years", "\nYou can drink in", yearsUntilDrinking, "years", "\nYou can rent a car in", yearsUntilCarRental, "years")
	}
}

//func test(age int) {
//fmt.Println("Age:", age)
//yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
//if yearsUntilAdult == 0 {
//fmt.Println("You are an adult")
//} else {
//fmt.Println("You are an adult in", yearsUntilAdult, "years")
//}
//if yearsUntilDrinking == 0 {
//fmt.Println("You can drink")
//} else {
//fmt.Println("You can drink in", yearsUntilDrinking, "years")
//}
//if yearsUntilCarRental == 0 {
//fmt.Println("You can rent a car")
//	} else {
//		fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
//	}
//	fmt.Println("=======================================")
//}

func main() {
	test(4)
	test(10)
	test(22)
	test(35)
}
