package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	currentSong  = 0
	musicCounter = 0 //used for seting up the audio
	//musicAddr    [20]Music     //addresses for music files
	//musicName    [20]MusicName //names of the music files
	music []Music //literally music
	//ARBITRARILY SET ARRAY LENGTH, CAN ADJUST LATER
)

//Songs ... Songs in the game
type Songs struct {
	menuSong Music
	gameSong Music
}

//Music ... Music for the game in a simple-to-use system, must be mp3 for now
type Music struct {
	location string
	name     string
}

//MusicName ... Names for the music files
type MusicName struct {
	name string
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

	volume := &effects.Volume{
		Streamer: loop,
		Base:     2,
		Volume:   -4.5,
		Silent:   true,
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(volume, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func loadMusic() {
	dirname := "./Resources/Sound/Music"

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	//set up flag, name, and address
	for _, file := range files {
		name := file.Name()
		location := dirname + "/" + file.Name()

		music = append(music, Music{
			location,
			name,
		})

		//musicName[musicCounter].name = file.Name()
		//musicAddr[musicCounter].location = dirname + "/" + file.Name()
		musicCounter++
	}
}

func closeSong() {
	speaker.Close()
}

//literally just change the song 4head
func switchSong(i int) {

	if i == currentSong {
		return //literally just don't do anything
	}

	currentSong = i
	closeSong()
	music[i].play()
}

//pass in file name, index for file returned
func searchMusic(s string) int {
	for i := 0; i < musicCounter; i++ {
		if music[i].name == s {
			return i //returns the index for the desired song
		}
	}
	//no filename found
	return -1
}
