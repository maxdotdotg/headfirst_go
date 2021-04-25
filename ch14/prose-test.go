package main

import (
	"fmt"
	"github.com/headfirstgo/prose"
)

func main() {
	// new version of JoinWithCommas doesn't work for two items
	phrases := []string{"my friends", "the best city"}
	fmt.Println("photos of", prose.JoinWithCommas(phrases))

	// but does work for three items
	phrases2 := []string{"my friends", "the best city", "an ice cream cone"}
	fmt.Println("photos of", prose.JoinWithCommas(phrases2))
}
