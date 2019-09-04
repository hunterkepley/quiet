package main

import (
	"fmt"
	"image/color"
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var (
	// Basic variables
	frames    = 0 // Fps
	second    = time.Tick(time.Second)
	gameState = 1 // 0 is in a game, 1 is in the menu. Keeps track of rendering and updating.
	dt        float64

	imageScale       = 2.
	winWidth         = 1024.
	winHeight        = 768.
	currentWinWidth  = winWidth
	currentWinHeight = winHeight

	player Player

	viewMatrix pixel.Matrix

	currentLevel Level

	currentShader string

	clearColor color.Color

	// Draws bounding boxes of the rain deadzones for debugging
	drawRainDeadzones = false
)

const ()

func run() {
	cfg := pixelgl.WindowConfig{ // Defines window struct
		Title:     "QUIET",
		Bounds:    pixel.R(0, 0, winWidth, winHeight),
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg) // Creates window
	if err != nil {                    // Deals with error
		panic(err)
	}

	viewCanvas := pixelgl.NewCanvas(pixel.R(win.Bounds().Min.X, win.Bounds().Min.Y, win.Bounds().W(), win.Bounds().H()))

	loadResources()

	// Set up the matrices for the view of the world
	letterBox(win)

	player = createPlayer(pixel.V(200, 200), 0, playerSpritesheets.playerIdleRightSheet.sheet, true, imageScale)

	// Set up all levels
	loadLevels()

	viewCanvas.SetFragmentShader(regularShader)

	clearColor = color.RGBA{0x0a, 0x0a, 0x0a, 0x0a}

	last := time.Now()  // For fps decoupled updates
	for !win.Closed() { // Game loop
		if currentWinHeight != win.Bounds().H() || currentWinWidth != win.Bounds().W() {
			// Resize event
			currentWinWidth = win.Bounds().W()
			currentWinHeight = win.Bounds().H()
			letterBox(win)
		}
		imd := imdraw.New(nil)
		dt = time.Since(last).Seconds() // For fps decoupled updates.
		if dt > 0.25 {
			dt = 0.
		}
		last = time.Now() // ^
		win.Clear(colornames.Black)
		viewCanvas.Clear(clearColor)
		imd.Clear()

		// TODO: temporary
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Println(win.MousePosition())
		}

		switch gameState {
		case 0: // In game, will probably change... Not sure
			updateGame(win, viewCanvas, dt)
			renderGame(win, viewCanvas, imd, dt)
			clearColor = color.RGBA{0x0a, 0x0a, 0x0a, 0x0a}
		case 1: // In menu [?Likely to be separate menus?]
			updateMenu(win, viewCanvas, dt)
			renderMenu(win, viewCanvas)
			clearColor = color.Black
		}

		viewCanvas.Draw(win, viewMatrix)
		imd.Draw(win)
		win.Update()

		frames++ // FPS is dealt here
		select { // Waits for the block to finish
		case <-second: // A second has passed
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames)) // Appends fps to title
			frames = 0                                                   // Reset it my dude
		default:
		} // FPS is done
	}
}

func loadResources() {
	//Load the player sprite sheets for the game
	loadPlayerSpritesheets()
	//Load the player spritebatches for the game
	loadPlayerBatches()
	//Load the sound wave spritebatches for the game
	loadSoundWaveBatches()
	//Load the object spritesheets for the game
	loadObjectSpritesheets()
	//Load the object spritebatches for the game
	loadObjectBatches()
	//Load images for game that aren't spritesheets
	loadObjectImages()
	//Load images for the visible sound bubbles/other sound images
	loadSoundImages()
	//Load spritesheets for sound waves
	loadSoundWaveSpritesheets()
	//Load images for the main menu
	loadMenuImages()
	//Load music for the game
	loadMusic()
	//Load sound effects for the game
	loadAudio() //ENABLE THIS WHEN READY TO TEST
	//Load enemy images for the game
	loadEnemyImages()
	//Load enemy spritesheets for the game
	loadEnemySpriteSheets()
	//Create main menu
	mainMenu = createMainMenu()
}

func main() {
	pixelgl.Run(run)
}
