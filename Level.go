package main

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

/*func (l *Level) renderRoom(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, dt float64) {
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		l.rooms[l.currentRoomIndex].objects[i].render(win, viewCanvas)
	}
}*/

func (l *Level) updateRoom(player *Player, dt float64) {
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		l.rooms[l.currentRoomIndex].objects[i].update(player)
	}
}

func (l *Level) changeRoom(roomIndex int, player *Player) {
	l.currentRoomIndex = roomIndex
	l.setupRoom(player)
}

func (l *Level) setupRoom(player *Player) {
	player.pos = l.rooms[l.currentRoomIndex].playerStartPos
	foregroundObjects = []Object{}
	backgroundObjects = []Object{}
	for i := 0; i < len(l.rooms[l.currentRoomIndex].objects); i++ {
		foregroundObjects = append(foregroundObjects, l.rooms[l.currentRoomIndex].objects[i])
	}
}
