package main

import (
	"github.com/faiface/pixel"
)

var (
	spritesheets Spritesheets
)

/*Spritesheet ... Holds a picture of a spritesheet and the frames of each single picture*/
type Spritesheet struct {
	sheet          pixel.Picture
	frames         []pixel.Rect
	numberOfFrames int
}

/*Spritesheets ... All the spritesheets in the game*/
type Spritesheets struct {
	playerIdleDownSheet Spritesheet
}

func loadSpritesheets() {
	// Player spritesheet
	playerSheet := loadPicture("./Art/Player/idle_down.png")
	spritesheets = Spritesheets{
		createSpriteSheet(playerSheet, 4),
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
