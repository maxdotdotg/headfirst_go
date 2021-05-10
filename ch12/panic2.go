package main

import "fmt"

func main() {
	freakOut()
	fmt.Println("exiting normally")
}

func calmDown() {
	recover()
}

func freakOut() {
	defer calmDown()
	panic("ohno")
}
