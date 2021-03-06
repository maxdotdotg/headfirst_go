// pass_fail checks if a grade is passing for failing

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	fmt.Println(input)
}
