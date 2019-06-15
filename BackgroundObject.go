package main

import "github.com/faiface/pixel"

type BackgroundObject struct {
	pos    pixel.Vec
	center pixel.Vec
	size   pixel.Vec
	pic    pixel.Picture
	sprite pixel.Sprite
	radius float64
}
