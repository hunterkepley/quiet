package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// Enemy ... All basic enemies in the game
type Enemy struct {
	pos             pixel.Vec
	center          pixel.Vec
	size            pixel.Vec
	pic             pixel.Picture
	sprite          *pixel.Sprite
	sizeDiminisher  float64
	moveSpeed       float64
	moveVector      pixel.Vec // 1, 1 for moving top right, 0, 1 for moving up, etc.
	noSoundTimer    float64   // Timer for how long until they stop chasing after not hearing a sound
	noSoundTimerMax float64
	targetPosition  pixel.Vec // The position the enemy will go to

	// Animations
	animation  Animation
	animations EnemyAnimations
}

//EnemyAnimations .. Enemy animations in the game
type EnemyAnimations struct {
	leftAnimation Animation
}

func createEnemy(pos pixel.Vec, pic pixel.Picture, sizeDiminisher float64, moveSpeed float64, noSoundTimer float64) Enemy {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)

	animationSpeed := 0.2
	return Enemy{
		pos,
		pixel.ZV,
		size,
		pic,
		sprite,
		sizeDiminisher,
		moveSpeed,
		pixel.ZV,
		0.,
		noSoundTimer,
		pixel.ZV,
		createAnimation(enemySpriteSheets.larvaSpriteSheets.leftSpriteSheet, animationSpeed),
		EnemyAnimations{
			createAnimation(enemySpriteSheets.larvaSpriteSheets.leftSpriteSheet, animationSpeed),
		},
	}
}

func (e *Enemy) render(viewCanvas *pixelgl.Canvas, imd *imdraw.IMDraw) {
	mat := pixel.IM.
		Moved(e.center).
		Scaled(e.center, imageScale)
	sprite := e.animation.animate(dt)
	sprite.Draw(viewCanvas, mat)
}

func (e *Enemy) update(dt float64, soundWaves []SoundWave) {
	e.moveVector = pixel.V(0, 0)
	for i := 0; i < len(soundWaves); i++ {
		if soundWaves[i].pos.X < e.pos.X+e.size.X &&
			soundWaves[i].pos.X+soundWaves[i].size.X > e.pos.X &&
			soundWaves[i].pos.Y < e.pos.Y+e.size.Y/e.sizeDiminisher &&
			soundWaves[i].pos.Y+soundWaves[i].size.Y > e.pos.Y {
			e.noSoundTimer = e.noSoundTimerMax
			e.targetPosition = soundWaves[i].startPos
		}
	}
	if e.noSoundTimer > 0. {
		if e.targetPosition.X > e.center.X {
			e.moveVector.X = 1
		} else if e.targetPosition.X < e.center.X {
			e.moveVector.X = -1
		}
		if e.targetPosition.Y > e.center.Y {
			e.moveVector.Y = 1
		} else if e.targetPosition.Y < e.center.Y {
			e.moveVector.Y = -1
		}
		e.noSoundTimer -= 1 * dt
	}
	e.pos = pixel.V(e.pos.X+(e.moveSpeed*dt)*e.moveVector.X, e.pos.Y+(e.moveSpeed*dt)*e.moveVector.Y)
	e.center = pixel.V(e.pos.X+(e.size.X/2), e.pos.Y+(e.size.Y/2))
}
