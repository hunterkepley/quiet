package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/**
 * This file deals with parsing options n shit
 * from a file
 * lmao
 */

//Options ... All the options for the game
type Options struct {
	videoOptions VideoOptions
}

//VideoOptions ... The video options for the game, like vsync
type VideoOptions struct {
	vsync bool
}

func (o *Options) loadOptions() {
	file, err := os.Open("Resources/settings.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	split := strings.Split(string(b), "\n")
	_ = split
}
