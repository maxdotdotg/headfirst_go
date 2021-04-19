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

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("loading", cargo)
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

	TryVehicle(Truck("fyordz not 150"))
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("test cargo")
	}
}
