package main

import (
	"fmt"
)

func main() {
	var length float64 = 1.2
	var width int = 2

	//length = width
	// above doesn't work because of type mis-match
	//length = float64(width)
	fmt.Println(length)
	fmt.Println("area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
}
