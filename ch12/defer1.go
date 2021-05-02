package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	// defer makes this line run after Socialize() exits
	// this will happen even when errors get thrown
	defer fmt.Println("goodbye")
	fmt.Println("hello")

	// Socialize exits here
	return fmt.Errorf("I don't feel like talking")

	// these never get called
	fmt.Println("weather, it's nice")
	return nil
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
