package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("today's location")
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLatitude(33.22)
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLongitude(-122.09)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())

}
