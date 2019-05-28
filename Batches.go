package main

import (
	"github.com/faiface/pixel"
)

var (
	batches Batches
)

/*Batches ... All the batches in the game*/
type Batches struct {
	playerBatch *pixel.Batch
}

func loadBatches() {
	batches = Batches{
		pixel.NewBatch(&pixel.TrianglesData{}, spritesheets.playerIdleDownSheet.sheet),
	}
}
