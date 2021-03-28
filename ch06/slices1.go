package main

import "fmt"

func main() {
	/*
		"A slice’s underlying array can’t grow in size. If there isn’t room in
		the array to add elements, all its elements will be copied to a new,
		larger array, and the slice will be updated to refer to this new array.
		But since all this happens behind the scenes in the append function,
		there’s no easy way to tell whether the slice returned from append has
		the same underlying array as the slice you passed in, or a different
		underlying array."
		ch06

	*/
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)

	// modify stuff
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
	/*
		go run slices1.go
		[s1 s1] [s1 s1 s2 s2] [s1 s1 s2 s2 s3 s3] [s1 s1 s2 s2 s3 s3 s4 s4]
		[s1 s1] [s1 s1 s2 s2] [XX s1 s2 s2 s3 s3] [XX s1 s2 s2 s3 s3 s4 s4]
	*/
}
