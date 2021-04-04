package main

import "fmt"

func main() {
	status("Alma")
	status("Carl")
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]

	// check if value is assigned
	// ok is true when values are assigned,
	// and false when using the zero value
	// known as the "comma-ok idiom"?
	if !ok {
		fmt.Printf("no grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing\n", name)
	}
}
