package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

//Node ... A* nodes
type Node struct {
	f        int
	g        int
	h        int
	pos      pixel.Vec
	size     pixel.Vec
	passable bool
}

func createNode(pos pixel.Vec, size pixel.Vec, passable bool) Node {
	return Node{0, 0, 0, pos, size, passable}
}

func (n *Node) render(imd *imdraw.IMDraw) {
	imd.Push(n.pos, pixel.V(n.pos.X+n.size.X, n.pos.Y+n.size.Y))
	if !n.passable {
		imd.Color = colornames.Red
		imd.Rectangle(1.)
	} else {
		imd.Color = colornames.Green
		imd.Rectangle(1.)
	}
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
						*openNodes = append(*openNodes, createNode(pos, size, false))
						break
					}
				}
			}
			*openNodes = append(*openNodes, createNode(pos, size, true))
		}
	}
}


func astar(start int, end int) { // start and end being the position
	
}