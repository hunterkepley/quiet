package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//SoundWave ... Sound waves emitted from sounds
type SoundWave struct {
	pos      pixel.Vec
	center   pixel.Vec
	velocity pixel.Vec // 0 for no change, 1 for  change via dB level [-1 for reverse]
	size     pixel.Vec
	pic      pixel.Picture
	sprite   *pixel.Sprite
	dB       float64 // Speed/strength of the sound
}

func createSoundWave(pos pixel.Vec, pic pixel.Picture, velocity pixel.Vec, dB float64) SoundWave {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return SoundWave{
		pos,
		pixel.ZV,
		velocity,
		size,
		pic,
		sprite,
		dB}
}

func (w *SoundWave) update(dt float64) {
	w.center = pixel.V(w.pos.X+(w.size.X/2), w.pos.Y+(w.size.Y/2))
	w.pos = pixel.V(w.pos.X+((w.velocity.X*w.dB)*dt), w.pos.Y+((w.velocity.Y*w.dB)*dt))
}

func (w *SoundWave) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(w.center).
		Scaled(w.center, imageScale)
	w.sprite.Draw(viewCanvas, mat)
}

//SoundEmitter ... Emitter for sound in the game
type SoundEmitter struct {
	pos   pixel.Vec
	waves []SoundWave
}

func (s *SoundEmitter) emit(dB float64) {
	s.waves = append(s.waves, createSoundWave(s.pos, soundImages.playerSoundWaveTL, pixel.V(-1, 1), dB))  // Top left
	s.waves = append(s.waves, createSoundWave(s.pos, soundImages.playerSoundWaveTR, pixel.V(1, 1), dB))   // Top right
	s.waves = append(s.waves, createSoundWave(s.pos, soundImages.playerSoundWaveBL, pixel.V(-1, -1), dB)) // Bottom left
	s.waves = append(s.waves, createSoundWave(s.pos, soundImages.playerSoundWaveBR, pixel.V(1, -1), dB))  // Bottom right
}
