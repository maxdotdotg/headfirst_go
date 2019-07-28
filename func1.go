package main

import (
	"fmt"
)

func main() {
	var amount, total float64
	var err error
	amount, err = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount, err = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount, err = paintNeeded(5.0, 3.3)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount, err = paintNeeded(5.0, -3.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Total: %0.2f liters\n", total)

}

// function name (input vars with types) return type {
// body of the function
// return some_stuff
// }

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
}
