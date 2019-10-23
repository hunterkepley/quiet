package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	levels = []Level{}
)

func loadLevels() {
	levels = []Level{
		createLevel( // First level [outside gas dealer]
			[]Room{
				createRoom( // Outside 1
					[]Object{ // Objects in room
						createObject(pixel.V(0., 0.), l1ObjectImages.gasStreet, 1., true, false, false, false, 0.),
						createObject(pixel.V(0., 0.), l1ObjectImages.gasFence, 1., true, false, false, true, 100.),
						createObject(pixel.V(0., 0.), l1ObjectImages.gasBody, 1., true, false, false, true, 100.),
						createObject(pixel.V(0., 0.), l1ObjectImages.gasRoof, 1., false, true, false, false, 0.),
						createObject(pixel.V(0., 0.), l1ObjectImages.gasLight, 1., false, true, false, false, 0.),
						createObject(pixel.V(518., 352.), l1ObjectImages.gasLeftPole, 15., false, false, true, true, 100.),
						createAnimatedObject(pixel.V(654., 411.), l1ObjectSpritesheets.trashCanSheet, 0.03, 1., false, true, true, true, 0.),
						createObject(pixel.V(764., 352.), l1ObjectImages.gasRightPole, 15., false, false, true, true, 100.),
					},
					[]Enemy{ // Enemies in the room

					},
					[]Entrance{ // Entrances in the room
						createEntrance(pixel.V(340, 421), pixel.V(70, 95), 5., 1, -1, pixelgl.KeyE),
					},
					pixel.V(50., 50.), // Player starting position
					true,              // Has rain
					[]pixel.Rect{ // Rain dead zones
						pixel.R(223, 420, 285, 636),
						pixel.R(282, 424, 461, 589),
						pixel.R(461, 424, 723, 533),
						pixel.R(509, 354, 560, 504),
						pixel.R(752, 354, 804, 504),
						pixel.R(495, 501, 819, 543),
						pixel.R(396, 636, 669, 663),
						pixel.R(728, 553, 768, 570),
						pixel.R(722, 570, 754, 590),
						pixel.R(713, 590, 730, 600),
						pixel.R(708, 600, 723, 612),
						pixel.R(692, 605, 711, 628),
						pixel.R(680, 620, 697, 642),
						pixel.R(667, 632, 682, 653),
						pixel.R(357, 592, 420, 662),
						pixel.R(369, 427, 493, 593),
						pixel.R(0, 237, 1024, 255),
						pixel.R(409, 576, 432, 600),
						pixel.R(239, 636, 362, 678),
						pixel.R(356, 674, 691, 686),
						pixel.R(502, 354, 759, 434),
					},
					stormShader, // Shader
					l1r1,        // Function for the room
					false,       // Has sound waves
				),
				createRoom( // Inside 1
					[]Object{
						createObject(pixel.V(28., 639.), l1ObjectImages.backWall1, 1., false, false, true, true, 100.),
						createObject(pixel.V(28., 0.), l1ObjectImages.bottomWall1, 2., false, false, true, true, 100.),
						createObject(pixel.V(0., 0.), l1ObjectImages.wall1, 1., false, false, true, true, 0.),
						createObject(pixel.V(996., 0.), l1ObjectImages.wall1, 1., false, false, true, true, 0.),
						createObject(pixel.V(0., 0.), l1ObjectImages.floor1, 1., true, false, false, false, 0.),
						createObject(pixel.V(100., 180.), l1ObjectImages.box1, 2., false, false, true, true, 15.),
						createObject(pixel.V(200., 80.), l1ObjectImages.box1, 2., false, false, true, true, 15.),
					},
					[]Enemy{
						createEnemy(pixel.V(300., 50.), enemyImages.larvaImages.stillLeft, 1., 20., 5., 0.15, 0.5, 0.1, 5., 80., 10, 30),
					},
					[]Entrance{
						createEntrance(pixel.V(150, 150), pixel.V(50, 50), 5., 2, -1, pixelgl.KeyE),
					},
					pixel.V(100., 100.),
					false,
					[]pixel.Rect{},
					grayscaleShader,
					l1r2,
					true,
				),
				createRoom( // Inside 2? Testing room for now
					[]Object{
						createObject(pixel.V(100., 180.), l1ObjectImages.box1, 2., false, false, true, true, 15.),
					},
					[]Enemy{},
					[]Entrance{},
					pixel.V(100., 100.),
					false,
					[]pixel.Rect{},
					redShader1,
					l1r3,
					true,
				),
			},
		),
	}
}

// Functions for rooms

// Level 1
// Room 1
func l1r1(player *Player) {
	// Collision against gas station
	if player.pos.Y > 422 && player.pos.X > 243 && player.pos.X < 716 {
		player.pos.Y = 422
	} else if player.pos.Y > 428 {
		player.pos.Y = 428
	}
}

// Room 2
func l1r2(player *Player) {

}

// Room 3
func l1r3(player *Player) {

}

// Level 2
