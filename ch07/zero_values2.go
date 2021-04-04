package main

import (
	"fmt"
	"sort"
)

func main() {
	status("Alma")
	status("Carl")

	// duplication for the same of the sorting example
	grades := map[string]float64{"Alma": 60.4, "Rohit": 86.5, "Carl": 50.4}
	var names []string
	for name := range grades {
		names = append(names, name)
	}

	// sort alpabetically
	// and do so by permanently modifying the content of the variable?
	// that's pretty weird...
	sort.Strings(names)
	for _, name := range names {
		// why is this using double-percentage signs? IDK
		// prints "Alma has a grade of 60.4%"
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}

}

func status(name string) {
	grades := map[string]float64{"Alma": 60.4, "Rohit": 86.5, "Carl": 50.4}
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
