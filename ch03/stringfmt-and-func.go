package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		string formatting notes!
		%5.3f
		% start of formatting verb
		5 min width of the entire number, with spaces for padding of needed
		3 width after decimal point
		f formatted verb type (float, in this case)

		remember, printf doesn't add a newline
	*/

	/*
		var width, height, area float64

		width = 4.2
		height = 3.0
		area = width * height
		fmt.Printf("%.2f liters of paint needed\n", area/10)

		width = 5.2
		height = 3.5
		area = width * height
		fmt.Printf("%.2f liters of paint needed\n", area/10)
	*/

	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters of paint needed\n", amount)

	amount, err = paintNeeded(-4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.2f liters of paint needed\n", amount)
}

func paintNeeded(width float64, height float64) (float64, error) {
	// function definition!
	// define required args with types, and define return type
	if width < 0 {
		return 0, fmt.Errorf("width can't be negative, you entered %f", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("height can't be negative, you entered %f", height)
	}

	area := width * height

	/*
		// fmt.Sprintf _returns_ a formatted string instead
		// of printing to stdout
		return fmt.Sprintf("%.2f liters of paint needed", area/10)
	*/
	return area / 10.0, nil

}
