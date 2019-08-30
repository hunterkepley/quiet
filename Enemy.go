package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

//Eye ... The eye above enemies heads
type Eye struct {
	pos    pixel.Vec
	center pixel.Vec
	size   pixel.Vec
	sprite *pixel.Sprite
	state  int // 0 is looking, 1 is opening, 2 is closed [still image], 3 is closing

	// Animations
	animation  Animation
	animations EyeAnimations
}

//EyeAnimations ... Eye animations in the game
type EyeAnimations struct {
	lookingAnimation Animation
	openingAnimation Animation
	closingAnimation Animation
}

//Enemy ... All basic enemies in the game
type Enemy struct {
	pos                pixel.Vec
	center             pixel.Vec
	size               pixel.Vec
	pic                pixel.Picture
	sprite             *pixel.Sprite
	sizeDiminisher     float64
	moveSpeed          float64
	moveVector         pixel.Vec // 1, 1 for moving top right, 0, 1 for moving up, etc.
	noSoundTimer       float64   // Timer for how long until they stop chasing after not hearing a sound
	noSoundTimerMax    float64
	targetPosition     pixel.Vec // The position the enemy will go to
	currentAnimation   int       // Int of the current animation. 0 = top, 3 = left
	moveAnimationSpeed float64
	idleAnimationSpeed float64
	canAttack          bool
	attacking          bool
	attackCooldown     float64
	attackCooldownMax  float64
	attackCheckRadius  float64
	eye                Eye

	// Nodes
	nodes       []Node
	currentPath []Node

	// Animations
	animation  Animation
	animations EnemyAnimations
}

//EnemyAnimations .. Enemy animations in the game
type EnemyAnimations struct {
	leftAnimation        Animation
	rightAnimation       Animation
	attackRaiseAnimation Animation
}

func createEnemy(pos pixel.Vec, pic pixel.Picture, sizeDiminisher float64, moveSpeed float64, noSoundTimer float64, moveAnimationSpeed float64, idleAnimationSpeed float64, attackCooldown float64, attackCheckRadius float64) Enemy {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	eyeLookingAnimationSpeed := 0.1
	eyeOpeningAnimationSpeed := 0.1
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
		3,
		moveAnimationSpeed,
		idleAnimationSpeed,
		false,
		false,
		attackCooldown,
		attackCooldown,
		attackCheckRadius,
		Eye{
			pixel.ZV,
			pixel.ZV,
			pixel.V(enemyImages.eye.Bounds().Size().X*imageScale, enemyImages.eye.Bounds().Size().Y*imageScale),
			pixel.NewSprite(enemyImages.eye, enemyImages.eye.Bounds()),
			2,
			createAnimation(enemySpriteSheets.eyeLookingSheet, eyeLookingAnimationSpeed),
			EyeAnimations{
				createAnimation(enemySpriteSheets.eyeLookingSheet, eyeLookingAnimationSpeed),
				createAnimation(enemySpriteSheets.eyeOpeningSheet, eyeOpeningAnimationSpeed),
				createAnimation(enemySpriteSheets.eyeClosingSheet, eyeOpeningAnimationSpeed),
			},
		},
		[]Node{},
		[]Node{},
		createAnimation(enemySpriteSheets.larvaSpriteSheets.leftSpriteSheet, idleAnimationSpeed),
		EnemyAnimations{
			createAnimation(enemySpriteSheets.larvaSpriteSheets.leftSpriteSheet, idleAnimationSpeed),
			createAnimation(enemySpriteSheets.larvaSpriteSheets.rightSpriteSheet, idleAnimationSpeed),
			createAnimation(enemySpriteSheets.larvaSpriteSheets.attackRaiseSpriteSheet, idleAnimationSpeed),
		},
	}
}

func (e *Enemy) render(viewCanvas *pixelgl.Canvas, imd *imdraw.IMDraw) {
	mat := pixel.IM.
		Moved(e.center).
		Scaled(e.center, imageScale)
	if e.eye.state != 2 {
		*e.eye.sprite = e.eye.animation.animate(dt)
	} else {
		e.eye.sprite = pixel.NewSprite(enemyImages.eye, enemyImages.eye.Bounds())
	}

	sprite := e.animation.animate(dt)
	sprite.Draw(viewCanvas, mat)
	// Render nodes, temporary
	for _, j := range e.currentPath {
		j.render(imd)
	}
}

func (e *Enemy) eyeRender(viewCanvas *pixelgl.Canvas) {
	eyeMat := pixel.IM.
		Moved(e.eye.center).
		Scaled(e.eye.center, imageScale)
	e.eye.sprite.Draw(viewCanvas, eyeMat)
}

