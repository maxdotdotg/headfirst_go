package main

import (
	//	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// all of this doesn't work, because stripping newlines
	// is fucking tedious

	// fmt.Print("Enter a filename: ")
	// reader := bufio.NewReader(os.Stdin)
	// filename, err := reader.ReadString('\n')
	// if err != nil {
	// 	// log the error and break
	// 	log.Fatal(err)
	// }
	//fileinfo, err := os.Stat(filename)

	fileinfo, err := os.Stat("notes")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileinfo.Size())
}
