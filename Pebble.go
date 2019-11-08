package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//Pebble ... Throwable objects in game that make sound
type Pebble struct {
	pos             pixel.Vec
	startPos        pixel.Vec
	size            pixel.Vec
	center          pixel.Vec
	pic             pixel.Picture
	sprite          *pixel.Sprite
	velocity        pixel.Vec
	maxSpeed        float64
	declinationRate float64
	diminisher      float64
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
		startPos.Y = player.center.Y
		break
	case (1):
		velocity = pixel.V(1, 1)
		break
	case (2):
		velocity = pixel.V(1, -1)
		startPos.Y = player.center.Y - player.size.Y/2
		break
	case (3):
		velocity = pixel.V(-1, 1)
	}
	return Pebble{
		startPos,
		startPos,
		size,
		pixel.V(0, 0),
		pic,
		sprite,
		velocity,
		maxSpeed,
		8, // Decination rate
		1,
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

	switch p.direction {
	case (0):
		sin := math.Abs(math.Sin((p.startPos.X-p.pos.X)/15)) * 100 / p.diminisher
		if sin <= 0 {
			p.diminisher += 0.5
		}
		p.gravityVelocity.X = p.maxSpeed * p.velocity.X
		p.gravityVelocity.Y = sin
		if p.maxSpeed > 0 {
			p.pos.Y += p.gravityVelocity.X * p.velocity.X * dt
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		break
	case (1):
		sin := math.Abs(math.Sin((p.startPos.X-p.pos.X)/15)) * 100 / p.diminisher
		if sin <= 2 {
			p.diminisher++
		}
		p.gravityVelocity.X = p.maxSpeed * p.velocity.X
		p.gravityVelocity.Y = sin
		if p.maxSpeed > 0 {
			p.pos.X += (p.velocity.X * p.gravityVelocity.X) * dt
			p.pos.Y = p.startPos.Y + p.gravityVelocity.Y
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		if p.maxSpeed <= 45 && sin <= 2 {
			p.maxSpeed = 0
		}
		break
	case (2):
		sin := math.Abs(math.Sin((p.startPos.X-p.pos.X)/15)) * 100 / p.diminisher
		if sin <= 0 {
			p.diminisher += 0.5
		}
		p.gravityVelocity.X = p.maxSpeed * p.velocity.X
		p.gravityVelocity.Y = sin
		if p.maxSpeed > 0 {
			p.pos.Y -= p.gravityVelocity.X * p.velocity.X * dt
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		break
	case (3):
		sin := math.Abs(math.Sin((p.startPos.X-p.pos.X)/15)) * 100 / p.diminisher
		if sin <= 2 {
			p.diminisher++
		}
		p.gravityVelocity.X = p.maxSpeed * p.velocity.X
		p.gravityVelocity.Y = sin
		if p.maxSpeed > 0 {
			p.pos.X -= (p.velocity.X * p.gravityVelocity.X) * dt
			p.pos.Y = p.startPos.Y + p.gravityVelocity.Y
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		if p.maxSpeed <= 45 && sin <= 2 {
			p.maxSpeed = 0
		}
		break
	}
}

func (p *Pebble) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(p.center).
		Scaled(p.center, imageScale)
	p.sprite.Draw(viewCanvas, mat)
}
