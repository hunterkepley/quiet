package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// Enemy ... All basic enemies in the game
type Enemy struct {
	pos            pixel.Vec
	center         pixel.Vec
	size           pixel.Vec
	pic            pixel.Picture
	sprite         *pixel.Sprite
	sizeDiminisher float64
	moveSpeed      float64
	moveVector     pixel.Vec // 1, 1 for moving top right, 0, 1 for moving up, etc.
}

func createEnemy(pos pixel.Vec, pic pixel.Picture, sizeDiminisher float64, moveSpeed float64) Enemy {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return Enemy{
		pos,
		pixel.ZV,
		size,
		pic,
		sprite,
		sizeDiminisher,
		moveSpeed,
		pixel.ZV,
	}
}

func (e Enemy) render(viewCanvas *pixelgl.Canvas, imd *imdraw.IMDraw) {
	mat := pixel.IM.
		Moved(e.center).
		Scaled(e.center, imageScale)
	e.sprite.Draw(viewCanvas, mat)
}

func (e *Enemy) update(dt float64, p Player) {
	e.moveVector = pixel.V(0, 0)
	if p.center.X > e.center.X {
		e.moveVector.X = 1
	} else if p.center.X < e.center.X {
		e.moveVector.X = -1
	}
	if p.center.Y > e.center.Y {
		e.moveVector.Y = 1
	} else if p.center.Y < e.center.Y {
		e.moveVector.Y = -1
	}
	e.pos = pixel.V(e.pos.X+(e.moveSpeed*dt)*e.moveVector.X, e.pos.Y+(e.moveSpeed*dt)*e.moveVector.Y)
	e.center = pixel.V(e.pos.X+(e.size.X/2), e.pos.Y+(e.size.Y/2))
}
