package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	rain []Rain
)

// Rain ... It's rain
type Rain struct {
	pos       pixel.Vec
	center    pixel.Vec
	size      pixel.Vec
	pic       pixel.Picture
	sprite    pixel.Sprite
	rainSpeed float64
	endHeight float64
}

func createRain(pos pixel.Vec) Rain {
	rainChoice := rand.Intn(objectSpritesheets.rainSheet.numberOfFrames)
	pic := objectSpritesheets.rainSheet.sheet
	sprite := pixel.NewSprite(pic, objectSpritesheets.rainSheet.frames[rainChoice])
	rainSpeed := float64(rand.Intn(600) + 580)
	endHeight := float64(rand.Intn(int(winHeight) / 2))
	size := pixel.V(pic.Bounds().Size().X/float64(len(playerSpritesheets.playerIdleRightSheet.frames)), pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return Rain{
		pos,
		pixel.ZV,
		size,
		pic,
		*sprite,
		rainSpeed,
		endHeight,
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
	r.pos.Y -= r.rainSpeed * dt
	r.center = pixel.V(r.pos.X+(r.size.X/2), r.pos.Y+(r.size.Y/2))
}

func updateRain(dt float64) {
	for i := 0; i < len(rain); i++ {
		rain[i].update(dt)
		if rain[i].pos.Y < rain[i].endHeight {
			rain = append(rain[:i], rain[i+1:]...)
		}
	}
}

func renderRain(viewCanvas *pixelgl.Canvas) {
	objectBatches.rainBatch.Clear()
	for i := 0; i < len(rain); i++ {
		rain[i].render(viewCanvas)
	}
}
