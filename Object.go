package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Object struct {
	pos    pixel.Vec
	center pixel.Vec
	size   pixel.Vec
	pic    pixel.Picture
	sprite pixel.Sprite
	radius float64
}

func createObject(pos pixel.Vec, pic pixel.Picture) Object {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return Object{
		pos,
		pixel.ZV,
		pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y),
		pic,
		*sprite,
		pic.Bounds().Size().Y / 2,
	}
}

func (o *Object) update() {
	o.center = pixel.V(o.pos.X+(o.size.X/2), o.pos.Y+(o.size.Y/2))
}

func (o Object) render(win *pixelgl.Window, imd *imdraw.IMDraw) {
	mat := pixel.IM.
		Moved(o.center)
	o.sprite.Draw(win, mat)
}
