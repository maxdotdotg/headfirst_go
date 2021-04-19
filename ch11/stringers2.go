package main

import "fmt"

type Gallons float64

// satisfy the Stringer interface by implementing String()
// and here we decide how type Gallons gets printed
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Mililiters float64

func (m Mililiters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}
func main() {
	fmt.Println(Gallons(12.0987654))
	fmt.Println(Liters(12.0987654))
	fmt.Println(Mililiters(12.0987654))
}
