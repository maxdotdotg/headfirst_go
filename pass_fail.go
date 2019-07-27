// pass_fail reports if a student's grade is passing or failing

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("enter a grade: ")
	reader := bufio.NewReader(os.Stdin)   // buffered read from stdin
	input, err := reader.ReadString('\n') //return up up to newline
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // strip the new line
	grade, err := strconv.ParseFloat(input, 64) // convert to float64
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("a grade of", grade, "is", status)
}
