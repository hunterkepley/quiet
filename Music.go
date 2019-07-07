package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	songs Songs
)

//Songs ... Songs in the game
type Songs struct {
	menuSong Music
}

//Music ... Music for the game in a simple-to-use system, must be mp3 for now
type Music struct {
	location string
}

func createMusic(location string) Music {
	return Music{
		location,
	}
}

func (m *Music) play() {
	song, err := os.Open(m.location)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(song)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func loadMusic() {
	songs = Songs{
		createMusic("./Resources/Sound/Music/menuMusic.mp3"),
	}
}
