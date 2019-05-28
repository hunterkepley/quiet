package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var (
	// Basic variables
	frames    = 0 // Fps
	second    = time.Tick(time.Second)
	gameState = 0 // 0 is in a game, 1 is in the menu. Keeps track of rendering and updating.
	dt        float64

	camZoom = 2.

	player Player

	testBox Object
)

const (
	winWidth  = 1024
	winHeight = 768
)

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

	//Load the sprite sheets for the game
	loadSpritesheets()
	//Load batches for the game
	loadBatches()
	//load images for game that aren't spritesheets
	loadImages()

	testBox = createObject(pixel.V(100., 100.), images.box1)

	player = createPlayer(pixel.V(0, 0), 0, spritesheets.playerIdleDownSheet.sheet, true)

	last := time.Now()  // For fps decoupled updates
	for !win.Closed() { // Game loop
		imd := imdraw.New(nil)
		dt = time.Since(last).Seconds() // For fps decoupled updates.
		if dt > 0.25 {
			dt = 0.
		}
		last = time.Now() // ^
		win.Clear(color.RGBA{0x0a, 0x0a, 0x0a, 0x0a})
		imd.Clear()
		switch gameState {
		case 0: // In game, will probably change... Not sure
			updateGame(win, dt)
			renderGame(win, imd)
		case 1: // In menu [?Likely to be separate menus?]
			updateMenu(dt)
			renderMenu(win)
		}

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
func main() {
	pixelgl.Run(run)

}
