package main

import (
	"github.com/faiface/pixel"
)

// Animation ... Animations........
type Animation struct {
	sheet          Spritesheet
	currentSprite  *pixel.Sprite
	frameSpeed     float64
	frameSpeedMax  float64
	frameNumber    int
	frameNumberMax int
}

func createAnimation(sheet Spritesheet, frameSpeedMax float64) Animation {
	frameSpeed := 0.
	frameNumber, frameNumberMax := 0, sheet.numberOfFrames
	sprite := pixel.NewSprite(sheet.sheet, sheet.frames[frameNumber])
	return Animation{sheet, sprite, frameSpeed, frameSpeedMax, frameNumber, frameNumberMax}
}

func (a *Animation) animate(dt float64) pixel.Sprite {
	if a.frameSpeed > a.frameSpeedMax {
		a.frameNumber++
		if a.frameNumber >= a.frameNumberMax {
			a.frameNumber = 0
		}
		a.frameSpeed = 0.
	} else {
		a.frameSpeed += 1. * dt
		a.currentSprite = pixel.NewSprite(a.sheet.sheet, a.sheet.frames[a.frameNumber])
	}

	return *a.currentSprite
}
