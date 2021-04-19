package main

import (
	"github.com/headfirstgo/gadget"
)

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

	// this will work at runtime because type conversion is used to access
	// methods associated with the concrete type
	TryOut(gadget.TapeRecorder{})

	// this won't work at runtime for the same reason
	// "panic: interface conversion: main.Player is gadget.TapePlayer,
	// not gadget.TapeRecorder"
	TryOut(gadget.TapePlayer{})
}

func TryOut(player Player) {
	player.Play("test track")
	player.Stop()

	// type assertion, now with error handling
	recorder, ok := player.(gadget.TapeRecorder)

	if ok {
		recorder.Record()
	}
}
