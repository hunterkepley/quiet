package main

import (
	"fmt"

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
	blink              bool // Blink when damage is taken
	blinkRender        bool // When to render while blinking
	blinkTimer         float64
	blinkTimerMax      float64 // How long each blink lasts
	blinkAmount        int
	blinkAmountMax     int // How many times to blink
	knockbackTimer     float64
	knockbackTimerMax  float64 // How long to be knocked back for
	pic                pixel.Picture
	health             int
	maxHealth          int
	animation          Animation
	batch              *pixel.Batch
	footSizeDiminisher float64 // Diminisher for where the feet are for collisions
	imageScale         float64
	hitBox             pixel.Rect
	footHitBox         pixel.Rect
	// Throwables
	throwables       []Pebble
	throwCooldown    float64
	throwCooldownMax float64
	// Sound emitter
	activateSoundEmitter bool
	allowSoundEmitter    bool
	soundEmitter         SoundEmitter
	soundTimer           float64
	soundTimerMax        float64
	soundDB              float64 // The starting dB of sound currently

	// Animations
	idleAnimationSpeed float64
	moveAnimationSpeed float64
	animations         PlayerAnimations
}

//PlayerAnimations ... Player animations in the game
type PlayerAnimations struct { // Holds all the animations for the player
	idleRightAnimation Animation
	idleUpAnimation    Animation
	idleDownAnimation  Animation
	idleLeftAnimation  Animation
}

func createPlayer(pos pixel.Vec, cID int, pic pixel.Picture, movable bool, playerImageScale float64) Player { // Player constructor
	size := pixel.V(pic.Bounds().Size().X/float64(len(playerSpritesheets.playerIdleRightSheet.frames)), pic.Bounds().Size().Y)
	size = pixel.V(size.X*playerImageScale, size.Y*playerImageScale)

	idleAnimationSpeed := 0.6
	moveAnimationSpeed := 0.15

	soundDB := 55.

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
		false,
		false,
		0.5,
		0.5,
		3,
		3,
		0,
		0.5,
		pic,
		100,
		100,
		createAnimation(playerSpritesheets.playerIdleRightSheet, idleAnimationSpeed),
		playerBatches.playerIdleRightBatch,
		10.,
		playerImageScale,
		pixel.R(0, 0, 0, 0),
		pixel.R(0, 0, 0, 0),
		[]Pebble{},
		0,
		5,
		true,
		true,
		createSoundEmitter(pos),
		1.,
		1.,
		soundDB,
		idleAnimationSpeed,
		moveAnimationSpeed,
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

	for i := range p.throwables {
		p.throwables[i].update(dt)
	}

	p.updateHitboxes()

	if p.activeMovement {
		p.activateSoundEmitter = true
		p.animation.frameSpeedMax = p.moveAnimationSpeed
	} else {
		p.activateSoundEmitter = false
		p.animation.frameSpeedMax = p.idleAnimationSpeed
	}

	// Update sound emitter
	if p.allowSoundEmitter { // If the sound emitter is allowed in a room
		p.soundEmitter.update(pixel.V(p.footHitBox.Min.X+p.size.X/2, p.pos.Y), dt)
		if p.soundTimer > 0 { // Constantly tick down the timer to prevent tapping a key to avoid sound emitting, won't emit until moving
			p.soundTimer -= 1 * dt
		}
		if p.activateSoundEmitter { // If the player is currently walking
			if p.soundTimer < 0 {
				p.soundEmitter.emit(p.soundDB, 10)
				p.soundTimer = p.soundTimerMax
			}
		}
	}

	if p.knockbackTimer > 0 { // Knockback!
		p.canMove = false
		p.knockbackTimer -= 1 * dt
		fmt.Println(p.knockbackTimer)
	} else {
		p.canMove = true
	}

	// Screen edge collision detection/response
	if p.center.X-p.size.X/2 < 0. || p.center.X+p.size.X/2 > winWidth { // Left / Right
		p.pos.X += (p.velocity.X * -1) * dt
	}
	if p.center.Y-p.size.Y/2 < 0. || p.center.Y+p.size.Y/2 > winHeight { // Bottom / Top
		p.pos.Y += (p.velocity.Y * -1) * dt
	}
}

