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

	fmt.Println(lines)
}
