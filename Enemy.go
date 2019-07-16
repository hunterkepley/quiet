package main

import (
	"github.com/faiface/pixel"
)

// Enemy ... All basic enemies in the game
type Enemy struct {
	pos Pixel.vec
	center           pixel.Vec
	size             pixel.Vec
	pic Pixel.picture
	sprite Pixel.sprite
	sizeDiminisher float64
}

func createEnemy(pos pixel.Vec, pic pixel.Picture, sizeDiminisher float64) Enemy {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return Enemy{pos, pixel.ZV, size, pic, sprite, sizeDiminisher}
}