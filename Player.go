package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*Player ... struct for controllable players in the game.*/
type Player struct {
	pos                pixel.Vec
	center             pixel.Vec
	velocity           pixel.Vec
	maxSpeed           float64
	currSpeed          float64
	rotation           float64
	radius             float64
	size               pixel.Vec
	currDir            int // Current direction of moving, 0 W, 1 D, 2 S, 3 A
	canMove            bool
	activeMovement     bool
	pic                pixel.Picture
	health             int8
	maxHealth          int8
	animation          Animation
	batch              *pixel.Batch
	footSizeDiminisher float64 // Diminisher for where the feet are for collisions

	// Animations
	animations PlayerAnimations
}

// PlayerAnimations ... Player animations in the game
type PlayerAnimations struct { // Holds all the animations for the player
	idleRightAnimation Animation
	idleUpAnimation    Animation
	idleDownAnimation  Animation
	idleLeftAnimation  Animation
}

func createPlayer(pos pixel.Vec, cID int, pic pixel.Picture, movable bool) Player { // Player constructor
	size := pixel.V(pic.Bounds().Size().X/float64(len(playerSpritesheets.playerIdleRightSheet.frames)), pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)

	idleAnimationSpeed := 0.3

	return Player{
		pos,
		pixel.ZV,
		pixel.ZV,
		20.0,
		35.0,
		0.0,
		pic.Bounds().Size().Y / 2,
		size,
		1,
		movable,
		false,
		pic,
		100,
		100,
		createAnimation(playerSpritesheets.playerIdleRightSheet, idleAnimationSpeed),
		playerBatches.playerIdleRightBatch,
		10.,
		PlayerAnimations{
			createAnimation(playerSpritesheets.playerIdleRightSheet, idleAnimationSpeed),
			createAnimation(playerSpritesheets.playerIdleUpSheet, idleAnimationSpeed),
			createAnimation(playerSpritesheets.playerIdleDownSheet, idleAnimationSpeed),
			createAnimation(playerSpritesheets.playerIdleLeftSheet, idleAnimationSpeed),
		},
	}

}

func (p *Player) update(win *pixelgl.Window, dt float64) { // Updates player
	if p.canMove {
		p.input(win, dt)
	}
	p.center = pixel.V(p.pos.X+(p.size.X/2), p.pos.Y+(p.size.Y/2))
}

func (p *Player) render(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, dt float64) { // Draws the player
	p.batch.Clear()
	sprite := p.animation.animate(dt)
	sprite.Draw(p.batch, pixel.IM.Rotated(pixel.ZV, p.rotation).Moved(p.center).Scaled(p.center, imageScale))
	p.batch.Draw(viewCanvas)
}

func (p *Player) input(win *pixelgl.Window, dt float64) {
	if p.velocity.X > p.maxSpeed {
		p.velocity.X = p.maxSpeed
	} else if p.velocity.X < -1*p.maxSpeed {
		p.velocity.X = -1 * p.maxSpeed
	}
	if p.velocity.Y > p.maxSpeed {
		p.velocity.Y = p.maxSpeed
	} else if p.velocity.Y < -1*p.maxSpeed {
		p.velocity.Y = -1 * p.maxSpeed
	}

	if p.canMove {
		p.velocity = pixel.V(0., 0.)
	}

	if win.Pressed(pixelgl.KeyW) && win.Pressed(pixelgl.KeyD) {
		if p.currDir != 0 {
			p.currDir = 0
			p.batch = playerBatches.playerIdleUpBatch
			p.animation = p.animations.idleUpAnimation
		}
		p.velocity.Y = p.currSpeed
		p.velocity.X = p.currSpeed
	} else if win.Pressed(pixelgl.KeyW) && win.Pressed(pixelgl.KeyA) {
		if p.currDir != 0 {
			p.currDir = 0
			p.batch = playerBatches.playerIdleUpBatch
			p.animation = p.animations.idleUpAnimation
		}
		p.velocity.Y = p.currSpeed
		p.velocity.X = -p.currSpeed
	} else if win.Pressed(pixelgl.KeyS) && win.Pressed(pixelgl.KeyD) {
		if p.currDir != 2 {
			p.currDir = 2
			p.batch = playerBatches.playerIdleDownBatch
			p.animation = p.animations.idleDownAnimation
		}
		p.velocity.Y = -p.currSpeed
		p.velocity.X = p.currSpeed
	} else if win.Pressed(pixelgl.KeyS) && win.Pressed(pixelgl.KeyA) {
		if p.currDir != 2 {
			p.currDir = 2
			p.batch = playerBatches.playerIdleDownBatch
			p.animation = p.animations.idleDownAnimation
		}
		p.velocity.Y = -p.currSpeed
		p.velocity.X = -p.currSpeed
	} else {
		if win.Pressed(pixelgl.KeyW) { // Up, 0
			if p.currDir != 0 {
				p.currDir = 0
				p.batch = playerBatches.playerIdleUpBatch
				p.animation = p.animations.idleUpAnimation
			}
			p.velocity.Y = p.currSpeed
		}
		if win.Pressed(pixelgl.KeyD) { // Right, 1
			if p.currDir != 1 {
				p.currDir = 1
				p.batch = playerBatches.playerIdleRightBatch
				p.animation = p.animations.idleRightAnimation
			}
			p.velocity.X = p.currSpeed
		}
		if win.Pressed(pixelgl.KeyS) { // Down, 2
			if p.currDir != 2 {
				p.currDir = 2
				p.batch = playerBatches.playerIdleDownBatch
				p.animation = p.animations.idleDownAnimation
			}
			p.velocity.Y = p.currSpeed

			p.velocity.Y = -p.currSpeed
		}
		if win.Pressed(pixelgl.KeyA) { // Left, 3
			if p.currDir != 3 {
				p.currDir = 3
				p.batch = playerBatches.playerIdleLeftBatch
				p.animation = p.animations.idleLeftAnimation
			}

			p.velocity.X = -p.currSpeed
		}
	}

	p.pos = pixel.V(p.pos.X+p.velocity.X*dt, p.pos.Y+p.velocity.Y*dt)

	p.isMoving(win)
}

func (p *Player) isMoving(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyD) {
		p.activeMovement = true
	} else {
		p.activeMovement = false
	}
}
