package main

import "fmt"

func main() {
	// print a hand-drawn table
	// width of 12 for the string, width of 2 for the int
	fmt.Printf("%12s | %s\n", "product", "cost in cents")
	fmt.Printf("--------------------------------\n")
	fmt.Printf("%12s | %2d\n", "paper clips", 5)
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "tape", 99)

}
