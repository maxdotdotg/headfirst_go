package main

// works because the greeting package is in ~/go/src/greeting
//import "greeting"

import (
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/deutsch"
)

func main() {
	greeting.Hi()
	greeting.Hello()
	deutsch.Hallo()
}
