package main

import (
	"github.com/faiface/pixel"
)

var (
	playerSpritesheets PlayerSpritesheets
)

/*Spritesheet ... Holds a picture of a spritesheet and the frames of each single picture*/
type Spritesheet struct {
	sheet          pixel.Picture
	frames         []pixel.Rect
	numberOfFrames int
}

/*Spritesheets ... All the spritesheets in the game*/
type PlayerSpritesheets struct {
	playerIdleRightSheet Spritesheet
	playerIdleUpSheet    Spritesheet
}

func loadPlayerSpritesheets() {
	// Player spritesheet
	playerIdleRightSheet := loadPicture("./Art/Player/idle_right.png")
	playerIdleUpSheet := loadPicture("./Art/Player/idle_up.png")
	playerSpritesheets = PlayerSpritesheets{
		createSpriteSheet(playerIdleRightSheet, 4),
		createSpriteSheet(playerIdleUpSheet, 4),
	}
}

func createSpriteSheet(sheet pixel.Picture, frames float64) Spritesheet {
	w := float64(int(sheet.Bounds().W() / frames))
	h := sheet.Bounds().H()
	var sheetFrames []pixel.Rect
	for x := sheet.Bounds().Min.X; x < sheet.Bounds().Max.X; x += w {
		for y := sheet.Bounds().Min.Y; y < sheet.Bounds().Max.Y; y += h {
			sheetFrames = append(sheetFrames, pixel.R(x, y, x+w, y+h))
		}
	}
	numberOfFrames := frames
	return Spritesheet{sheet, sheetFrames, int(numberOfFrames)}
}
