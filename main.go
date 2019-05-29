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
	gameState = 0 // 0 is in a game, 1 is in the menu. Keeps track of rendering and updating.
	dt        float64

	camZoom   = 2.
	winWidth  = 1024.
	winHeight = 768.

	player Player

	testBox Object

	viewMatrix pixel.Matrix
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
	viewCanvas := pixelgl.NewCanvas(pixel.R(win.Bounds().Min.X, win.Bounds().Min.Y, win.Bounds().W()/camZoom, win.Bounds().H()/camZoom))

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
		if winHeight != win.Bounds().H() || winWidth != win.Bounds().W() {
			// Resize event
			letterBox(win, player)
		}
		imd := imdraw.New(nil)
		dt = time.Since(last).Seconds() // For fps decoupled updates.
		if dt > 0.25 {
			dt = 0.
		}
		last = time.Now() // ^
		win.Clear(colornames.Black)
		viewCanvas.Clear(color.RGBA{0x0a, 0x0a, 0x0a, 0x0a})
		imd.Clear()
		switch gameState {
		case 0: // In game, will probably change... Not sure
			updateGame(win, dt)
			renderGame(win, viewCanvas, imd)
		case 1: // In menu [?Likely to be separate menus?]
			updateMenu(dt)
			renderMenu(win)
		}

		viewMatrix = pixel.IM. // This centers the camera on the player
					Scaled(player.pos, camZoom).
					Moved(win.Bounds().Center()).
					Moved(pixel.ZV.Sub(player.pos))

		winMatrix := pixel.IM. // This centers the camera on the player
					Scaled(player.pos, camZoom).
					Moved(win.Bounds().Center()).
					Moved(pixel.ZV.Add(player.pos))

		win.SetMatrix(winMatrix)

		imd.Draw(win)
		viewCanvas.Draw(win, viewMatrix)
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
