package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("~/go/head-first-go/ch05/data.txt")
	if err != nil {
		log.Print(err)
	}
	fmt.Println(file)

	file2, err := os.Open("/home/user/go/head-first-go/ch05/data.txt")
	if err != nil {
		log.Print(err)
	}
	fmt.Println(file2)
}
