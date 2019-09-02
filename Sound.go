package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//SoundWave ... Sound waves emitted from sounds
type SoundWave struct {
	pos                 pixel.Vec
	startPos            pixel.Vec // Starting position for enemies to go to
	center              pixel.Vec
	velocity            pixel.Vec // 0 for no change, 1 for  change via dB level [-1 for reverse]
	size                pixel.Vec
	pic                 pixel.Picture
	sprite              *pixel.Sprite
	dB                  float64          // Speed/strength of the sound
	startingDB          float64          // What dB it started with
	depletionRate       float64          // How quickly the sound wave loses dB
	passedThrough       []int            // Object indexes the soundwave passed through so it doesn't glitch and delete
	trail               []SoundWaveTrail // THE TRAILS FOR THE SOUND WAVE
	createTrailTimer    float64          // How quickly a new trail particle is made
	createTrailTimerMax float64
}

//SoundWaveTrail ... Trails for the soundwaves dude
type SoundWaveTrail struct {
	pos            pixel.Vec
	size           pixel.Vec
	animationSheet Spritesheet
	animation      Animation
	center         pixel.Vec
	velocity       pixel.Vec
	sprite         *pixel.Sprite
	animationEnded bool
	end            bool
	speed          float64
}

func createSoundWaveTrail(pos pixel.Vec, velocity pixel.Vec, animationSheet Spritesheet, speed float64) SoundWaveTrail {
	animationSpeed := 0.25
	sprite := pixel.NewSprite(animationSheet.sheet, animationSheet.sheet.Bounds())
	size := pixel.V(animationSheet.sheet.Bounds().Size().X/float64(animationSheet.numberOfFrames), animationSheet.sheet.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)
	return SoundWaveTrail{
		pos,
		size,
		animationSheet,
		createAnimation(animationSheet, animationSpeed),
		pixel.ZV,
		velocity,
		sprite,
		false,
		false,
		speed,
	}
}

func (t *SoundWaveTrail) update(dt float64) {
	if !t.end {
		t.pos = pixel.V(t.pos.X+((t.velocity.X*t.speed)*dt), t.pos.Y+((t.velocity.Y*t.speed)*dt))
		t.center = pixel.V(t.pos.X+(t.size.X/2), t.pos.Y+(t.size.Y/2))
		if t.animation.frameNumber >= t.animation.frameNumberMax-1 {
			t.animationEnded = true
		}
		if t.animationEnded && t.animation.frameNumber < t.animation.frameNumberMax-1 {
			t.end = true
		}
	}
}

func (t *SoundWaveTrail) render(viewCanvas *pixelgl.Canvas) {
	if !t.end {
		mat := pixel.IM.
			Moved(t.center).
			Scaled(t.center, imageScale)

		*t.sprite = t.animation.animate(dt)
		t.sprite.Draw(soundWaveBatches.soundWaveBTrailBatch, mat)
	}
}

func createSoundWave(pos pixel.Vec, pic pixel.Picture, velocity pixel.Vec, dB float64, depletionRate float64, startPos pixel.Vec) SoundWave {
	sprite := pixel.NewSprite(pic, pic.Bounds())
	size := pixel.V(pic.Bounds().Size().X, pic.Bounds().Size().Y)
	size = pixel.V(size.X*imageScale, size.Y*imageScale)

	trailTimer := 0.5

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
		[]SoundWaveTrail{},
		trailTimer,
		trailTimer,
	}
}

func (w *SoundWave) update(dt float64) {
	w.pos = pixel.V(w.pos.X+((w.velocity.X*w.startingDB)*dt), w.pos.Y+((w.velocity.Y*w.startingDB)*dt))
	w.center = pixel.V(w.pos.X+(w.size.X/2), w.pos.Y+(w.size.Y/2))
	w.objectCollision()
	for i := 0; i < len(w.trail); i++ {
		w.trail[i].update(dt)
	}
	if w.createTrailTimer > 0. {
		w.createTrailTimer -= 1 * dt
	} else {
		w.createTrailTimer = w.createTrailTimerMax
		w.trail = append(w.trail, createSoundWaveTrail(w.pos, w.velocity, playerSpritesheets.soundWaveBTrailSheet, w.dB/1.5))
	}
}

func (w *SoundWave) render(viewCanvas *pixelgl.Canvas) {
	// Don't be lazy clean this up it's disgusting and you should feel bad
	mat := pixel.IM.
		Moved(w.center).
		Scaled(w.center, imageScale)
	if w.velocity == pixel.V(0, -1) {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[0], mat)
	} else if w.velocity == pixel.V(-1, -1) {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[1], mat)
	} else if w.velocity == pixel.V(1, -1) {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[2], mat)
	} else if w.velocity == pixel.V(-1, 0) {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[3], mat)
	} else if w.velocity == pixel.V(1, 0) {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[4], mat)
	} else if w.velocity == pixel.V(0, 1) {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[5], mat)
	} else if w.velocity == pixel.V(-1, 1) {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[6], mat)
	} else {
		w.sprite.Draw(soundWaveBatches.soundWaveBatches[7], mat)
	}
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
	soundWaveBatches.soundWaveBTrailBatch.Clear()
	for _, j := range soundWaveBatches.soundWaveBatches {
		j.Clear()
	}
	for i := 0; i < len(s.waves); i++ {
		s.waves[i].render(viewCanvas)
		for j := 0; j < len(s.waves[i].trail); j++ {
			s.waves[i].trail[j].render(viewCanvas)
		}
	}
	soundWaveBatches.soundWaveBTrailBatch.Draw(viewCanvas)
	for _, j := range soundWaveBatches.soundWaveBatches {
		j.Draw(viewCanvas)
	}
}
