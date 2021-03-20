package main

import (
	"fmt"
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

	fmt.Println(litersOfPaint(4.2, 3.0))
	fmt.Println(litersOfPaint(5.2, 3.5))
}

func litersOfPaint(width float64, height float64) string {
	// function definition!
	// define required args with types, and define return type
	area := width * height

	// fmt.Sprintf _returns_ a formatted string instead
	// of printing to stdout
	return fmt.Sprintf("%.2f liters of paint needed", area/10)
}
