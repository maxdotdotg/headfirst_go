package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	fmt.Println(doorNumber)

	if doorNumber == 1 {
		fmt.Println("win a book")
	} else if doorNumber == 2 {
		fmt.Println("win a goat")
	} else if doorNumber == 3 {
		fmt.Println("win a game")
	} else {
		panic("invalid door number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
