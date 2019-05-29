package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var (
	windowBounds = pixel.ZV
)

func renderGame(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, imd *imdraw.IMDraw) {
	player.render(win, viewCanvas)
	testBox.render(win, viewCanvas)
}

func updateGame(win *pixelgl.Window, dt float64) {
	player.update(win, dt)

	testBox.update()

	// This is pretty badly done, but it does the trick for making the stars not decrease their bounds
	if win.Bounds().W() > windowBounds.X {
		windowBounds.X = win.Bounds().W()
	}
	if win.Bounds().H() > windowBounds.Y {
		windowBounds.Y = win.Bounds().H()
	}
}
