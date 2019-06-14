package main

import (
	"github.com/faiface/pixel"
)

// Room ... It's a damn room for a level
type Room struct {
	objects        []Object
	playerStartPos pixel.Vec
	hasRain        bool
	rainTimer      float64
	rainTimerMax   float64
}

func createRoom(objects []Object, playerStartPos pixel.Vec, hasRain bool) Room {
	rainTimer := 0.000001
	rainTimerMax := rainTimer
	return Room{
		objects,
		playerStartPos,
		hasRain,
		rainTimer,
		rainTimerMax,
	}
}
