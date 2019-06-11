package main

import "github.com/faiface/pixel"

var (
	levels = []Level{}
)

func loadLevels() {
	levels = []Level{
		createLevel( // First level [outside first house]
			[]Room{
				createRoom(
					[]Object{
						createObject(pixel.V(100., 100.), images.box1, 2.),
						createObject(pixel.V(200., 300.), images.box1, 2.),
					},
					pixel.V(0., 0.),
				),
				createRoom(
					[]Object{
						createObject(pixel.V(200., 200.), images.box1, 2.),
					},
					pixel.V(100., 100.),
				),
			},
		),
	}
}
