package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("speeding up")
}

func (c Car) Brake() {
	fmt.Println("stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("speeding up")
}

func (t Truck) Brake() {
	fmt.Println("stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}

func main() {
	var vehicle Vehicle = Car("moo-stang")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("fjord")
	vehicle.Brake()
	vehicle.Steer("right")
}
