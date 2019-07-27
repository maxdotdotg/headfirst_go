package main

import "fmt"

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("%0.2f liters of paint needed\n", area/10.0)
	// resultString := fmt.Sprintf("about 1/3: %0.2f\n", 1.0/3.0)
	// fmt.Println(resultString)

	width = 5.2
	height = 3.5
	area = width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)

}

//func area(width,height)
