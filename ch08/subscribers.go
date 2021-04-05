package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	// doing it this way, without creating a custom type,
	// means creating a new struct for each instance
	var subscriber_var struct {
		name   string
		rate   float64
		active bool
	}

	subscriber_var.name = "Amand"
	subscriber_var.rate = 4.99
	subscriber_var.active = true

	fmt.Println("Subscriber:", subscriber_var.name)
	fmt.Println("monthly rate:", subscriber_var.rate)
	fmt.Println("active?", subscriber_var.active)

	subscriber1 := defaultSubscriber("lily")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("jimmy")
	printInfo(subscriber2)

	// pass a pointer to the struct
	applyDiscount(&subscriber2)
	printInfo(subscriber2)
}

func printInfo(s subscriber) {
	fmt.Println("Subscriber:", s.name)
	fmt.Println("monthly rate:", s.rate)
	fmt.Println("active?", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

// use this function to update the struct
// by taking a pointer to the struct
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
