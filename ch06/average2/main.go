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

	// create an empty var for sum and average
	var sum float64 = 0

	// convert the slice of strings into a slice of float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Printf("total is: %0.2f\n", sum)

	sampleCount := float64(len(arguments))

	// calculate average
	average := sum / sampleCount
	// print float with 2 decimal places
	fmt.Printf("average is: %0.2f\n", average)
}
