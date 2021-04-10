package main

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func main() {
	// defined types that don't use structs
	// used for readability, I suppose?
	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)

	// can't assign different defined types, but you can convert
	// when both have the same underlying type (here, float64)
	carFuel = Gallons(Liters(40) * 0.264)
	busFuel = Liters(Gallons(63) * 3.785)
	fmt.Println(carFuel, busFuel)

	// converting between defined types that use the same
	// underlying type
	carFuel += ToGallons(Liters(30.0))
	busFuel += ToLiters(Gallons(40.0))
	fmt.Println(carFuel, busFuel)

}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}
