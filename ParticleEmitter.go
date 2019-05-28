package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ParticleEmitter struct {
	pos           pixel.Vec
	particles     []Particle
	sheet         Spritesheet
	spawnTimer    float64 // Time between particles
	spawnTimerMax float64
	endTimer      float64 // Time before they get deleted
	reverseX      bool    // Reverse halfway through
	reverseY      bool
}

type Particle struct {
	pos         pixel.Vec
	pic         pixel.Picture
	sprite      pixel.Sprite
	rotation    float64
	change      pixel.Vec
	endTimer    float64 // Time before they get deleted
	endTimerMax float64
	offset      pixel.Vec // Offset for movement
	reverseX    bool
	reverseY    bool
}

func createParticleEmitter(sheet Spritesheet, pos pixel.Vec, spawnTimer float64, endTimer float64, reverseX bool, reverseY bool) ParticleEmitter {
	return ParticleEmitter{pos, []Particle{}, sheet, spawnTimer, spawnTimer, endTimer, reverseX, reverseY}
}

func (p *ParticleEmitter) update(dt float64, newPos pixel.Vec, change pixel.Vec, offset pixel.Vec, canUpdate bool) {
	p.pos = newPos
	for i := 0; i < len(p.particles); i++ {
		p.particles[i].pos.X += p.particles[i].change.X*dt + p.particles[i].offset.X*dt
		p.particles[i].pos.Y += p.particles[i].change.Y*dt + p.particles[i].offset.Y*dt
	}
	if canUpdate {
		if p.spawnTimer < 0 {
			randomChoice := rand.Intn(p.sheet.numberOfFrames)
			sprite := pixel.NewSprite(p.sheet.sheet, p.sheet.frames[randomChoice])
			offset = pixel.V((float64(rand.Intn(int(offset.X)*2)) - offset.X), (float64(rand.Intn(int(offset.Y)*2)) - offset.Y))
			p.particles = append(p.particles, Particle{
				p.pos,
				p.sheet.sheet,
				*sprite,
				float64(rand.Intn(360)),
				change,
				p.endTimer,
				p.endTimer,
				offset,
				p.reverseX,
				p.reverseY,
			})
			p.spawnTimer = p.spawnTimerMax
		} else {
			p.spawnTimer -= 1 * dt
		}
	}
	if len(p.particles) > 1 {
		for i := 0; i < len(p.particles); i++ {
			if p.particles[i].endTimer < 0 {
				p.particles = p.particles[1:]
			} else {
				if p.reverseX || p.reverseY {
					timeLeftRatio := p.particles[i].endTimerMax / p.particles[i].endTimer
					if (p.particles[i].reverseX || p.particles[i].reverseY) && timeLeftRatio > 2. {
						if p.reverseX {
							p.particles[i].change.X *= -1
							p.particles[i].reverseX = false
						} else {
							p.particles[i].change.Y *= -1
							p.particles[i].reverseY = false
						}
					}
				}
				p.particles[i].endTimer -= 1 * dt
			}
		}
	} else {
		if !canUpdate {
			p.particles = []Particle{}
		}
	}
}

func (p *ParticleEmitter) render(win *pixelgl.Window) {
	for i := 0; i < len(p.particles); i++ {
		mat := pixel.IM.
			Rotated(pixel.ZV, p.particles[i].rotation).
			Moved(p.particles[i].pos)
		//p.particles[i].sprite.Draw(batches.fireParticlesBatch, pixel.IM.Rotated(pixel.ZV, p.particles[i].rotation).Moved(p.particles[i].pos))
		p.particles[i].sprite.Draw(win, mat)
	}
}
