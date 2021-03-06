package main

import (
	"fmt"
	"time"
)

func main() {
	// create var now of type time.Time
	var now time.Time = time.Now()

	var year int = now.Year()
	fmt.Println(year, now)

}
