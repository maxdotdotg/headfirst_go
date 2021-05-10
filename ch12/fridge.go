package main

import "fmt"

// TIL defer runs before panic
func snack() {
	defer fmt.Println("closing fridge")
	fmt.Println("opening fridge")
	panic("empty fridge")
}

func main() {
	snack()
}

/* expected stdout
opening fridge
closing fridge??? IDK
panic empty fridge
main.snack() stack trace
*/
