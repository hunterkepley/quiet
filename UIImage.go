package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// UIImage ... Images to be displayed on UI (menu images mainly), basically any non-interactable/non-updating UI
type UIImage struct {
	pos    pixel.Vec
	center pixel.Vec
	size   pixel.Vec
	pic    pixel.Picture
	sprite *pixel.Sprite
}

func createUIImage(pos pixel.Vec, pic pixel.Picture) UIImage {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	center := pixel.V(pos.X+(size.X/2), pos.Y+(size.Y/2))
	return UIImage{
		pos,
		center,
		size,
		pic,
		sprite,
	}
}

func (i *UIImage) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(i.center).
		Scaled(i.center, imageScale)
	i.sprite.Draw(viewCanvas, mat)
}

func (i *UIImage) update() {
	// Empty for now I guess
}
