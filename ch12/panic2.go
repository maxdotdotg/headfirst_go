package main

import "fmt"

func calmDown() {
	p := recover()
	// assert the type of panic is "error"
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("an error")
	panic(err)

}
