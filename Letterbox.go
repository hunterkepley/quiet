package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func letterBox(win *pixelgl.Window) {
	sizeX := 1.
	sizeY := 1.

	if win.Bounds().H()-winHeight > win.Bounds().W()-winWidth {
		sizeX = win.Bounds().W() / winWidth
		sizeY = win.Bounds().W() / winWidth
	} else {
		sizeX = win.Bounds().H() / winHeight
		sizeY = win.Bounds().H() / winHeight
	}

	viewMatrix = pixel.IM.
		Moved(pixel.V(win.Bounds().Center().X/camZoom, win.Bounds().Center().Y/camZoom)).
		ScaledXY(pixel.V(win.Bounds().Center().X/camZoom, win.Bounds().Center().Y/camZoom), pixel.V(sizeX, sizeY))
}
