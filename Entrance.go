package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/**
 * This file is for entrances, like a door to another room. They are just zones defined, no image attached.
 */

// Entrance ... An entrace to a new room
type Entrance struct {
	pos              pixel.Vec
	size             pixel.Vec
	floatingUI       FloatingUI
	roomIndex        int // What room to switch to in the current level, if -1, changes level instead
	levelIndex       int // What level to switch to, if -1, changes room instead
	renderFloatingUI bool
}

func createEntrance(pos pixel.Vec, size pixel.Vec, bounceRange float64, roomIndex int, levelIndex int) Entrance {
	floatingUIPosition := pixel.V(pos.X, pos.Y+size.Y)
	return Entrance{
		pos,
		size,
		createFloatingUI(floatingUIPosition, floatingImages.e, bounceRange),
		roomIndex,
		levelIndex,
		false,
	}
}

func (e *Entrance) update(dt float64) {
	e.playerCollision(&player)
	if e.renderFloatingUI {
		e.floatingUI.update(dt)
	}
}

func (e *Entrance) render(viewCanvas *pixelgl.Canvas) {
	if e.renderFloatingUI {
		e.floatingUI.render(viewCanvas)
	}
}

func (e *Entrance) playerCollision(p *Player) {
	if p.pos.X < e.pos.X+e.size.X &&
		p.pos.X+p.size.X > e.pos.X &&
		p.pos.Y < e.pos.Y+e.size.Y &&
		p.pos.Y+p.size.Y > e.pos.Y {
		e.renderFloatingUI = true
	} else {
		e.renderFloatingUI = false
	}
}
