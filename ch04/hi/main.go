package main

// works because the greeting package is in ~/go/src/greeting
import "greeting"

func main() {
	greeting.Hi()
	greeting.Hello()
}
