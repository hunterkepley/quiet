package main

import (
	"github.com/faiface/pixel"
)

var (
	playerSpritesheets PlayerSpritesheets
	objectSpritesheets ObjectSpritesheets
	enemySpriteSheets  EnemySpriteSheets
)

/*Spritesheet ... Holds a picture of a spritesheet and the frames of each single picture*/
type Spritesheet struct {
	sheet          pixel.Picture
	frames         []pixel.Rect
	numberOfFrames int
}

/*PlayerSpritesheets ... All the player spritesheets in the game*/
type PlayerSpritesheets struct {
	playerIdleRightSheet Spritesheet
	playerIdleUpSheet    Spritesheet
	playerIdleDownSheet  Spritesheet
	playerIdleLeftSheet  Spritesheet
	// Sound wave stuff
	soundWaveBTrailSheet Spritesheet
}

//LarvaSpriteSheets ... All the larva spritesheets in the game
type LarvaSpriteSheets struct {
	leftSpriteSheet             Spritesheet
	rightSpriteSheet            Spritesheet
	attackRaiseSpriteSheetLeft  Spritesheet
	attackRaiseSpriteSheetRight Spritesheet
}

//EnemySpriteSheets ... All the enemy spritesheets in the game
type EnemySpriteSheets struct {
	larvaSpriteSheets LarvaSpriteSheets
	eyeLookingSheet   Spritesheet
	eyeOpeningSheet   Spritesheet
	eyeClosingSheet   Spritesheet
}

/*ObjectSpritesheets ... All the object spritesheets in the game*/
type ObjectSpritesheets struct {
	rainSheet       Spritesheet
	rainSplashSheet Spritesheet
	trashCanSheet   Spritesheet
}

func loadPlayerSpritesheets() {
	// Player spritesheets
	playerIdleRightSheet := loadPicture("./Resources/Art/Player/idle_right.png")
	playerIdleUpSheet := loadPicture("./Resources/Art/Player/idle_up.png")
	playerIdleDownSheet := loadPicture("./Resources/Art/Player/idle_down.png")
	playerIdleLeftSheet := loadPicture("./Resources/Art/Player/idle_left.png")

	// Sound wave trail sheets
	soundWaveBTrailSheet := loadPicture("./Resources/Art/Player/sound_wave_b_trail_sheet.png")
	playerSpritesheets = PlayerSpritesheets{
		createSpriteSheet(playerIdleRightSheet, 4),
		createSpriteSheet(playerIdleUpSheet, 4),
		createSpriteSheet(playerIdleDownSheet, 4),
		createSpriteSheet(playerIdleLeftSheet, 4),
		// Sound wave trail sheets
		createSpriteSheet(soundWaveBTrailSheet, 4),
	}
}

func loadEnemySpriteSheets() {
	// Enemy spritesheets
	larvaLeftSheet := loadPicture("./Resources/Art/Enemies/Larva/left_spritesheet.png")
	larvaRightSheet := loadPicture("./Resources/Art/Enemies/Larva/right_spritesheet.png")
	larvaAttackRaiseSheetLeft := loadPicture("./Resources/Art/Enemies/Larva/attack_raise_spritesheet_left.png")
	larvaAttackRaiseSheetRight := loadPicture("./Resources/Art/Enemies/Larva/attack_raise_spritesheet_right.png")
	eyeLookingSheet := loadPicture("./Resources/Art/Enemies/eye_looking.png")
	eyeOpeningSheet := loadPicture("./Resources/Art/Enemies/eye_opening.png")
	eyeClosingSheet := loadPicture("./Resources/Art/Enemies/eye_closing.png")
	enemySpriteSheets = EnemySpriteSheets{
		LarvaSpriteSheets{
			createSpriteSheet(larvaLeftSheet, 4),
			createSpriteSheet(larvaRightSheet, 4),
			createSpriteSheet(larvaAttackRaiseSheetLeft, 7),
			createSpriteSheet(larvaAttackRaiseSheetRight, 4),
		},
		createSpriteSheet(eyeLookingSheet, 12),
		createSpriteSheet(eyeOpeningSheet, 4),
		createSpriteSheet(eyeClosingSheet, 4),
	}
}

func loadObjectSpritesheets() {
	// Object spritesheets
	rainSheet := loadPicture("./Resources/Art/Weather/rain.png")
	rainSplashSheet := loadPicture("./Resources/Art/Weather/rain_splash.png")
	trashCanSheet := loadPicture("./Resources/Art/Objects/Scenery/l1/trash_can_sheet.png")
	objectSpritesheets = ObjectSpritesheets{
		createSpriteSheet(rainSheet, 5),
		createSpriteSheet(rainSplashSheet, 6),
		createSpriteSheet(trashCanSheet, 4),
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
	return Spritesheet{sheet, sheetFrames, int(frames)}
}
