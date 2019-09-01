package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//SoundWave ... Sound waves emitted from sounds
type SoundWave struct {
	pos           pixel.Vec
	startPos      pixel.Vec // Starting position for enemies to go to
	center        pixel.Vec
	velocity      pixel.Vec // 0 for no change, 1 for  change via dB level [-1 for reverse]
	size          pixel.Vec
	pic           pixel.Picture
	sprite        *pixel.Sprite
	dB            float64 // Speed/strength of the sound
	startingDB    float64 // What dB it started with
	depletionRate float64 // How quickly the sound wave loses dB
	passedThrough []int   // Object indexes the soundwave passed through so it doesn't glitch and delete
}

func createSoundWave(pos pixel.Vec, pic pixel.Picture, velocity pixel.Vec, dB float64, depletionRate float64, startPos pixel.Vec) SoundWave {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)

	return SoundWave{
		pixel.V(pos.X-size.X/2, pos.Y-size.Y/2),
		startPos,
		pixel.ZV,
		velocity,
		size,
		pic,
		sprite,
		dB,
		dB,
		depletionRate,
		[]int{},
	}
}

func (w *SoundWave) update(dt float64) {
	w.pos = pixel.V(w.pos.X+((w.velocity.X*w.startingDB)*dt), w.pos.Y+((w.velocity.Y*w.startingDB)*dt))
	w.center = pixel.V(w.pos.X+(w.size.X/2), w.pos.Y+(w.size.Y/2))
	w.objectCollision()
}

func (w *SoundWave) render(viewCanvas *pixelgl.Canvas) {
	mat := pixel.IM.
		Moved(w.center).
		Scaled(w.center, imageScale)
	w.sprite.Draw(viewCanvas, mat)
}

func (w *SoundWave) objectCollision() {
	for i := range foregroundObjects {
		o := foregroundObjects[i]
		if o.soundCollidable {
			if w.pos.X < o.pos.X+o.size.X &&
				w.pos.X+w.size.X > o.pos.X &&
				w.pos.Y < o.pos.Y+o.size.Y &&
				w.pos.Y+w.size.Y > o.pos.Y {

				w.reflect(o, i)
			}
		}
	}
	for i := range backgroundObjects {
		o := backgroundObjects[i]
		if o.soundCollidable {
			if w.pos.X < o.pos.X+o.size.X &&
				w.pos.X+w.size.X > o.pos.X &&
				w.pos.Y < o.pos.Y+o.size.Y &&
				w.pos.Y+w.size.Y > o.pos.Y {

				w.reflect(o, i)
			}
		}
	}
}

func (w *SoundWave) reflect(o Object, i int) {
	con := true
	for j := range w.passedThrough {
		if w.passedThrough[j] == i {
			con = false
		}
	}
	if con {
		w.passedThrough = append(w.passedThrough, i)
		w.dB -= o.dBDiminisher
		w.startingDB = w.dB
	}
}

//SoundEmitter ... Emitter for sound in the game
type SoundEmitter struct {
	pos   pixel.Vec
	waves []SoundWave
	//audio *beep.Buffer
}

func createSoundEmitter(pos pixel.Vec) SoundEmitter {
	return SoundEmitter{
		pos,
		[]SoundWave{},
		//gameAudio[0],
	}
}

func (s *SoundEmitter) emit(dB float64, depletionRate float64) {
	offset := 12.  // Offset for the corner waves
	offsetS := 16. // Offset for the top, bottom, left, right waves

	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X-offset, s.pos.Y+offset), soundImages.playerSoundWaveTL, pixel.V(-1, 1), dB, depletionRate, s.pos))  // Top left
	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X+offset, s.pos.Y+offset), soundImages.playerSoundWaveTR, pixel.V(1, 1), dB, depletionRate, s.pos))   // Top right
	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X-offset, s.pos.Y-offset), soundImages.playerSoundWaveBL, pixel.V(-1, -1), dB, depletionRate, s.pos)) // Bottom left
	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X+offset, s.pos.Y-offset), soundImages.playerSoundWaveBR, pixel.V(1, -1), dB, depletionRate, s.pos))  // Bottom right
	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X, s.pos.Y+offsetS), soundImages.playerSoundWaveT, pixel.V(0, 1), dB, depletionRate, s.pos))          // Top
	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X, s.pos.Y-offsetS), soundImages.playerSoundWaveB, pixel.V(0, -1), dB, depletionRate, s.pos))         // Bottom
	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X+offsetS, s.pos.Y), soundImages.playerSoundWaveR, pixel.V(1, 0), dB, depletionRate, s.pos))          // Right
	s.waves = append(s.waves, createSoundWave(pixel.V(s.pos.X-offsetS, s.pos.Y), soundImages.playerSoundWaveL, pixel.V(-1, 0), dB, depletionRate, s.pos))         // Left

	footstepIndex := searchAudio("footstep.mp3") //should look up file based on the provided string
	go selectAudio(footstepIndex)           //should play audio
}

func (s *SoundEmitter) update(pos pixel.Vec, dt float64) {
	s.pos = pos
	for i := 0; i < len(s.waves); i++ {
		s.waves[i].update(dt)

		s.waves[i].dB -= s.waves[i].depletionRate * dt
		// Screen edge collision detection/response
		// These have to be separate to avoid crashing
		if s.waves[i].center.X-s.waves[i].size.X/2 < 0. || s.waves[i].center.X+s.waves[i].size.X/2 > winWidth { // Left / Right
			s.waves = append(s.waves[:i], s.waves[i+1:]...)
		} else if s.waves[i].center.Y-s.waves[i].size.Y/2 < 0. || s.waves[i].center.Y+s.waves[i].size.Y/2 > winHeight { // Bottom / Top
			s.waves = append(s.waves[:i], s.waves[i+1:]...)
		} else if s.waves[i].dB <= 0. { // Remove if no dB
			s.waves = append(s.waves[:i], s.waves[i+1:]...)
		}
	}
}

func (s *SoundEmitter) render(viewCanvas *pixelgl.Canvas) {
	for i := 0; i < len(s.waves); i++ {
		s.waves[i].render(viewCanvas)
	}
}
