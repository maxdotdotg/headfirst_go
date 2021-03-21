package main

import "fmt"

func main() {
	var myInt int
	// initialize variable of type pointer to int
	// to use it, we have to give it an address
	var myPointer *int

	myInt = 42

	// assign myPointer the _address_ of myInt
	myPointer = &myInt
	fmt.Println(*myPointer)
	fmt.Println(myPointer)
}
