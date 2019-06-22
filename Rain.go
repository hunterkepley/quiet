package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	rain     []Rain
	splashes []Splash
)

// Rain ... It's rain
type Rain struct {
	pos        pixel.Vec
	center     pixel.Vec
	size       pixel.Vec
	pic        pixel.Picture
	sprite     pixel.Sprite
	rainSpeedY float64
	rainSpeedX float64
	endHeight  float64
}

// Splash ... It's rain splash
type Splash struct {
	pos       pixel.Vec
	center    pixel.Vec
	size      pixel.Vec
	pic       pixel.Picture
	sprite    pixel.Sprite
	animation Animation
}

func createRain(pos pixel.Vec) Rain {
	rainChoice := rand.Intn(objectSpritesheets.rainSheet.numberOfFrames)
	pic := objectSpritesheets.rainSheet.sheet
	sprite := pixel.NewSprite(pic, objectSpritesheets.rainSheet.frames[rainChoice])
	rainSpeedY := float64(rand.Intn(1000) + 600)
	rainSpeedX := float64(rand.Intn(200) + 10)
	endHeight := float64(rand.Intn(int(winHeight)))
	size := pixel.V(pic.Bounds().Size().X/float64(len(objectSpritesheets.rainSheet.frames)), pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return Rain{
		pos,
		pixel.ZV,
		size,
		pic,
		*sprite,
		rainSpeedY,
		rainSpeedX,
		endHeight,
	}
}

func createSplash(pos pixel.Vec) Splash {
	pic := objectSpritesheets.rainSplashSheet.sheet
	sprite := pixel.NewSprite(pic, objectSpritesheets.rainSplashSheet.frames[0])
	size := pixel.V(pic.Bounds().Size().X/float64(len(objectSpritesheets.rainSplashSheet.frames)), pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	center := pixel.V(pos.X+(size.X/2), pos.Y+(size.Y/2))
	rainSplashSpeed := 0.04

	return Splash{
		pos,
		center,
		size,
		pic,
		*sprite,
		createAnimation(objectSpritesheets.rainSplashSheet, rainSplashSpeed),
	}
}

func (r *Rain) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(r.pos).
		Scaled(r.center, imageScale)
	r.sprite.Draw(objectBatches.rainBatch, mat)
	objectBatches.rainBatch.Draw(viewCanvas)
}

func (r *Rain) update(dt float64) {
	r.pos.Y -= r.rainSpeedY * dt
	r.pos.X += r.rainSpeedX * dt
	r.center = pixel.V(r.pos.X+(r.size.X/2), r.pos.Y+(r.size.Y/2))
}

func updateRain(rainDeadZones []pixel.Rect, player Player, dt float64) {
	for i := 0; i < len(rain); i++ {
		rain[i].update(dt)
		if rain[i].pos.Y < rain[i].endHeight {
			canSplash := true
			for j := 0; j < len(rainDeadZones); j++ {
				if rain[i].pos.X > rainDeadZones[j].Min.X &&
					rain[i].pos.X < rainDeadZones[j].Max.X &&
					rain[i].pos.Y > rainDeadZones[j].Min.Y &&
					rain[i].pos.Y < rainDeadZones[j].Max.Y {
					// This checks if a splash is possible [if the rain ends on the side of an object, etc]
					canSplash = false
				}
			}
			if rain[i].pos.X > player.pos.X &&
				rain[i].pos.X < player.pos.X+player.size.X &&
				rain[i].pos.Y > player.pos.Y &&
				rain[i].pos.Y < player.pos.Y+player.size.Y {
				// Check if a rain drop ends on the player [same reason as dead zones]
				canSplash = false
			}
			if canSplash {
				splashes = append(splashes, createSplash(rain[i].pos))
			}
			rain = append(rain[:i], rain[i+1:]...)
		}
	}
	for i := 0; i < len(splashes); i++ {
		splashes[i].update(dt)
		if splashes[i].animation.frameNumber == splashes[i].animation.frameNumberMax-1 {
			splashes = append(splashes[:i], splashes[i+1:]...)
		}
	}
}

func renderRain(viewCanvas *pixelgl.Canvas) {
	objectBatches.rainBatch.Clear()
	for i := 0; i < len(rain); i++ {
		rain[i].render(viewCanvas)
	}
}

func renderSplashes(viewCanvas *pixelgl.Canvas) {
	objectBatches.rainSplashBatch.Clear()
	for i := 0; i < len(splashes); i++ {
		splashes[i].render(viewCanvas)
	}
}

func (s *Splash) update(dt float64) {
	s.sprite = s.animation.animate(dt)
}

func (s *Splash) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(s.pos).
		Scaled(s.center, imageScale)
	s.sprite.Draw(objectBatches.rainSplashBatch, mat)
	objectBatches.rainSplashBatch.Draw(viewCanvas)
}
