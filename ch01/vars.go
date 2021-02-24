package main

import "fmt"

func main() {
	// long way to declare vars
	var quantity int
	var length, width float64
	var customerName string

	// assign values
	quantity = 4
	length, width = 1.2, 2.4
	customerName = "damon"

	// do it in one step and skip the type declaration
	var otherCustomer = "tim"

	// use'em
	fmt.Println(otherCustomer, "hasn't done shit")
	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	// zero values
	var yoInt int
	var yoFloat float64
	var yoString string
	var yoBool bool

	fmt.Println(yoInt)    // 0
	fmt.Println(yoFloat)  // 0
	fmt.Println(yoString) // ""
	fmt.Println(yoBool)   // false

}
