package main

import (
	"github.com/faiface/pixel"
)

var (
	objectImages ObjectImages
)

//ObjectImages ... All the non-spritesheet images in the game
type ObjectImages struct {
	box1         pixel.Picture
	gasBody      pixel.Picture
	gasRoof      pixel.Picture
	gasLeftPole  pixel.Picture
	gasRightPole pixel.Picture
	gasStreet    pixel.Picture
	gasFence     pixel.Picture
}

func loadObjectImages() {
	objectImages = ObjectImages{
		loadPicture("./Art/Objects/Scenery/box1.png"),
		loadPicture("./Art/Objects/Buildings/gas_body.png"),
		loadPicture("./Art/Objects/Buildings/gas_roof.png"),
		loadPicture("./Art/Objects/Buildings/gas_left_pole.png"),
		loadPicture("./Art/Objects/Buildings/gas_right_pole.png"),
		loadPicture("./Art/Objects/Backgrounds/street1.png"),
		loadPicture("./Art/Objects/Buildings/gas_fence.png"),
	}
}
