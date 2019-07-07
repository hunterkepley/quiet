package main

import "github.com/faiface/pixel"

//SoundBubble ... Sound bubbles emitted from sounds
type SoundBubble struct {
	pos      pixel.Vec
	velocity pixel.Vec
	size     pixel.Vec
	pic      pixel.Picture
	sprite   *pixel.Sprite
}

func createSoundBubble(pos pixel.Vec, pic pixel.Picture) SoundBubble {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return SoundBubble{pos, pixel.V(0, 0), size, pic, sprite}
}

//SoundEmitter ... Emitter for sound in the game
type SoundEmitter struct {

}