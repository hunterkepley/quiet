package main

import (
	"github.com/faiface/pixel"
)

//Pebble ... Throwable objects in game that make sound
type Pebble struct {
	pos      pixel.Vec
	size     pixel.Vec
	pic      pixel.Picture
	sprite   *pixel.Sprite
	velocity pixel.Vec
	maxSpeed float64

	// Sound emitter
	activateSoundEmitter bool
	allowSoundEmitter    bool
	soundEmitter         SoundEmitter
	soundTimer           float64
	soundTimerMax        float64
	soundDB              float64 // The starting dB of sound currently
}

func createPebble(startPos pixel.Vec, pic pixel.Picture, maxSpeed float64, soundDB float64) Pebble {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return Pebble{
		startPos,
		size,
		pic,
		sprite,
		pixel.V(0, 0),
		maxSpeed,
		// Sound emitter
		false,
		true,
		createSoundEmitter(startPos),
		1.,
		1.,
		soundDB,
	}
}
