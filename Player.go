package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*Player ... struct for controllable players in the game.*/
type Player struct {
	pos            pixel.Vec
	center         pixel.Vec
	velocity       pixel.Vec
	maxSpeed       float64
	currSpeed      float64
	rotation       float64
	radius         float64
	size           pixel.Vec
	currDir        int // Current direction of moving, 0 W, 1 D, 2 S, 3 A
	canMove        bool
	activeMovement bool
	pic            pixel.Picture
	health         int8
	maxHealth      int8
}

func createPlayer(pos pixel.Vec, cID int, pic pixel.Picture, movable bool) Player { // Player constructor
	size := pixel.V(pic.Bounds().Size().X/float64(len(spritesheets.playerIdleDownSheet.frames)), pic.Bounds().Size().Y)

	return Player{
		pos,
		pixel.ZV,
		pixel.ZV,
		50.0,
		100.0,
		0.0,
		pic.Bounds().Size().Y / 2,
		size,
		1,
		movable,
		false,
		pic,
		100,
		100,
	}

}

func (p *Player) update(win *pixelgl.Window, dt float64) { // Updates player
	if p.canMove {
		p.input(win, dt)
	}
	/*mat := pixel.IM. // This centers the camera on the player
	Scaled(p.pos, camZoom).
	Moved(pixel.ZV.Sub(p.pos)).
	Moved(win.Bounds().Center())*/
	p.center = pixel.V(p.pos.X+(p.size.X/2), p.pos.Y+(p.size.Y/2))
	//win.SetMatrix(mat)
}

func (p *Player) render(win *pixelgl.Window, viewCanvas *pixelgl.Canvas) { // Draws the player
	batches.playerBatch.Clear()
	sprite := pixel.NewSprite(spritesheets.playerIdleDownSheet.sheet, spritesheets.playerIdleDownSheet.frames[p.currDir])
	sprite.Draw(batches.playerBatch, pixel.IM.Rotated(pixel.ZV, p.rotation).Moved(p.center))
	batches.playerBatch.Draw(viewCanvas)
}

func (p *Player) input(win *pixelgl.Window, dt float64) {
	if p.velocity.X > p.maxSpeed {
		p.velocity.X = p.maxSpeed
	} else if p.velocity.X < -1*p.maxSpeed {
		p.velocity.X = -1 * p.maxSpeed
	}
	if p.velocity.Y > p.maxSpeed {
		p.velocity.Y = p.maxSpeed
	} else if p.velocity.Y < -1*p.maxSpeed {
		p.velocity.Y = -1 * p.maxSpeed
	}

	if p.canMove {
		p.velocity = pixel.V(0., 0.)
	}

	if win.Pressed(pixelgl.KeyW) { // Up, 0
		p.currDir = 0

		p.rotation = 0
		p.velocity.Y = p.currSpeed //pixel.V(p.pos.X, p.pos.Y+(p.currSpeed*dt))
	}
	if win.Pressed(pixelgl.KeyD) { // Right, 1
		p.currDir = 1

		p.rotation = 0
		p.velocity.X = p.currSpeed //pixel.V(p.pos.X+(p.currSpeed*dt), p.pos.Y)
	}
	if win.Pressed(pixelgl.KeyS) { // Down, 2
		p.currDir = 2

		p.rotation = 0
		p.velocity.Y = -p.currSpeed //pixel.V(p.pos.X, p.pos.Y-(p.currSpeed*dt))
	}
	if win.Pressed(pixelgl.KeyA) { // Left, 3
		p.currDir = 3

		p.rotation = 0
		p.velocity.X = -p.currSpeed //pixel.V(p.pos.X-(p.currSpeed*dt), p.pos.Y)
	}

	p.pos = pixel.V(p.pos.X+p.velocity.X*dt, p.pos.Y+p.velocity.Y*dt)

	p.isMoving(win)
}

func placeWire(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		fmt.Println("Ok")
	}
}

func (p *Player) isMoving(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyD) {
		p.activeMovement = true
	} else {
		p.activeMovement = false
	}
}
