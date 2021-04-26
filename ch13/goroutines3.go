package main

import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// create channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	// pass channels to functions using goroutines
	go abc(channel1)
	go def(channel2)

	// receive and print values
	// order is consistent because channels block until the value is received
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
}
