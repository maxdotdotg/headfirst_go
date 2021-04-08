package main

import (
	"fmt"
	"geo"
)

func main() {
	location := geo.Landmark{}
	location.Name = "some place"
	location.Latitude = 90.6
	location.Longitude = -100
	fmt.Println(location)
}
