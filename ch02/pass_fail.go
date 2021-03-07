// pass_fail checks if a grade is passing for failing

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
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)

	// gotta use all vars and catch errors
	// bufio.NewReader.ReadString returns a string or an err

	// we're ignoring errors, so :shrug:
	// in this case, use the blank variable "_"
	// input, _ := reader.ReadString('\n')

	input, err := reader.ReadString('\n')
	if err != nil {
		// log the error and break
		log.Fatal(err)
	}

	// strip newlines from input and convert to float
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		// log the error and break
		log.Fatal(err)
	}

	if grade < 60 {
		fmt.Println("failing")
	} else {
		fmt.Println("passing")
	}
}
