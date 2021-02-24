package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// both of these use return values
	// they'd be stored in a variable or
	// something I guess
	math.Floor(2.75)
	strings.Title("head first go")

	// this time print to stdout
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
}
