package main

import (
	"github.com/faiface/pixel"
)

type Room struct {
	objects        []Object
	playerStartPos pixel.Vec
}

func createRoom(objects []Object, playerStartPos pixel.Vec) Room {
	return Room{
		objects,
		playerStartPos,
	}
}
