// count tallies the number of times each line occurs in a file
package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	// build two slices where counts[i] is the number of times
	// that names[i] occurs in votes.txt

	var names []string
	var counts []int

	for _, line := range lines {
		matched := false

		// if this line matches the current name, increment count
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}

		// if the line doesn't match the current name, add it to names
		// and add a count of 1
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
	fmt.Println(lines)
}
