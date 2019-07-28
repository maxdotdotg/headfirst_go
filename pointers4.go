package main

import "fmt"

func main() {
	myInt := 6
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
