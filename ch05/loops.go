package main

import "fmt"

func main() {
	notes := [7]string{"do", "ray", "me", "fa", "so", "la", "ti"}

	// loop over the array using the index
	for i := 0; i <= 2; i++ {
		fmt.Println(notes[i])
	}
	// loop over the array using len
	// this will throw an index out of range err
	// when i == 7, because we start at 0
	// so we cheat and subtract 1
	for i := 0; i <= (len(notes) - 1); i++ {
		fmt.Println(notes[i])
	}

	// use range
	for i, v := range notes {
		fmt.Println("index", i, "value", v)
	}

	// use range but lose the index
	for _, v := range notes {
		fmt.Println("value", v)
	}
}
