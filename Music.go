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
	songs       Songs
	currentSong = 0
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
	//NEEDS TO BE CHANGED, FOR AUDIO TO WORK SPEAKER CAN'T BE CLOSED
	//NEED TO END CURRENT STREAMER
	speaker.Close()
}

//WIP attempting to simplify the process of closing a song and playing a new one
func switchSong(i int) {

	//fmt.Println("song id: ", i)

	//kind of ugly but def functional and would allow for switching from menu music to game music and vice versa
	if i == currentSong { //this if/else block added as a bugfix to speaker overload bug as a way to check before closing and restarting the speaker
		//literally just don't do anything
	} else {

		if i == 1 { //id 1 is normal game song
			currentSong = i
			closeSong()
			songs.gameSong.play()
		} else if i == 0 { //id 0 is menu song
			currentSong = i
			closeSong()
			songs.menuSong.play()
		} else { //id doesn't exist, some error has occurred
			//fmt.Println("err")
			return
		}
	}
}
