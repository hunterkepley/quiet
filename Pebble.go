package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//Pebble ... Throwable objects in game that make sound
type Pebble struct {
	pos      pixel.Vec
	size     pixel.Vec
	center   pixel.Vec
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
		pixel.V(0, 0),
		pic,
		sprite,
		pixel.V(1, 1),
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

func (p *Pebble) update(dt float64) {
	p.pos.X += (p.velocity.X * p.maxSpeed) * dt
	p.pos.Y += (p.velocity.Y * p.maxSpeed) * dt
	p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
}

func (p *Pebble) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(p.center).
		Scaled(p.center, imageScale)
	p.sprite.Draw(viewCanvas, mat)
}
