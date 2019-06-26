package main

import (
	"github.com/faiface/pixel/pixelgl"
)

var (
	mainMenu Menu
)

// This file handles the main menu for now. Will make another one to handle the pause menu most likely? Maybe just make it overlay the game, maybe not, I don't know anymore

func renderMenu(win *pixelgl.Window, viewCanvas *pixelgl.Canvas) {
	mainMenu.render(viewCanvas)
}

func updateMenu(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, dt float64) {
	mainMenu.update(win, viewCanvas)
}
