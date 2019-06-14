package main

import (
	"fmt"
	"math/rand"

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

func (l *Level) updateRoom(player *Player, dt float64) {
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		l.rooms[l.currentRoomIndex].objects[i].update(player)
	}
	if l.rooms[l.currentRoomIndex].hasRain {
		updateRain(dt)
		if l.rooms[l.currentRoomIndex].rainTimer < 0. {
			l.rooms[l.currentRoomIndex].rainTimer = l.rooms[l.currentRoomIndex].rainTimerMax
			rain = append(rain, createRain(pixel.V(float64(rand.Intn(int(winWidth))), winHeight)))
		} else {
			l.rooms[l.currentRoomIndex].rainTimer -= 1 * dt
		}
	}
}

func (l *Level) changeRoom(roomIndex int, player *Player) {
	if roomIndex < len(l.rooms) {
		l.currentRoomIndex = roomIndex
		l.setupRoom(player)
	} else {
		fmt.Println("Room ", roomIndex, " does not exist!")
	}
}

func (l *Level) setupRoom(player *Player) {
	player.pos = l.rooms[l.currentRoomIndex].playerStartPos
	foregroundObjects = []Object{}
	backgroundObjects = []Object{}
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		foregroundObjects = append(foregroundObjects, l.rooms[l.currentRoomIndex].objects[i])
	}
}
