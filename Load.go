package main

import (
	"strconv"
	"strings"

	"github.com/faiface/pixel/pixelgl"
)

/*
* This file deals with parsing and loading the save file as well as
* sending the engine the correct information for what level/room
* to load. Check the save file for format and stipulations
 */

func loadGame(viewCanvas *pixelgl.Canvas) {
	split := strings.Split(string(decryptFile("Resources/game.dat", "egg")), "\n")
	currentLevelIndex, _ = strconv.Atoi(split[0]) // Set current level index
	currentLevel = levels[currentLevelIndex]      // Set current level to that index
	room, _ := strconv.Atoi(split[1])
	currentLevel.changeRoom(room, &player, viewCanvas) // Set current room
}
