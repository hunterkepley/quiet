package main

import (
	"github.com/faiface/pixel"
)

var (
	playerBatches PlayerBatches
)

/*Batches ... All the batches in the game*/
type PlayerBatches struct {
	playerIdleRightBatch *pixel.Batch
	playerIdleUpBatch    *pixel.Batch
}

func loadPlayerBatches() {
	playerBatches = PlayerBatches{
		pixel.NewBatch(&pixel.TrianglesData{}, spritesheets.playerIdleRightSheet.sheet),
		pixel.NewBatch(&pixel.TrianglesData{}, spritesheets.playerIdleUpSheet.sheet),
	}
}
