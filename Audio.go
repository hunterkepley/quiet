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
	audioCounter = 0     //used for setup of the above arrays
	audio        []Audio //instance of audio struct
)

//Audio ... location of file, name of file, flag for if file is playing or not, buffer that is holding the file for playback
type Audio struct {
	location string
	name     string
	flag     bool
	buffer   *beep.Buffer
}

//play audio
func (a *Audio) play() {

	fmt.Sprintln("Grabbing audio file w/ name: " + a.name) //FOR TESTING PURPOSES
	buffer := a.buffer

	sound := buffer.Streamer(0, buffer.Len())

	volume := &effects.Volume{
		Streamer: sound,
		Base:     2,
		Volume:   -3,
		Silent:   true,
	}

	done := make(chan bool) //prob wont need, will leave here as reminder just in case
	speaker.Play(volume, beep.Callback(func() {
		//make flag false so that audio can be played again
		a.flag = false
		done <- true
	}))

	<-done

}

//setup the buffers for playing audio
func setupBuffer(location string) *beep.Buffer {
	audio, err := os.Open(location)
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
		flag := false
		name := file.Name()
		location := dirname + "/" + file.Name()
		buffer := setupBuffer(location)

		audio = append(audio, Audio{
			location,
			name,
			flag,
			buffer,
		})

		audioCounter++
	}
}

//this func is for selecting the audio file to play
func selectAudio(i int) {
	//checks for index < 0
	if i < 0 || i > audioCounter {
		fmt.Sprintln("Error, desired audio file doesn't exist")
		return
	}

	if audio[i].flag {
		return //don't want to repeat the same sound over and over again
	}
	//this code only executes if the audio selected isn't already playing
	audio[i].flag = true
	audio[i].play()

}

//pass in the name of the file you want to play and this func will return the index of the file to be played back
func searchAudio(s string) int {
	for i := 0; i < audioCounter; i++ {
		if audio[i].name == s {
			return i //returns the index for the desired file
		}
	}
	//no filename found
	return -1

}
