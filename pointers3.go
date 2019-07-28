package main

import "fmt"

func main() {
	myInt := 4
	myIntPointer := &myInt     // pointer to var
	fmt.Println(myIntPointer)  // address of pointer
	fmt.Println(*myIntPointer) // value at pointer

	myFloat := 98.6
	myFloatPointer := &myFloat   // pointer to var
	fmt.Println(myFloatPointer)  // address of pointer
	fmt.Println(*myFloatPointer) // value at pointer

	myBool := true
	myBoolPointer := &myBool    // pointer to var
	fmt.Println(myBoolPointer)  // address of pointer
	fmt.Println(*myBoolPointer) // value at pointer
}
