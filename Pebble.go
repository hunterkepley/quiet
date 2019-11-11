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
	maxSpeedBase    float64
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

// And thus, throwable objects became one of the most filled and complicated structs in the game

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

	// This shit got a lil complicated a lil quickly :)
	// Lots of badly done physics that work somehow,
	// ~~please don't look~~
	switch p.direction {
	case (0):
		p.maxSpeed -= 140 * dt
		if p.maxSpeed > 0 {
			p.pos.Y += p.maxSpeedBase * dt
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		if p.maxSpeed <= 0 {
			p.maxSpeedBase = 0
		}
		break
	case (1):
		sin := math.Abs(math.Sin((p.startPos.X-p.pos.X)/80)) * 400 / p.diminisher
		if sin <= 2 {
			p.diminisher += 200 * dt
		}
		p.gravityVelocity.X = p.maxSpeed * p.velocity.X
		p.gravityVelocity.Y = sin
		if p.maxSpeed > 0 {
			p.pos.X += (p.velocity.X * p.gravityVelocity.X) * dt
			p.pos.Y = p.startPos.Y + p.gravityVelocity.Y
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		if p.maxSpeed <= 146 && sin <= 2 {
			p.maxSpeed = 0
		}
		break
	case (2):
		p.maxSpeed -= 140 * dt
		if p.maxSpeed > 0 {
			p.pos.Y -= p.maxSpeedBase * dt
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		if p.maxSpeed <= 0 {
			p.maxSpeedBase = 0
		}
		break
	case (3):
		sin := math.Abs(math.Sin((p.startPos.X-p.pos.X)/80)) * 400 / p.diminisher
		if sin <= 2 {
			p.diminisher += 200 * dt
		}
		p.gravityVelocity.X = p.maxSpeed * p.velocity.X
		p.gravityVelocity.Y = sin
		if p.maxSpeed > 0 {
			p.pos.X -= (p.velocity.X * p.gravityVelocity.X) * dt
			p.pos.Y = p.startPos.Y + p.gravityVelocity.Y
			p.center = pixel.V(p.pos.X+p.size.X/2, p.pos.Y+p.size.Y/2)
			p.maxSpeed -= p.diminisher * dt
		}
		if p.maxSpeed <= 146 && sin <= 2 {
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
