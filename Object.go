package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Object struct {
	pos             pixel.Vec
	center          pixel.Vec
	size            pixel.Vec
	pic             pixel.Picture
	sprite          pixel.Sprite
	radius          float64
	sizeDiminisher  float64
	inFrontOfPlayer bool // If the object is rendered in front of the player or not
}

var (
	backgroundObjects []Object
	foregroundObjects []Object
)

func createObject(pos pixel.Vec, pic pixel.Picture, sizeDiminisher float64) Object {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return Object{
		pos,
		pixel.ZV,
		pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y),
		pic,
		*sprite,
		pic.Bounds().Size().Y / 2,
		sizeDiminisher,
		true,
	}
}

func (o *Object) update(p *Player) {
	o.center = pixel.V(o.pos.X+(o.size.X/2), o.pos.Y+(o.size.Y/2))
	o.playerCollision(p)
}

func (o Object) render(win *pixelgl.Window, viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(o.center).
		Scaled(o.center, imageScale)
	o.sprite.Draw(viewCanvas, mat)
}

func (o *Object) playerCollision(p *Player) {
	if p.pos.Y > o.pos.Y+(o.size.Y*imageScale)/o.sizeDiminisher {
		o.inFrontOfPlayer = true
	} else {
		o.inFrontOfPlayer = false
	}
	if p.pos.X < o.pos.X+(o.size.X*imageScale) &&
		p.pos.X+(p.size.X*imageScale) > o.pos.X &&
		p.pos.Y < o.pos.Y+(o.size.Y*imageScale)/o.sizeDiminisher &&
		p.pos.Y+(p.size.Y*imageScale)/p.footSizeDiminisher > o.pos.Y {
		fmt.Println("Collided")
	}
}
