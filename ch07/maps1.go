package main

import "fmt"

func main() {
	jewelery := make(map[string]float64)
	jewelery["necklace"] = 89.99
	jewelery["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Earrings:", jewelery["earrings"])
	fmt.Println("Necklace:", jewelery["necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
}
