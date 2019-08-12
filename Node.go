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
	parent   *Node
}

func createNode(pos pixel.Vec, size pixel.Vec, passable bool) Node {
	return Node{0, 0, 0, pos, size, passable, nil}
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

func astar(start int, end int, open []Node, closed []Node) []Node { // start and end being the position

	// Using
	// https://medium.com/@nicholas.w.swift/easy-a-star-pathfinding-7e6689c7f7b2
	// To make this

	startNode := createNode(open[start].pos, open[start].size, open[start].passable)
	endNode := createNode(open[end].pos, open[end].size, open[end].passable)

	for len(open) > 0 {
		currentNode := open[0]
		currentIndex := 0
		for i, j := range open {
			if j.f < currentNode.f {
				currentNode = j
				currentIndex = i
			}
		}

		// Pop current off open list, add to closed list
		open = append(open[:currentIndex], open[currentIndex+1:]...)
		closed = append(closed, currentNode)

		// Found the goal
		if currentIndex == end {
			path := []Node{}
			current := currentNode
			emptyNode := Node{}
			for current != emptyNode {
				path = append(path, current)
				current = *current.parent
			}
			// Reverse slice
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path
		}

		// Generate children
		children := []Node{}
		for _, newPosition := range []pixel.Vec{pixel.V(0, -1), pixel.V(0, 1), pixel.V(-1, 0), pixel.V(1, 0), pixel.V(-1, -1), pixel.V(-1, 1), pixel.V(1, -1), pixel.V(1, 1)} {
			// Get node position
			//nodePosition := pixel.V(current)
			// Make the nodes listed like so:
			// [....]
			// [....]
			// [....]
		}
	}
}
