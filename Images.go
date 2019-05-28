package main

import (
	"github.com/faiface/pixel"
)

var (
	images Images
)

/*Images ... All the non-spritesheet images in the game*/
type Images struct {
	box1 pixel.Picture
}

func loadImages() {
	images = Images{
		loadPicture("Art/Objects/Scenery/box1.png"),
	}
}
