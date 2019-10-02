package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/**
 * This file is for things like the floating E above a door to enter a new room
 */

// FloatingUI ... Images to be displayed as UI, but float in the game, like an E above a door
type FloatingUI struct {
	pos                 pixel.Vec
	size                pixel.Vec
	center              pixel.Vec
	sprite              *pixel.Sprite
	pic                 pixel.Picture
	bounceRange         float64 // How far up and down does the image bounce from it's orgin?
	currentBounceOffset float64 // How far is the UI currently bouncing?
	bouncingUp          bool    // Is it bouncing up or down?
}

func createFloatingUI(pos pixel.Vec, pic pixel.Picture, bounceRange float64) FloatingUI {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	center := pixel.V(pos.X+(size.X/2), pos.Y+(size.Y/2))
	return FloatingUI{
		pos,
		size,
		center,
		sprite,
		pic,
		bounceRange,
		0.,
		true,
	}
}

func (f *FloatingUI) render(viewCanvas *pixelgl.Canvas) {
	combined := pixel.V(f.center.X, f.center.Y+f.currentBounceOffset) // For the bounce offset
	mat := pixel.IM.
		Moved(combined).
		Scaled(f.center, imageScale)
	f.sprite.Draw(viewCanvas, mat)
}

func (f *FloatingUI) update(dt float64, entranceWidth float64) {
	sub := f.size.X - entranceWidth
	if sub < 0 {
		sub = entranceWidth - f.size.X
	}

	f.center = pixel.V(f.pos.X+(f.size.X/2)-sub, f.pos.Y+(f.size.Y/2))
	if f.bouncingUp {
		if f.currentBounceOffset < f.bounceRange {
			f.currentBounceOffset += 5. * dt
		} else {
			f.bouncingUp = false
		}
	} else {
		if f.currentBounceOffset > (0) {
			f.currentBounceOffset -= 5. * dt
		} else {
			f.bouncingUp = true
		}
	}
}
