package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
* This file deals with the saving the game. It is important and if there are any
* altercations, the save files for any user will not work as this is a parsed
* file system. I'm going to use encryption for this file so people can't easily
* skip a hard level bc they're being pussies using the cryptos package
* -- There really won't be a need to change it most likely as the game is level
* -- based and my plan is just to have it load the level and room using the
* -- current level switching system.
*** FORMAT ***
currentLevel (integer)
currentRoom (integer)
yeah uh, pretty much it for now lol
*/

func saveGame(level int, room int) {
	f, err := os.Create("save.txt") // Create file if it doesn't exist
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encrypt the sucker and write
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	encryptFile("game.dat", []byte(strconv.Itoa(level)+"\n"+strconv.Itoa(room)), "egg")
	// Yes, egg is the password. It doesn't matter because they'd have to decompile
	// a go exe, which is pretty much impossible because it compiles into machine code
	// and no one's gone the extra mile to make it translate back into Go. The only
	// way to get the information out of the save file is for someone to know the
	// type of encryption we're using, and go through the trouble of cracking
	// the password, and in that case, god damn go ahead man, you REALLY don't
	// want to play the current level do you?
}
