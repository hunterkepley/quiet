package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//Pebble ... Throwable objects in game that make sound
type Pebble struct {
	pos             pixel.Vec
	size            pixel.Vec
	center          pixel.Vec
	pic             pixel.Picture
	sprite          *pixel.Sprite
	velocity        pixel.Vec
	maxSpeed        float64
	declinationRate float64
	direction       int
	groundLevel     float64
	gravityVelocity pixel.Vec
	bounce          bool

	// Sound emitter
	activateSoundEmitter bool
	allowSoundEmitter    bool
	soundEmitter         SoundEmitter
	soundTimer           float64
	soundTimerMax        float64
	soundDB              float64 // The starting dB of sound currently
}

func createPebble(startPos pixel.Vec, pic pixel.Picture, maxSpeed float64, soundDB float64, direction int, groundLevel float64) Pebble {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	velocity := pixel.V(1, 1)
	switch direction {
	case (0):
		velocity = pixel.V(1, 1)
		break
	case (1):
		velocity = pixel.V(1, 1)
		break
	case (2):
		velocity = pixel.V(1, -1)
		break
	case (3):
		velocity = pixel.V(-1, 1)
	}
	return Pebble{
		startPos,
		size,
		pixel.V(0, 0),
		pic,
		sprite,
		velocity,
		maxSpeed,
		8,             // Decination rate
		direction,     // The direction in which the object in thrown
		groundLevel,   // The level at which the object touches the ground
		pixel.V(0, 0), // Velocity at which gravity affects the object
		true,          // True when the object bounces up, false when falling
		//            0 Up, 1 Right, 2 Down, 3 Left
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

	bounceMultiplier := 4.
	gravityVelocityMax := 50.

	switch p.direction {
	case (0):
		if p.bounce {
			if p.gravityVelocity.Y < 10 {
				p.gravityVelocity.Y += p.maxSpeed * bounceMultiplier * dt
			} else {
				p.bounce = false
			}
		} else {
			if p.gravityVelocity.Y > 0 {
				p.gravityVelocity.Y -= p.maxSpeed * bounceMultiplier * dt
			} else {
				p.bounce = true
			}
		}
		break
	case (1):
		p.gravityVelocity.X = p.maxSpeed
		if p.bounce {
			if p.gravityVelocity.Y < gravityVelocityMax {
				p.gravityVelocity.Y += p.maxSpeed * bounceMultiplier * dt
			} else {
				p.bounce = false
			}
		} else {
			if p.gravityVelocity.Y > -gravityVelocityMax {
				p.gravityVelocity.Y -= p.maxSpeed * bounceMultiplier * dt
			} else {
				p.bounce = true
				gravityVelocityMax /= 2
			}
		}
		break
	}
	if p.maxSpeed > 0 {
		p.pos.X += (p.velocity.X * p.gravityVelocity.X) * dt
		p.pos.Y += (p.velocity.Y * p.gravityVelocity.Y) * dt
		p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
		p.maxSpeed -= p.declinationRate * dt
	}
	if p.maxSpeed <= 0 {
		p.maxSpeed = 0
	}
}

func (p *Pebble) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(p.center).
		Scaled(p.center, imageScale)
	p.sprite.Draw(viewCanvas, mat)
}
