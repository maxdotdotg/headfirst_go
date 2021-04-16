package main

import (
	"fmt"
	"geo"
)

func main() {
	var coordinates geo.Coordinates
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
