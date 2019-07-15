package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var quantity int
	//var length, width float64
	//var customerName string

	//quantity = 4
	//length, width = 1.2, 2.4
	//customerName = "Damon Cole"

	//var quantity int = 4
	//var length, width float64 = 1.2, 2.4
	//var customerName string = "Damon Cole"

	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(customerName))

	var mystring string
	fmt.Println(mystring)
	fmt.Println(reflect.TypeOf(mystring))
}
