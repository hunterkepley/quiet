package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	windowBounds = pixel.ZV
)

func renderGame(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, imd *imdraw.IMDraw, dt float64) {
	for i := 0; i < len(backgroundObjects); i++ {
		backgroundObjects[i].render(viewCanvas, imd, player)

	}
	player.render(win, viewCanvas, dt)
	for i := 0; i < len(foregroundObjects); i++ {
		foregroundObjects[i].render(viewCanvas, imd, player)
		if foregroundObjects[i].hitboxes {
			foregroundObjects[i].renderHitboxes(imd, player)
		}
	}
	if currentLevel.rooms[currentLevel.currentRoomIndex].hasRain {
		renderSplashes(viewCanvas)
	}
	if currentLevel.rooms[currentLevel.currentRoomIndex].hasRain {
		renderRain(viewCanvas)
		if drawRainDeadzones {
			imd.Color = colornames.Cyan
			width := 1.
			for i := 0; i < len(currentLevel.rooms[currentLevel.currentRoomIndex].rainDeadZones); i++ {
				imd.Push(currentLevel.rooms[currentLevel.currentRoomIndex].rainDeadZones[i].Min, currentLevel.rooms[currentLevel.currentRoomIndex].rainDeadZones[i].Max)
				imd.Rectangle(width)
			}
		}
	}
}

func updateGame(win *pixelgl.Window, dt float64) {
	player.update(win, dt)

	if win.Pressed(pixelgl.KeyE) {
		currentLevel.changeRoom(1, &player)
	}

	currentLevel.updateRoom(&player, dt)

	if len(backgroundObjects) >= 1 {
		for i := 0; i < len(backgroundObjects); i++ {
			backgroundObjects[i].update(&player)
			if backgroundObjects[i].inFrontOfPlayer {
				foregroundObjects = append(foregroundObjects, backgroundObjects[i])
				backgroundObjects = append(backgroundObjects[:i], backgroundObjects[i+1:]...)
			}
		}
	}

	if len(foregroundObjects) >= 1 {
		for i := 0; i < len(foregroundObjects); i++ {
			foregroundObjects[i].update(&player)
			if !foregroundObjects[i].inFrontOfPlayer {
				backgroundObjects = append(backgroundObjects, foregroundObjects[i])
				foregroundObjects = append(foregroundObjects[:i], foregroundObjects[i+1:]...)
			}
		}
	}

	// This is pretty badly done, but it does the trick for making the stars not decrease their bounds
	if win.Bounds().W() > windowBounds.X {
		windowBounds.X = win.Bounds().W()
	}
	if win.Bounds().H() > windowBounds.Y {
		windowBounds.Y = win.Bounds().H()
	}
}
