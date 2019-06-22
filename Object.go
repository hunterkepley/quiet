package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Object ... Objects that the player can collide with
type Object struct {
	pos              pixel.Vec
	center           pixel.Vec
	size             pixel.Vec
	pic              pixel.Picture
	sprite           pixel.Sprite
	radius           float64
	sizeDiminisher   float64
	inFrontOfPlayer  bool // If the object is rendered in front of the player or not
	backgroundObject bool // true if in background
	foregroundObject bool // true if in foreground
	playerCollidable bool // true if collides with player
}

var (
	backgroundObjects []Object
	foregroundObjects []Object
)

func createObject(pos pixel.Vec, pic pixel.Picture, sizeDiminisher float64, backgroundObject bool, foregroundObject bool, playerCollidable bool) Object {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	inFrontOfPlayer := true
	if backgroundObject {
		inFrontOfPlayer = false
	}
	return Object{
		pos,
		pixel.ZV,
		size,
		pic,
		*sprite,
		pic.Bounds().Size().Y / 2,
		sizeDiminisher,
		inFrontOfPlayer,
		backgroundObject,
		foregroundObject,
		playerCollidable,
	}
}

func (o *Object) update(p *Player) {
	o.center = pixel.V(o.pos.X+(o.size.X/2), o.pos.Y+(o.size.Y/2))
	if !o.playerCollidable {
		o.playerCollision(p)
	}
}

func (o Object) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(o.center).
		Scaled(o.center, imageScale)
	o.sprite.Draw(viewCanvas, mat)
}

func (o *Object) playerCollision(p *Player) {
	if p.pos.Y > o.pos.Y+(o.size.Y*imageScale)/(o.sizeDiminisher+1.) {
		o.inFrontOfPlayer = true
	} else {
		o.inFrontOfPlayer = false
	}
	if p.pos.X < o.pos.X+o.size.X &&
		p.pos.X+p.size.X > o.pos.X &&
		p.pos.Y < o.pos.Y+o.size.Y/o.sizeDiminisher &&
		p.pos.Y+p.size.Y/p.footSizeDiminisher > o.pos.Y {
		fmt.Println("Collided")
	}
}
