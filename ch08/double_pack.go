package main

import "fmt"

type part struct {
	description string
	count       int
}

// use a pointer because we need access to
// the underlying data, and passing the variable
// itself will only modify the value when in scope
// of the doublePack function; this change will not
// persist, and we will not have moar parts
func doublePack(p *part) {
	p.count *= 2
}

func main() {
	var fuses part
	fuses.description = "fuses"
	fuses.count = 5
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
	doublePack(&fuses)
	fmt.Println(fuses.count)
}
