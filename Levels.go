package main

import "github.com/faiface/pixel"

var (
	levels = []Level{}
)

func loadLevels() {
	levels = []Level{
		createLevel( // First level [outside gas dealer]
			[]Room{
				createRoom( // Outside 1
					[]Object{ // Objects in room
						createObject(pixel.V(0., 0.), objectImages.gasStreet, 1., true, false, false, false, 0.),
						createObject(pixel.V(0., 0.), objectImages.gasFence, 1., true, false, false, true, 100.),
						createObject(pixel.V(0., 0.), objectImages.gasBody, 1., true, false, false, true, 100.),
						createObject(pixel.V(0., 0.), objectImages.gasRoof, 1., false, true, false, false, 0.),
						createObject(pixel.V(0., 0.), objectImages.gasLight, 1., false, true, false, false, 0.),
						createObject(pixel.V(518., 352.), objectImages.gasLeftPole, 15., false, false, true, true, 100.),
						createAnimatedObject(pixel.V(654., 411.), objectSpritesheets.trashCanSheet, 0.03, 1., false, true, true, true, 0.),
						createObject(pixel.V(764., 352.), objectImages.gasRightPole, 15., false, false, true, true, 100.),
					},
					[]Enemy{ // Enemies in the room

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
						createObject(pixel.V(28., 639.), objectImages.backWall1, 1., false, false, true, true, 100.),
						createObject(pixel.V(28., 0.), objectImages.bottomWall1, 2., false, false, true, true, 100.),
						createObject(pixel.V(0., 0.), objectImages.wall1, 1., false, false, true, true, 0.),
						createObject(pixel.V(996., 0.), objectImages.wall1, 1., false, false, true, true, 0.),
						createObject(pixel.V(0., 0.), objectImages.floor1, 1., true, false, false, false, 0.),
						createObject(pixel.V(100., 180.), objectImages.box1, 2., false, false, true, true, 15.),
						createObject(pixel.V(200., 80.), objectImages.box1, 2., false, false, true, true, 15.),
					},
					[]Enemy{
						createEnemy(pixel.V(300., 50.), enemyImages.larvaImages.stillLeft, 1., 20., 5., 0.15, 0.5, 0.05, 5., 80.),
					},
					pixel.V(100., 100.),
					false,
					[]pixel.Rect{},
					grayscaleShader,
					l1r2,
					true,
				),
			},
		),
	}
}

// Functions for rooms

// Level 1
func l1r1(player *Player) {
	// Collision against gas station
	if player.pos.Y > 422 && player.pos.X > 243 && player.pos.X < 716 {
		player.pos.Y = 422
	} else if player.pos.Y > 428 {
		player.pos.Y = 428
	}
}

func l1r2(player *Player) {

}

// Level 2
