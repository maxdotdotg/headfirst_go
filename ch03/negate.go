package main

import "fmt"

func main() {
	truth := true
	fmt.Println("was", truth)
	negate(&truth)
	fmt.Println("now is", truth)

	lies := false
	fmt.Println("was", lies)
	negate(&lies)
	fmt.Println("now is", lies)
}

func negate(myBool *bool) {
	// take the address of the bool as an arg
	// update the value with its opposite
	*myBool = !*myBool
}
