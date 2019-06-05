package main

import (
	"github.com/faiface/pixel"
)

type Room struct {
	objects        []Object
	playerStartPos pixel.Vec
	roomStarted    bool // If the room has been instantiated or not
}

func createRoom(objects []Object, playerStartPos pixel.Vec) Room {
	return Room{
		objects,
		playerStartPos,
		false,
	}
}
