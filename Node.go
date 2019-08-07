package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

//Node ... A* nodes
type Node struct {
	f    int
	g    int
	h    int
	pos  pixel.Vec
	size pixel.Vec
}

func createNode(pos pixel.Vec, size pixel.Vec) Node {
	return Node{0, 0, 0, pos, size}
}

func (n *Node) render(imd *imdraw.IMDraw, col color.RGBA) {
	imd.Color = col
	imd.Push(n.pos, pixel.V(n.pos.X+n.size.X, n.pos.Y+n.size.Y))
	imd.Rectangle(1.)
}

func createNodes(size pixel.Vec, openNodes *[]Node, closedNodes *[]Node) {
	for i := 0; i < int(winHeight/size.Y); i++ {
		for j := 0; j < int(winWidth/size.X); j++ {
			pos := pixel.V(float64(j)*size.X, float64(i)*size.Y)
			// Check if an object occupies the node
			for _, o := range currentLevel.rooms[currentLevel.currentRoomIndex].objects {
				if pos.X < o.pos.X+o.size.X &&
					pos.X+size.X > o.pos.X &&
					pos.Y < o.pos.Y+o.size.Y &&
					pos.Y+size.Y > o.pos.Y {
					if !o.backgroundObject {
						*closedNodes = append(*closedNodes, createNode(pos, size))
						break
					}
				}
			}
			*openNodes = append(*openNodes, createNode(pos, size))
		}
	}
}
