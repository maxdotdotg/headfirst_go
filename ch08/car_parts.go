package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	// declare a variable of type car
	var porche car
	porche.name = "porche of some model"
	porche.topSpeed = 99.9 // IDGAF about cars, this is a nice number
	fmt.Println("Name:", porche.name)
	fmt.Println("top speed:", porche.topSpeed)

	// declare a variable of type part
	var bolts part
	bolts.description = "hex bolts"
	bolts.count = 88
	showInfo(bolts)

	otherPart := minimumOrder("another special car part")
	fmt.Println(otherPart.description, otherPart.count)
}

func showInfo(p part) {
	// this function accepts arguments of type "part" only
	// this is also when we start to use short variable names
	// because we hate people
	fmt.Println("description:", p.description)
	fmt.Println("count:", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}
