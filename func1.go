package main

import "fmt"

func main() {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount = paintNeeded(5.0, 3.3)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	fmt.Printf("Total: %0.2f liters\n", total)

}

// function name (input vars with types) return type {
// body of the function
// return some_stuff
// }

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}
