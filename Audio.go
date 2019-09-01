package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faiface/beep/effects"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	audioCounter = 0              //used for setup of the above arrays
	playingAudio [20]bool         //flags to tell if audio is currently playing or not
	audioAddr    [20]Addr         //addresses for the audio files
	audioName    [20]AudioName    //names for the audio files
	gameAudio    [20]*beep.Buffer //actual audio buffers
	//ARRAY LENGTH ARBITRARILY SET TO 20 FOR TESTING PURPOSES
)

//Addr ... location of the audio file
type Addr struct {
	location string
}

//AudioName ... the name of the audio file
type AudioName struct {
	name string
}

//play audio
func (a *Addr) play(i int) {

	fmt.Sprintln("Grabbing audio file w/ name: " + audioName[i].name) //FOR TESTING PURPOSES
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

//run at startup, preloads audio files into buffers
func loadAudio() {
	dirname := "./Resources/Sound/Audio" //directory of all in game audio files

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
		playingAudio[audioCounter] = false
		audioName[audioCounter].name = file.Name()
		audioAddr[audioCounter].location = dirname + "/" + file.Name()
		audioCounter++
	}
	//sets up the buffer array
	for i := 0; i < audioCounter; i++ {
		gameAudio[i] = audioAddr[i].setupBuffer()
	}

}

//this func is for selecting the audio file to play
func selectAudio(i int) {
	//checks for index < 0
	if i < 0 {
		fmt.Sprintln("Error, desired audio file doesn't exist")
		return
	}

	if playingAudio[i] {
		return //don't want to repeat the same sound over and over again
	}
	//this code only executes if the audio selected isn't already playing
	playingAudio[i] = true
	audioAddr[i].play(i)

}

//pass in the name of the file you want to play and this func will return the index of the file to be played back
func searchAudio(s string) int {
	for i := 0; i < audioCounter; i++ {
		if audioName[i].name == s {
			return i //returns the index for the desired file
		}
	}
	//no filename found
	return -1

}
