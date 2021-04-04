package main

import "fmt"

func main() {
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool

	value, ok = counters["a"]
	fmt.Println(value, ok) // prints "3 true", because there's a value at "a"
	value, ok = counters["b"]
	fmt.Println(value, ok) // prints "0 true", because there's a value at "b"
	value, ok = counters["c"]
	fmt.Println(value, ok) // prints "0 false" because there's NO value at "c"
}
