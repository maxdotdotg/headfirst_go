package main

import (
	"fmt"
	"reflect"
)

// find types with reflect.TypeOf(VAR_GOES_HERE)
func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hi there"))
}
