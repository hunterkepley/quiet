package main

import "github.com/faiface/pixel"

var (
	levels = []Level{}
)

func loadLevels() {
	levels = []Level{
		createLevel( // First level [outside gas dealer]
			[]Room{
				createRoom(
					[]Object{ // Objects in room
						createObject(pixel.V(0., 0.), objectImages.gasStreet, 1., true, false),
						createObject(pixel.V(0., 0.), objectImages.gasBody, 1., true, false),
						createObject(pixel.V(0., 0.), objectImages.gasRoof, 1., false, true),
						createObject(pixel.V(518., 352.), objectImages.gasLeftPole, 15., false, false),
						createObject(pixel.V(764., 352.), objectImages.gasRightPole, 15., false, false),
					},
					pixel.V(50., 50.), // Player starting position
					true,              // Has rain
				),
				createRoom(
					[]Object{
						createObject(pixel.V(200., 200.), objectImages.box1, 2., false, false),
					},
					pixel.V(100., 100.),
					false,
				),
			},
		),
	}
}
