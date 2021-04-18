package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

// setter
func (e *Event) SetTitle(title string) error {
	// validation!
	// check to make sure the event title isn't too long
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

// getter
func (e *Event) Title() string {
	return e.title
}
