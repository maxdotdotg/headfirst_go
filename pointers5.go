package main

import "fmt"

func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println("value: ", *myFloatPointer)
	fmt.Println("pointer: ", myFloatPointer)
}

func createPointer() *float64 { // return type is pointer
	var myFloat = 98.5
	return &myFloat // return pointer of myFloat

}
