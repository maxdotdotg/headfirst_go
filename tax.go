package main

import (
	"fmt"
)

func main() {
	price := 100
	fmt.Println("price is", price, "dollars")

	taxRate := 0.08
	tax := taxRate * float64(price)
	fmt.Println("tax is", int(tax), "dollars")

	total := float64(price) + tax
	fmt.Println("Total cost is", int(total), "dollars")

	availableFunds := 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("Within budget?", float64(availableFunds) > total)

}
