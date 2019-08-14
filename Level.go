package main

import (
	"fmt"
	"math/rand"

	"github.com/faiface/pixel/pixelgl"

	"github.com/faiface/pixel"
)

// Level ... Levels for the game
type Level struct {
	rooms            []Room // Each level has X rooms
	currentRoomIndex int
}

func createLevel(rooms []Room) Level {
	return Level{
		rooms,
		0,
	}
}

func (l *Level) updateRoom(player *Player, dt float64, win *pixelgl.Window) {
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		l.rooms[l.currentRoomIndex].objects[i].update(player)
	}
	for i := 0; i < len(l.rooms[l.currentRoomIndex].enemies); i++ {
		l.rooms[l.currentRoomIndex].enemies[i].update(dt, player.soundEmitter.waves)
	}
	l.rooms[l.currentRoomIndex].exec(player)
	if l.rooms[l.currentRoomIndex].hasRain {
		updateRain(l.rooms[l.currentRoomIndex].rainDeadZones, *player, dt)
		if l.rooms[l.currentRoomIndex].rainTimer < 0. {
			l.rooms[l.currentRoomIndex].rainTimer = l.rooms[l.currentRoomIndex].rainTimerMax
			rain = append(rain, createRain(pixel.V(float64(rand.Intn(int(winWidth))), winHeight)))
		} else {
			l.rooms[l.currentRoomIndex].rainTimer -= 1 * dt
		}
	}
}

func (l *Level) changeRoom(roomIndex int, player *Player, viewCanvas *pixelgl.Canvas) {
	if roomIndex < len(l.rooms) {
		l.currentRoomIndex = roomIndex
		l.setupRoom(player, viewCanvas)
		clearProjectiles(player)
	} else {
		fmt.Println("Room ", roomIndex, " does not exist!")
	}
}

func clearProjectiles(player *Player) {
	player.soundEmitter.waves = []SoundWave{} // Clear sound waves from player
}

func (l *Level) setupRoom(player *Player, viewCanvas *pixelgl.Canvas) {
	player.pos = l.rooms[l.currentRoomIndex].playerStartPos
	foregroundObjects = []Object{}
	backgroundObjects = []Object{}
	currentShader = l.rooms[l.currentRoomIndex].shader
	viewCanvas.SetFragmentShader(currentShader)

	if l.rooms[l.currentRoomIndex].hasSoundWaves {
		player.allowSoundEmitter = true
	} else {
		player.allowSoundEmitter = false
	}
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		foregroundObjects = append(foregroundObjects, l.rooms[l.currentRoomIndex].objects[i])
	}
	for i := 0; i < len(l.rooms[l.currentRoomIndex].enemies); i++ {
		createNodes(pixel.V(17., 17.), &l.rooms[l.currentRoomIndex].enemies[i].nodes)
	}
}
