package main

import "fmt"

func printPointer(myBoolPointer *bool) { // pointer as input
	fmt.Println(*myBoolPointer) // print value of pointer
}

func main() {
	var mybool bool = true
	printPointer(&mybool) // pass pointer
}
