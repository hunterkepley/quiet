package main

import "github.com/faiface/pixel"

func collisionCheck(x pixel.Rect, y pixel.Rect) bool {
	if x.Min.X < y.Max.X &&
		x.Max.X > y.Min.X &&
		x.Min.Y < y.Max.Y &&
		x.Max.Y > y.Min.Y {
		return true
	}
	return false
}