func (p *Player) updateHitboxes() { // Also updates size
	if p.currDir == 1 || p.currDir == 3 {
		p.size = pixel.V(playerSpritesheets.playerIdleRightSheet.sheet.Bounds().Size().X/float64(len(playerSpritesheets.playerIdleRightSheet.frames)), p.pic.Bounds().Size().Y)
	} else {
		p.size = pixel.V(playerSpritesheets.playerIdleUpSheet.sheet.Bounds().Size().X/float64(len(playerSpritesheets.playerIdleUpSheet.frames)), p.pic.Bounds().Size().Y)
	}
	p.size = pixel.V(p.size.X*p.imageScale, p.size.Y*p.imageScale)
	p.footHitBox = pixel.R(p.pos.X, p.pos.Y, p.pos.X+p.size.X, p.pos.Y+p.size.Y/p.footSizeDiminisher)
	p.hitBox = pixel.R(p.pos.X, p.pos.Y, p.pos.X+p.size.X, p.pos.Y+p.size.Y)
}

func (p *Player) render(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, dt float64) { // Draws the player
	p.batch.Clear()
	for i := range p.throwables {
		p.throwables[i].render(viewCanvas)
	}
	sprite := p.animation.animate(dt)
	if p.blink { // Blink from damage
		if p.blinkTimer > 0 {
			p.blinkTimer -= 1. * dt
		} else {
			p.blinkTimer = p.blinkTimerMax
			p.blinkAmount--
		}
		if p.blinkTimer > p.blinkTimerMax/2. {
			p.blinkRender = true
		} else {
			p.blinkRender = false
		}
		if p.blinkAmount < 0 {
			p.blinkAmount = p.blinkAmountMax
			p.blinkTimer = p.blinkTimerMax
			p.blinkRender = false
			p.blink = false
		}
	}
	if !p.blinkRender {
		sprite.Draw(p.batch, pixel.IM.Rotated(pixel.ZV, p.rotation).Moved(p.center).Scaled(p.center, p.imageScale))
		p.batch.Draw(viewCanvas)
	}
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

	if p.throwCooldown > 0 {
		p.throwCooldown -= 1 * dt
	}
	if win.Pressed(pixelgl.KeyQ) /*&& p.throwCooldown <= 0*/ {
		p.throwables = append(p.throwables, createPebble(pixel.V(p.center.X, p.pos.Y), throwablesImages.pebble, 150., 15., p.currDir, p.pos.Y))
		p.throwCooldown = p.throwCooldownMax
	}

	if win.Pressed(pixelgl.KeyW) && win.Pressed(pixelgl.KeyD) {
		if p.currDir != 1 {
			p.currDir = 1
			p.batch = playerBatches.playerIdleRightBatch
			p.animation = p.animations.idleRightAnimation
		}
		p.velocity.Y = p.currSpeed
		p.velocity.X = p.currSpeed
	} else if win.Pressed(pixelgl.KeyW) && win.Pressed(pixelgl.KeyA) {
		if p.currDir != 3 {
			p.currDir = 3
			p.batch = playerBatches.playerIdleLeftBatch
			p.animation = p.animations.idleLeftAnimation
		}
		p.velocity.Y = p.currSpeed
		p.velocity.X = -p.currSpeed
	} else if win.Pressed(pixelgl.KeyS) && win.Pressed(pixelgl.KeyD) {
		if p.currDir != 1 {
			p.currDir = 1
			p.batch = playerBatches.playerIdleRightBatch
			p.animation = p.animations.idleRightAnimation
		}
		p.velocity.Y = -p.currSpeed
		p.velocity.X = p.currSpeed
	} else if win.Pressed(pixelgl.KeyS) && win.Pressed(pixelgl.KeyA) {
		if p.currDir != 3 {
			p.currDir = 3
			p.batch = playerBatches.playerIdleLeftBatch
			p.animation = p.animations.idleLeftAnimation
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

	if p.canMove {
		p.pos = pixel.V(p.pos.X+p.velocity.X*dt, p.pos.Y+p.velocity.Y*dt)
	}

	p.isMoving(win)
}

func (p *Player) isMoving(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyD) {
		p.activeMovement = true
	} else {
		p.activeMovement = false
	}
}

func (p *Player) takeDamage(damage int, knockbackTimer float64) {
	if !p.blink {
		p.blink = true
		p.knockbackTimerMax = knockbackTimer
		p.knockbackTimer = p.knockbackTimerMax
	}
	p.health -= damage
}
