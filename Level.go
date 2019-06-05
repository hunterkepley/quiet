package main

import "github.com/faiface/pixel/pixelgl"

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

func (l *Level) renderRoom(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, dt float64) {
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		l.rooms[l.currentRoomIndex].objects[i].render(win, viewCanvas)
	}
}

func (l *Level) updateRoom(player *Player, dt float64) {
	if !l.rooms[l.currentRoomIndex].roomStarted {
		player.pos = l.rooms[l.currentRoomIndex].playerStartPos
		l.rooms[l.currentRoomIndex].roomStarted = true
	}
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		l.rooms[l.currentRoomIndex].objects[i].update(player)
	}
}
