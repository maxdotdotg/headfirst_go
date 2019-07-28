package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(&myInt)                 // print the pointer
	fmt.Println(reflect.TypeOf(&myInt)) // print the type

	var myFloat float64
	fmt.Println(&myFloat)                 // print the pointer
	fmt.Println(reflect.TypeOf(&myFloat)) // print the type

	var myBool bool
	fmt.Println(&myBool)                 // print the pointer
	fmt.Println(reflect.TypeOf(&myBool)) // print the type
}
