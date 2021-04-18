package main

import "github.com/headfirstgo/gadget"

// "We could define [the Player interface] in the gadget
// package instead, but defining the interface in the same
// package where we use it gives us more flexibility."
// ch11

type Player interface {
	Play(string)
	Stop()
}

/*
func playList(device gadget.TapePlayer, songs []string) {
	// right now, playList only works with TapePlayer,
	// despite TapeRecorder having the same functions
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
*/

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"battery!", "shifting planes", "pulsar III"}
	playList(player, mixtape)

	anotherDevice := gadget.TapeRecorder{}
	anotherTape := []string{"the island, part ii", "hold your colour", "tarantula"}
	playList(anotherDevice, anotherTape)
}
