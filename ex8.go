package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars")

	var taxRate float32 = 0.08
	var tax float32 = float32(price) * taxRate
	fmt.Println("Tax is", tax, "dollars")

	var total float32 = float32(price) + tax
	fmt.Println("Total cost is", total, "dollars")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("Within budget?", total <= float32(availableFunds))
}
