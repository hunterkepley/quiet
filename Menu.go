package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//Menu ... The main menu of the game and all the UI and functions that belong to it
type Menu struct {
	images []UIImage
}

func createMainMenu() Menu {
	menu := Menu{
		[]UIImage{
			createUIImage(pixel.V(0, 0), menuImages.title),
		},
	}
	go menu.runMusic() // Plays music
	return menu
}

func (m *Menu) update(win *pixelgl.Window, viewCanvas *pixelgl.Canvas) {
	if win.Pressed(pixelgl.KeyEnter) {
		// Set up level
		currentLevel = levels[0]
		currentLevel.setupRoom(&player, viewCanvas)
		gameState = 0

		gameSongIndex := searchMusic("gameMusic.mp3")
		go switchSong(gameSongIndex)
	}
	for i := 0; i < len(m.images); i++ {
		m.images[i].update()
	}
}

func (m *Menu) render(viewCanvas *pixelgl.Canvas) {
	for i := 0; i < len(m.images); i++ {
		m.images[i].render(viewCanvas)
	}
}

func (m *Menu) runMusic() {
	//songs.menuSong.play()
	menuSongIndex := searchMusic("menuMusic.mp3")
	currentSong = menuSongIndex
	music[menuSongIndex].play()
}
