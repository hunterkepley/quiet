package main

import (
	"github.com/faiface/pixel"
)

// Room ... It's a damn room for a level
type Room struct {
	objects        []Object
	playerStartPos pixel.Vec
	hasRain        bool
	rainDeadZones  []pixel.Rect
	rainTimer      float64
	rainTimerMax   float64
	shader         string
	exec           func(player *Player)
	hasSoundWaves bool
}

func createRoom(objects []Object, playerStartPos pixel.Vec, hasRain bool, rainDeadZones []pixel.Rect, shader string, exec func(player *Player), hasSoundWaves bool) Room {
	rainTimer := 0.000001
	rainTimerMax := rainTimer
	return Room{
		objects,
		playerStartPos,
		hasRain,
		rainDeadZones,
		rainTimer,
		rainTimerMax,
		shader,
		exec,
		hasSoundWaves,
	}
}
