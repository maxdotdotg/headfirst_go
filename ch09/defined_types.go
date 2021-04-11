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

	// receiver methods for different types
	// can have the same name
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	fmt.Printf("%0.3f liters equals %0.3f milliliters\n", soda, soda.ToMilliliters())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
	fmt.Printf("%0.3f milliliters equals %0.3f liters\n", water, water.ToLiters())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())

}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

// receiver method, only works on variables of type Liters
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

// receiver method, only works on variables of type Milliliters
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}
