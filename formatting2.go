package main

import "fmt"

func main() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("an int: %d\n", 15)
	fmt.Printf("a str: %s\n", "hello")
	fmt.Printf("a bool: %t\n", false)
	fmt.Printf("vals: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("more vals: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("percent sign: %%\n")

}
