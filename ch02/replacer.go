package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# kick r#cks"

	// replacer, variable of type strings.NewReplacer
	replacer := strings.NewReplacer("#", "o")

	// calling the Replace method on a string
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
