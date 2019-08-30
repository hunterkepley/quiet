package main

import (
	"math"

	"github.com/faiface/pixel"
)

func collisionCheck(x pixel.Rect, y pixel.Rect) bool {
	if x.Min.X < y.Max.X &&
		x.Max.X > y.Min.X &&
		x.Min.Y < y.Max.Y &&
		x.Max.Y > y.Min.Y {
		return true
	}
	return false
}

func calculateDistance(a pixel.Vec, b pixel.Vec) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

func circlularCollisionCheck(r1 float64, r2 float64, d float64) bool {
	if d < r1+r2 {
		return true
	}
	return false
}
