// pass_fail checks if a grade is passing for failing

package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")

	// logic to read stdin and parse
	// factored into getFloat()
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	if grade < 60 {
		fmt.Println("failing")
	} else {
		fmt.Println("passing")
	}
}
