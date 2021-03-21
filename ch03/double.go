package main

import "fmt"

func main() {
	amount := 6

	// pass the _address_ of amount, not the value
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	// we take the _address_ as an argument
	// then we mutate the _value_
	*number *= 2
}
