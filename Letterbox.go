package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func letterBox(win *pixelgl.Window, p Player) {
	windowRatio := winWidth / winHeight
	viewRatio := win.Bounds().W() / win.Bounds().H()
	sizeX := 1.
	sizeY := 1.
	posX := 0.
	posY := 0.
	_ = posX
	_ = posY

	horizontalSpacing := true
	if windowRatio < viewRatio {
		horizontalSpacing = false
	}
	// If horizontalSpacing is true, the black bars will appear on the left and right side.
	// Otherwise, the black bars will appear on the top and bottom.

	if horizontalSpacing {
		sizeX = viewRatio / windowRatio
		posX = (1 - sizeX) / 2.
	} else {
		sizeY = windowRatio / viewRatio
		posY = (1 - sizeY) / 2.
	}

	viewMatrix = pixel.IM.
		Scaled(p.pos, camZoom).
		Moved(pixel.ZV.Sub(p.pos)).
		ScaledXY(pixel.V(posX, posY), pixel.V(sizeX, sizeY))
}
