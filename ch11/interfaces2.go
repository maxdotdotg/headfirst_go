package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("tweet")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("honk")
}

type NoiseMaker interface {
	MakeSound()
}

type Robot string

func (r Robot) Walk() {
	fmt.Println("powering legs")
}

func (r Robot) MakeSound() {
	fmt.Println("boop boop")
}

// and we can have functions that take interfaces as parameters
func play(n NoiseMaker) {
	n.MakeSound()
}

// both Whistle and Horn satisfy the NoiseMaker interface
// so the varable "toy" can be of either type
func main() {
	var toy NoiseMaker
	fmt.Printf("%#v\n", toy)
	toy = Whistle("toyco Canary")
	toy.MakeSound()
	fmt.Printf("%#v\n", toy)
	toy = Horn("toyco blaster")
	toy.MakeSound()
	fmt.Printf("%#v\n", toy)

	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))

	// type assertions
	var noiseMaker NoiseMaker = Robot("Gundam Inc")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()

}