func (e *Enemy) update(dt float64, soundWaves []SoundWave, p *Player) {
	e.moveVector = pixel.V(0, 0)
	if e.noSoundTimer <= 0. {
		if e.eye.state != 2 { // Close eye
			if e.eye.state != 3 {
				e.eye.state = 3
				e.eye.animation = e.eye.animations.closingAnimation
			}
			if e.eye.animation.frameNumber >= e.eye.animation.frameNumberMax-1 {
				e.eye.state = 2
			}
			e.canAttack = false
		}
		for i := 0; i < len(soundWaves); i++ {
			if soundWaves[i].pos.X < e.pos.X+e.size.X &&
				soundWaves[i].pos.X+soundWaves[i].size.X > e.pos.X &&
				soundWaves[i].pos.Y < e.pos.Y+e.size.Y/e.sizeDiminisher &&
				soundWaves[i].pos.Y+soundWaves[i].size.Y > e.pos.Y {
				nodeIndexStart := 0
				nodeIndexEnd := 0
				for nI, n := range e.nodes {
					if n.pos.X < e.pos.X+1 &&
						n.pos.X+n.size.X > e.pos.X &&
						n.pos.Y < e.pos.Y+1 &&
						n.pos.Y+n.size.Y > e.pos.Y {
						if n.passable {
							nodeIndexStart = nI
						}
					} else if n.pos.X < soundWaves[i].startPos.X+1 &&
						n.pos.X+n.size.X > soundWaves[i].startPos.X &&
						n.pos.Y < soundWaves[i].startPos.Y+1 &&
						n.pos.Y+n.size.Y > soundWaves[i].startPos.Y {
						if n.passable {
							nodeIndexEnd = nI
						}
					}
				}
				e.currentPath = astar(nodeIndexStart, nodeIndexEnd, e.nodes, e.size)

				e.noSoundTimer = e.noSoundTimerMax
				e.eye.state = 1
				e.eye.animation = e.eye.animations.openingAnimation
				soundWaves[i].dB = 0. // Destroy the wave to show it hit the enemy
			}
		}
	}
	e.animation.frameSpeedMax = e.idleAnimationSpeed
	if e.noSoundTimer > 0. {
		for i := 0; i < len(soundWaves); i++ {
			if soundWaves[i].pos.X < e.pos.X+e.size.X &&
				soundWaves[i].pos.X+soundWaves[i].size.X > e.pos.X &&
				soundWaves[i].pos.Y < e.pos.Y+e.size.Y/e.sizeDiminisher &&
				soundWaves[i].pos.Y+soundWaves[i].size.Y > e.pos.Y {
				soundWaves[i].dB = 0. // Destroy the wave to show it hit the enemy
				e.noSoundTimer = e.noSoundTimerMax
				nodeIndexStart := 0
				nodeIndexEnd := 0
				if len(e.currentPath) <= 0 {
					for nI, n := range e.nodes {
						if n.pos.X < e.pos.X+1 &&
							n.pos.X+n.size.X > e.pos.X &&
							n.pos.Y < e.pos.Y+1 &&
							n.pos.Y+n.size.Y > e.pos.Y {
							if n.passable {
								nodeIndexStart = nI
							}
						} else if n.pos.X < player.pos.X+1 &&
							n.pos.X+n.size.X > player.pos.X &&
							n.pos.Y < player.pos.Y+1 &&
							n.pos.Y+n.size.Y > player.pos.Y {
							if n.passable {
								nodeIndexEnd = nI
							}
						}
					}
					e.currentPath = astar(nodeIndexStart, nodeIndexEnd, e.nodes, e.size)
				}
			}
		}
		if e.eye.state != 0 { // Open eye
			if e.eye.animation.frameNumber >= e.eye.animation.frameNumberMax-1 {
				e.eye.animation = e.eye.animations.lookingAnimation
				e.canAttack = true
			}
		}

		e.animation.frameSpeedMax = e.moveAnimationSpeed

		if len(e.currentPath) > 0 {
			if e.center.X < e.currentPath[0].pos.X+e.currentPath[0].size.X &&
				e.center.X+1 > e.currentPath[0].pos.X &&
				e.center.Y < e.currentPath[0].pos.Y+e.currentPath[0].size.Y &&
				e.center.Y+1 > e.currentPath[0].pos.Y {
				e.currentPath = append(e.currentPath[:0], e.currentPath[1:]...)
			} else {
				e.targetPosition = pixel.V(e.currentPath[0].pos.X+e.currentPath[0].size.X/2.0, e.currentPath[0].pos.Y+e.currentPath[0].size.Y/2.0)
			}
		}
		if !e.attacking {
			if e.targetPosition.X-(e.moveSpeed*dt) > e.center.X {
				e.moveVector.X = 1
				if e.currentAnimation != 2 {
					e.animation = e.animations.rightAnimation
					e.currentAnimation = 2
				}
			} else if e.targetPosition.X+(e.moveSpeed*dt) < e.center.X {
				e.moveVector.X = -1
				if e.currentAnimation != 1 {
					e.animation = e.animations.leftAnimation
					e.currentAnimation = 1
				}
			}
			if e.targetPosition.Y > e.center.Y {
				e.moveVector.Y = 1
			} else if e.targetPosition.Y < e.center.Y {
				e.moveVector.Y = -1
			}
			e.noSoundTimer -= 1 * dt
		}
	}
	e.pos = pixel.V(e.pos.X+(e.moveSpeed*dt)*e.moveVector.X, e.pos.Y+(e.moveSpeed*dt)*e.moveVector.Y)
	e.center = pixel.V(e.pos.X+(e.size.X/2), e.pos.Y+(e.size.Y/2))
	e.eye.center = pixel.V(e.center.X, e.center.Y+(e.size.Y/2)+e.eye.size.Y)

	if e.canAttack {
		e.attackHandler(p, dt)
	}
	if e.attackCooldown > 0. {
		e.attackCooldown -= 1 * dt
	}
}

func (e *Enemy) attackHandler(p *Player, dt float64) {
	if circlularCollisionCheck(e.attackCheckRadius, p.radius, calculateDistance(p.center, e.center)) {
		if e.attackCooldown <= 0. && !e.attacking {
			e.attacking = true
			e.attackCooldown = e.attackCooldownMax
		} else {
			fmt.Println(e.attackCooldown)
			/**
			 * TODO:
			 *
			 * Add animation
			 * Set attacking to false after animation finished
			 * Add shockwave when the enemy hits ground
			 * Response from player when attacked (maybe work on death)
			 * Maybe add a outline of where the enemy can attack?
			 **/
		}
	}
}

/*
	canAttack          bool
	attackCooldown     float64
	attackCooldownMax  float64
	attackCheckRadius  float64
*/
