package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
)

func main() {
	date := calendar.Date{}
	/*
		// these don't work anymore because the fields
		// of the struct aren't exported
		date.Year = 2020
		date.Month = 33
		date.Day = 88
	*/

	// we have to use setter methods to set the values
	// and getter methods to get'em
	// I guess it's also convention to have the getter
	// use the same name as the field?
	// yep: "The Go community has decided on a convention
	// of leaving the Get prefix off of getter method names."
	// ch10
	date.SetYear(2020)
	date.SetMonth(12)
	date.SetDay(13)
	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
