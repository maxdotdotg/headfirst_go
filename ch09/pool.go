package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println("Sleepy creel population", population)
	fmt.Println("welcome, new neighbor")
	population += 1
	fmt.Println("Sleepy creel population", population)

}
