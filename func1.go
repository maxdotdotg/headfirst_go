package main

import "fmt"

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
}

// function name (input vars with types) return type {
// body of the function
// return some_stuff
// }

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / 10.0
}
