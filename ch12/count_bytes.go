package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	/*
		// run serially
		responseSize("https://www.google.com")
		responseSize("http://www.example.com")
		responseSize("https://golang.org")
	*/

	/*
		// run with goroutines, no channels, and a sleep
		go responseSize("https://www.google.com")
		go responseSize("http://www.example.com")
		go responseSize("https://golang.org")
		time.Sleep(time.Second)
	*/

	/*
		// create a channel and have the sizes sent there
		// right now this does run faster, but I have no
		// idea what output corresponds to what input
		sizes := make(chan int)
		go responseSize("https://www.google.com", sizes)
		go responseSize("http://www.example.com", sizes)
		go responseSize("https://golang.org", sizes)

		fmt.Println(<-sizes)
		fmt.Println(<-sizes)
		fmt.Println(<-sizes)
	*/

	// make urls a slice of strings, then loop over it
	urls := []string{"https://www.google.com", "http://www.example.com", "https://golang.org"}
	pages := make(chan Page)

	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

// pass a channel to capture the output since we
// can't use returns with goroutines
func responseSize(url string, channel chan Page) {
	// GET the URL
	fmt.Println("getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// close the network connection when main exits
	// why does it hold open a GET request? I'm confused
	defer response.Body.Close()

	// read the data of the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print the length of body, since body is a slice of bytes
	//fmt.Printf("body of %s is %d bytes\n", url, len(body))

	// send the number of bytes and the URL over a channel
	channel <- Page{URL: url, Size: len(body)}

}
