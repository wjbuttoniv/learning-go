package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

// add missing getSalary() method
func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

// dont touch below this line

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func main() {
	test(fullTime{
		name:   "Jack",
		salary: 5000,
	})
	test(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})
}
