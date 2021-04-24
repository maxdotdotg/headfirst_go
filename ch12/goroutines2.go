package main

import "fmt"

// this function takes a channel that carries strings
func greeting(myChannel chan string) {
	// send this value over the channel
	myChannel <- "hi"
}

func main() {
	// create a channel that carries strings
	myChannel := make(chan string)

	// pass the channel to a function
	// that runs in a goroutine
	go greeting(myChannel)

	// receive values from the channel
	val := <-myChannel
	fmt.Println(val)

	// fmt.Println(<-myChannel) would also have worked
}
