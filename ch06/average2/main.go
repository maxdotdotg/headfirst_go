// Package average2 reads a space-delimited list of float64s from the command line
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	/*
		fmt.Printf("%#v", os.Args)
		average2 git:(main) ✗ go run main.go hello
		[]string{"/tmp/go-build406762656/b001/exe/main", "hello"}%
	*/

	// take args and skip the first,
	// which is the path to the bin itself
	arguments := os.Args[1:]
	var numbers []float64

	// convert the slice of strings into a slice of float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// print float with 2 decimal places
	// add an ellipsis (...) following the slice you want to use in place of variadic arguments.
	// ch06
	fmt.Printf("average is: %0.2f\n", average(numbers...))
}

// average takes a slice of float64s and returns their average as a float64
func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
