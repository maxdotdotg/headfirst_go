package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a filename: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		// log the error and break
		log.Fatal(err)
	}
	// clean up filename
	filename := strings.TrimSpace(input)

	fileinfo, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileinfo.Size())
}
