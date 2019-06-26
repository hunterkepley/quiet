package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//Menu ... The main menu of the game and all the UI and functions that belong to it
type Menu struct {
	images []UIImage
}

func createMainMenu() Menu {
	return Menu{
		[]UIImage{
			createUIImage(pixel.V(0, 0), menuImages.title),
		},
	}
}

func (m *Menu) update(win *pixelgl.Window, viewCanvas *pixelgl.Canvas) {
	if win.Pressed(pixelgl.KeyEnter) {
		// Set up level
		currentLevel = levels[0]
		currentLevel.setupRoom(&player, viewCanvas)
		gameState = 0
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

func (m *Menu) music() {
	// TODO: make the music a system like Images or Spritesheets or Animations
	music, err := os.Open("./Resources/Sound/Music/menuMusic.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(music)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
}
