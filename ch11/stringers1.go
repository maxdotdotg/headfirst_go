package main

import "fmt"

type CoffeePot string

// satisfy the Stringer interface
func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

// use fmt.Print stuff and it'll use the Stringer
// interface under the hood
func main() {
	coffeePot := CoffeePot("french press")
	fmt.Println(coffeePot.String())
	fmt.Println(coffeePot)
	fmt.Printf("%s\n", coffeePot)
}
