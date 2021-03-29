package main

import "fmt"

func main() {
	fmt.Println(sum(9, 7))
	fmt.Println(sum(1, 2, 4))
}

// sum adds the elements of a slice of integers and returns their sum as an int
func sum(numbers ...int) int {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
