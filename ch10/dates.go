package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

// receiver function that takes a pointer
// because we need the change to persist
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}

	d.Year = year
	return nil
}

// we've added some validation to these receiver functions
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}

	d.Day = day
	return nil
}

func main() {
	date := Date{}
	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	date.SetDay(2)
	date.SetMonth(1)
	fmt.Println(date)
	fmt.Printf("%#v\n", date)

	// invalid values can still be set with literals
	date3 := Date{-1, 14, 99}
	fmt.Printf("%#v\n", date3)

	// invalid values can still be passed by accessing
	// values directly in the struct
	date4 := Date{}
	date4.Year = -2
	date4.Month = 22
	date.Day = 66
	fmt.Printf("%#v\n", date4)

	var date2 Date
	err = date2.SetYear(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", date2)

}
