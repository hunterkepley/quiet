package main

import (
	"github.com/faiface/pixel"
)

var (
	objectImages ObjectImages
	menuImages   MenuImages
	soundImages  SoundImages
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
	gasLight     pixel.Picture
}

//SoundImages ... All the visible sound images
type SoundImages struct {
	playerSoundBubble pixel.Picture
}

//MenuImages ... All the menu images that aren't animated
type MenuImages struct {
	title pixel.Picture
}

func loadObjectImages() {
	objectImages = ObjectImages{
		loadPicture("./Resources/Art/Objects/Scenery/box1.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_body.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_roof.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_left_pole.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_right_pole.png"),
		loadPicture("./Resources/Art/Objects/Backgrounds/l1/street1.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_fence.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_light.png"),
	}
}

func loadSoundImages() {
	soundImages = SoundImages{
		loadPicture("./Resources/Art/Player/sound_bubble.png"),
	}
}

func loadMenuImages() {
	menuImages = MenuImages{
		loadPicture("./Resources/Art/UI/MM/title.png"),
	}
}
