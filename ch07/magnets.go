package main

import "fmt"

func main() {
	ranks := map[string]int{"gold": 1, "silver": 2, "bronze": 3}

	for medal, rank := range ranks {
		fmt.Printf("the %s medal's rank is %d\n", medal, rank)
	}
}

/*
print to stdout
the gold medal's rank is 1
the silver medal's rank is 2
the bronze medal's rank is 3
*/
