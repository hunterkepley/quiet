package main

import (
	"github.com/faiface/pixel"
)

var (
	objectImages L1ObjectImages
	menuImages   MenuImages
	soundImages  SoundImages
	enemyImages  EnemyImages
)

//L1ObjectImages ... All the non-spritesheet images in the game in level 1
type L1ObjectImages struct {
	box1         pixel.Picture
	gasBody      pixel.Picture
	gasRoof      pixel.Picture
	gasLeftPole  pixel.Picture
	gasRightPole pixel.Picture
	gasStreet    pixel.Picture
	gasFence     pixel.Picture
	gasLight     pixel.Picture
	concreteWall pixel.Picture
	woodFloor    pixel.Picture
	trashCan     pixel.Picture
}

//SoundImages ... All the visible sound images
type SoundImages struct {
	playerSoundWaveTL pixel.Picture
	playerSoundWaveTR pixel.Picture
	playerSoundWaveBL pixel.Picture
	playerSoundWaveBR pixel.Picture
	playerSoundWaveB  pixel.Picture
	playerSoundWaveT  pixel.Picture
	playerSoundWaveL  pixel.Picture
	playerSoundWaveR  pixel.Picture
}

//MenuImages ... All the menu images that aren't animated
type MenuImages struct {
	title pixel.Picture
}

//LarvaImages ... All the larva enemy images in the game
type LarvaImages struct {
	stillLeft pixel.Picture
}

//EnemyImages ... All the enemy images in the game that aren't animated
type EnemyImages struct {
	larvaImages LarvaImages
	eye         pixel.Picture
}

func loadObjectImages() {
	objectImages = L1ObjectImages{
		loadPicture("./Resources/Art/Objects/Scenery/l1/box1.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_body.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_roof.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_left_pole.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_right_pole.png"),
		loadPicture("./Resources/Art/Objects/Backgrounds/l1/street1.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_fence.png"),
		loadPicture("./Resources/Art/Objects/Buildings/l1/gas_light.png"),
		loadPicture("./Resources/Art/Objects/Walls/l1/concrete_wall.png"),
		loadPicture("./Resources/Art/Objects/Backgrounds/l1/floor.png"),
		loadPicture("./Resources/Art/Objects/Scenery/l1/trash_can.png"),
	}
}

func loadSoundImages() {
	soundImages = SoundImages{
		loadPicture("./Resources/Art/Player/sound_wave_tl.png"),
		loadPicture("./Resources/Art/Player/sound_wave_tr.png"),
		loadPicture("./Resources/Art/Player/sound_wave_bl.png"),
		loadPicture("./Resources/Art/Player/sound_wave_br.png"),
		loadPicture("./Resources/Art/Player/sound_wave_b.png"),
		loadPicture("./Resources/Art/Player/sound_wave_t.png"),
		loadPicture("./Resources/Art/Player/sound_wave_l.png"),
		loadPicture("./Resources/Art/Player/sound_wave_r.png"),
	}
}

func loadMenuImages() {
	menuImages = MenuImages{
		loadPicture("./Resources/Art/UI/MM/title.png"),
	}
}

func loadEnemyImages() {
	enemyImages = EnemyImages{
		LarvaImages{
			loadPicture("./Resources/Art/Enemies/Larva/still_left.png"),
		},
		loadPicture("./Resources/Art/Enemies/eye.png"),
	}
}
