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

	// we have to use setter methods now
	date.SetYear(2020)
	date.SetMonth(12)
	date.SetDay(13)
	fmt.Println(date)
	fmt.Println(date.GetYear())
	fmt.Println(date.GetMonth())
	fmt.Println(date.GetDay())
}
