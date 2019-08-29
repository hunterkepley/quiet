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
	gameSong Music
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

	loop := beep.Loop(-1, streamer) //will indefinitley loop the song selected

	done := make(chan bool)
	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func loadMusic() {
	songs = Songs{
		createMusic("./Resources/Sound/Music/menuMusic.mp3"),
		createMusic("./Resources/Sound/Music/gameMusic.mp3"),
	}
}

//for closing currently running songs
func closeSong() {
	speaker.Close()
}

//WIP attempting to simplify the process of closing a song and playing a new one
func switchSong(s string) {
	speaker.Close()
	//kind of ugly but def functional and would allow for switching from menu music to game music and vice versa
	if s == "g" {
		songs.gameSong.play()
	} else if s == "m" {
		songs.menuSong.play()
	}
}
