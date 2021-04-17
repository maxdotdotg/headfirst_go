package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	var coordinates geo.Coordinates
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates)

	// this breaks because of the validation
	// in the setter method
	coords2 := geo.Coordinates{}
	err = coords2.SetLongitude(200)
	if err != nil {
		log.Fatal(err)
	}
}
