package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	backgroundObjects []Object
	foregroundObjects []Object
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
	inFrontOfPlayer  bool    // If the object is rendered in front of the player or not
	backgroundObject bool    // true if in background
	foregroundObject bool    // true if in foreground
	playerCollidable bool    // true if collides with player
	dBDiminisher     float64 // Amount of dB it takes to go through the object.

	// Collision rects
	top             pixel.Rect
	left            pixel.Rect
	right           pixel.Rect
	bottom          pixel.Rect
	collisionOffset float64
	hitboxes        bool
}


func createObject(pos pixel.Vec, pic pixel.Picture, sizeDiminisher float64, backgroundObject bool, foregroundObject bool, playerCollidable bool, dBDiminisher float64) Object {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	inFrontOfPlayer := true

	collisionOffset := 2.
	top := pixel.R(0, 0, 0, 0)
	left := pixel.R(0, 0, 0, 0)
	right := pixel.R(0, 0, 0, 0)
	bottom := pixel.R(0, 0, 0, 0)
	hitboxes := false

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
		dBDiminisher,
		top,
		left,
		right,
		bottom,
		collisionOffset,
		hitboxes,
	}
}

func (o *Object) update(p *Player) {
	o.center = pixel.V(o.pos.X+(o.size.X/2), o.pos.Y+(o.size.Y/2))
	if o.playerCollidable {
		o.playerCollision(p)
		o.top = pixel.R(o.pos.X+o.collisionOffset, o.pos.Y+(o.size.Y/o.sizeDiminisher)-o.collisionOffset, o.pos.X+o.size.X-o.collisionOffset, o.pos.Y+(o.size.Y/o.sizeDiminisher))
		o.left = pixel.R(o.pos.X, o.pos.Y+o.collisionOffset, o.pos.X+o.collisionOffset, o.pos.Y+(o.size.Y/o.sizeDiminisher)-o.collisionOffset)
		o.right = pixel.R(o.pos.X+o.size.X-o.collisionOffset, o.pos.Y+o.collisionOffset, o.pos.X+o.size.X, o.pos.Y+(o.size.Y/o.sizeDiminisher)-o.collisionOffset)
		o.bottom = pixel.R(o.pos.X+o.collisionOffset, o.pos.Y, o.pos.X+o.size.X-o.collisionOffset, o.pos.Y+o.collisionOffset)
	}
}

func (o Object) render(viewCanvas *pixelgl.Canvas, imd *imdraw.IMDraw, p Player) {
	mat := pixel.IM.
		Moved(o.center).
		Scaled(o.center, imageScale)
	o.sprite.Draw(viewCanvas, mat)
	if o.hitboxes {
		o.renderHitboxes(imd, player)
	}
}

func (o *Object) playerCollision(p *Player) {
	if p.pos.Y > o.pos.Y+o.size.Y/(o.sizeDiminisher+1.) {
		o.inFrontOfPlayer = true
	} else {
		o.inFrontOfPlayer = false
	}
	if p.pos.X < o.pos.X+o.size.X &&
		p.pos.X+p.size.X > o.pos.X &&
		p.pos.Y < o.pos.Y+o.size.Y/o.sizeDiminisher &&
		p.pos.Y+p.size.Y/p.footSizeDiminisher > o.pos.Y {
		if collisionCheck(p.footHitBox, o.top) {
			p.pos.Y = o.top.Max.Y
		} else if collisionCheck(p.footHitBox, o.bottom) {
			p.pos.Y = o.bottom.Min.Y - p.size.Y/p.footSizeDiminisher
		}
		if collisionCheck(p.footHitBox, o.right) {
			p.pos.X = o.right.Max.X
		} else if collisionCheck(p.footHitBox, o.left) {
			p.pos.X = o.left.Min.X - p.size.X
		}
	}
}

func (o *Object) renderHitboxes(imd *imdraw.IMDraw, p Player) {
	imd.Color = colornames.Cyan
	width := 1.
	imd.Push(o.top.Min, o.top.Max)
	imd.Rectangle(width)
	imd.Push(o.left.Min, o.left.Max)
	imd.Rectangle(width)
	imd.Push(o.right.Min, o.right.Max)
	imd.Rectangle(width)
	imd.Push(o.bottom.Min, o.bottom.Max)
	imd.Rectangle(width)
	imd.Push(p.footHitBox.Min, p.footHitBox.Max)
	imd.Rectangle(width)
}
