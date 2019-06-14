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
						createObject(pixel.V(100., 100.), images.box1, 2.),
						createObject(pixel.V(200., 300.), images.box1, 2.),
					},
					pixel.V(50., 50.), // Player starting position
					true,              // Has rain
				),
				createRoom(
					[]Object{
						createObject(pixel.V(200., 200.), images.box1, 2.),
					},
					pixel.V(100., 100.),
					false,
				),
			},
		),
	}
}
