package main

import (
	"github.com/faiface/pixel"
)

// Enemy ... All basic enemies in the game
type Enemy struct {
	pos Pixel.vec
	pic Pixel.picture
	sprite Pixel.sprite
}