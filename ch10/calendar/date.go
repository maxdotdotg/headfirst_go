package calendar

import (
	"errors"
)

type Date struct {
	year  int
	month int
	day   int
}

// receiver function that takes a pointer
// because we need the change to persist
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}

	d.year = year
	return nil
}

// we've added some validation to these receiver functions
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}

	d.day = day
	return nil
}

// Getter methods, get access to non-exported struct fields
func (d *Date) GetYear() int {
	return d.year
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) GetDay() int {
	return d.day
}
