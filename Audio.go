package main

import (
	"log"
	"os"

	"github.com/faiface/beep/effects"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	dirname      = "./Resources/Sound/Audio" //directory of all in game audio files
	playingAudio []bool                      //flags to tell if audio is currently playing or not
	audioAddr    []Addr                      //addresses for the audio files
	audioName    []AudioName                 //names for the audio files
	gameAudio    []*beep.Buffer              //actual audio buffers
)

//Addr ... location of the audio file
type Addr struct {
	location string
}

//AudioName ... the name of the audio file
type AudioName struct {
	name string
}

func (a *Addr) play(i int) {

	buffer := gameAudio[i]

	sound := buffer.Streamer(0, buffer.Len())

	volume := &effects.Volume{
		Streamer: sound,
		Base:     2,
		Volume:   -3,
		Silent:   false,
	}

	done := make(chan bool) //prob wont need, will leave here as reminder just in case
	speaker.Play(volume, beep.Callback(func() {
		//make flag false so that audio can be played again
		playingAudio[i] = false
		done <- true
	}))

	<-done

}

//setup the buffers for playing audio
func (a *Addr) setupBuffer() *beep.Buffer {
	audio, err := os.Open(a.location)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(audio)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	return buffer
}

func loadAudio() {
	loadCounter := 0 //used for setup of the above arrays

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	//sets up the flag, name, and address arrays
	for _, file := range files {
		playingAudio[loadCounter] = false
		audioName[loadCounter].name = file.Name()
		audioAddr[loadCounter].location = dirname + "/" + file.Name()
		loadCounter++
	}
	//sets up the buffer array
	for i := 0; i < loadCounter; i++ {
		gameAudio[i] = audioAddr[i].setupBuffer()
	}

}

func selectAudio(i int) {
	if playingAudio[i] {
		return //don't want to repeat the same sound over and over again
	}
	//this code only executes if the audio selected isn't already playing
	playingAudio[i] = true
	audioAddr[i].play(i)

}
