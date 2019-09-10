package main

import (
	"github.com/faiface/pixel"
	//"github.com/faiface/pixel/pixelgl"
)

/**
 * This file is for entrances, like a door to another room. They are just zones defined, no image attached.
 */

// Entrance ... An entrace to a new room
type Entrance struct {
	pos        pixel.Vec
	size       pixel.Vec
	floatingUI FloatingUI
	roomIndex  int // What room to switch to in the current level, if -1, changes level instead
	levelIndex int // What level to switch to, if -1, changes room instead
}
