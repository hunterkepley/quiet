package main

import (
	"github.com/faiface/pixel"
)

var (
	playerBatches    PlayerBatches
	objectBatches    ObjectBatches
	soundWaveBatches SoundWaveBatches
)

// PlayerBatches ... All the batches in the game
type PlayerBatches struct {
	playerIdleRightBatch *pixel.Batch
	playerIdleUpBatch    *pixel.Batch
	playerIdleDownBatch  *pixel.Batch
	playerIdleLeftBatch  *pixel.Batch
}

// ObjectBatches .. All the objects in levels batches
type ObjectBatches struct {
	rainBatch       *pixel.Batch
	rainSplashBatch *pixel.Batch
}

// SoundWaveBatches ... The batches for the sound waves (waves and trails)
type SoundWaveBatches struct {
	soundWaveBTrailBatch *pixel.Batch
}

func loadPlayerBatches() {
	playerBatches = PlayerBatches{
		pixel.NewBatch(&pixel.TrianglesData{}, playerSpritesheets.playerIdleRightSheet.sheet),
		pixel.NewBatch(&pixel.TrianglesData{}, playerSpritesheets.playerIdleUpSheet.sheet),
		pixel.NewBatch(&pixel.TrianglesData{}, playerSpritesheets.playerIdleDownSheet.sheet),
		pixel.NewBatch(&pixel.TrianglesData{}, playerSpritesheets.playerIdleLeftSheet.sheet),
	}
}

func loadObjectBatches() {
	objectBatches = ObjectBatches{
		pixel.NewBatch(&pixel.TrianglesData{}, objectSpritesheets.rainSheet.sheet),
		pixel.NewBatch(&pixel.TrianglesData{}, objectSpritesheets.rainSplashSheet.sheet),
	}
}

func loadSoundWaveBatches() {
	soundWaveBatches = SoundWaveBatches{
		pixel.NewBatch(&pixel.TrianglesData{}, playerSpritesheets.soundWaveBTrailSheet.sheet),
	}
}
