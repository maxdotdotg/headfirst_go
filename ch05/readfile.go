package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open the file for reading
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// create scanner for file
	scanner := bufio.NewScanner(file)

	// loop through until EOF and Scan returns false
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// handle errors
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
