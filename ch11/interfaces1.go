package main

import (
	"fmt"
	"github.com/headfirstgo/mypkg"
)

func main() {
	// declare a variable using the interface type?
	var value mypkg.MyInterface

	value = mypkg.MyType(5)
	value.MethodWithoutParams()
	value.MethodWithParams(6.25)
	fmt.Println(value.MethodWithRreturnValue())
}
