package main

import (
	"log"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	soundBites   Audio
	playingAudio []bool
)

//Audio ... bites that are played at different points during the game
type Audio struct {
	footstep Addr
	wormSlam Addr
}

//Addr ... locations of the audio bites
type Addr struct {
	location string
}

func createAudio(location string) Addr {
	return Addr{
		location,
	}
}

func (a *Addr) play(i int) {
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

	sound := buffer.Streamer(0, buffer.Len())

	done := make(chan bool) //prob wont need, will leave here as reminder just in case
	speaker.Play(sound, beep.Callback(func() {
		//make flag false so that audio can be played again
		playingAudio[i] = false
		done <- true
	}))

	<-done

}

func loadAudio() {
	soundBites = Audio{
		createAudio("./Resources/Sound/Audio/footstep.mp3"),
		createAudio("./Resources/Sound/Audio/slam.mp3"),
	}
	playingAudio[0] = false
	playingAudio[1] = false
}

func selectAudio(i int) {
	if playingAudio[i] {
		return //don't want to repeat the same sound over and over again
	}
	//this code only executes if the audio selected isn't already playing
	playingAudio[i] = true
	if i == 0 { //audio id 0 footstep
		soundBites.footstep.play(0)
	} else if i == 1 { //audio id 1 worm slam
		soundBites.wormSlam.play(1)
	}

}
