package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {
	// commented out to use datafile
	// set up the array
	// numbers := [3]float64{71.8, 56.2, 89.5}

	// double-check the path
	numbers, err := datafile.GetFloats("../data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// create an empty var for sum and average
	var sum float64 = 0
	var average float64 = 0

	// add the elements of the numbers array
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("total is:", sum)

	// calculate average
	average = sum / float64(len(numbers))
	// print float with 2 decimal places
	fmt.Printf("average is: %0.2f\n", average)

}
