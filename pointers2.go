package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	var myFloat int
	var myFloatPointer *int
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
}
