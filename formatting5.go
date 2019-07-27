package main

import "fmt"

func main() {
	// print the float with a min with of 5
	// and 3 decimal places
	fmt.Printf("%5.3f\n", 2.3)
	fmt.Printf("%5.3f\n", 00000000000.3)
	fmt.Printf("%5.3f\n", 213.3)
	fmt.Printf("%5.3f\n", 23.0)
	fmt.Printf("%5.3f\n", 02.3333333333333333333)
}
