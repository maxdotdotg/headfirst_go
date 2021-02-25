package main

/* print out
I started with 10 apples.
Some jerk ate 4 apples.
There are 6 apples left
*/

import "fmt"

func main() {
	var originalCount = 10
	var eatenCount = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("some jerk ate", eatenCount, "apples.")
	fmt.Println("there are", originalCount-eatenCount, "apples left")
}
