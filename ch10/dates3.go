package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main() {
	// the Event type has Date embedded in it
	// this means the variable `event` can use exported
	// methods associated with the Event type and the Date type
	event := calendar.Event{}
	err := event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(29)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", event)
	fmt.Println(event)
	fmt.Println(event.Day())
	fmt.Println(event.Month())
	fmt.Println(event.Year())

	// accessing fields can be done by using the embedded type as well
	fmt.Println(event.Date.Year())

	err = event.SetTitle("today is the day!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", event)
	fmt.Println(event.Title())
}
