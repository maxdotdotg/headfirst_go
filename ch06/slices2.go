package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3]
	slice = append(slice, "x")

	// don't know why I thought this would append a concat'd string
	// but it doesn't, just a regular append of one item after the other
	slice = append(slice, "y", "z")
	for _, letter := range slice {
		fmt.Println(letter)
	}
	/*
		should print out
		b
		c
		d NOPE, forgot, slices are exclusive
		e
		x
		y z? NOPE, no concatenation
		y
		z
	*/

}
