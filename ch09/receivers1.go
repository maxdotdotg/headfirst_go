package main

import "fmt"

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

// receiver functions can exist for pointers as well
func (n *Number) Double() {
	// totally not hard to read at all, obviously
	*n *= 2
}
func main() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)

	four := Number(4)
	four.Add(3)
	four.Subtract(2)

	/*
		When you call a method that requires a pointer receiver
		on a variable with a nonpointer type, Go will automatically
		convert the receiver to a pointer for you
	*/
	ten.Double()
	fmt.Println(ten)
}
