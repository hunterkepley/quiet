package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

var (
	maxNodePosition = pixel.V(78, 59)
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
	index    pixel.Vec
}

func createNode(pos pixel.Vec, size pixel.Vec, passable bool, index pixel.Vec, parent Node) Node {
	return Node{0, 0, 0, pos, size, passable, &parent, index}
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

func createNodes(size pixel.Vec, nodes *[]Node, enemySize pixel.Vec) {
	for i := 0; i < int(winHeight/size.Y); i++ {
		for j := 0; j < int(winWidth/size.X); j++ {
			pos := pixel.V(float64(j)*size.X, float64(i)*size.Y)
			// Check if an object occupies the node
			for _, o := range currentLevel.rooms[currentLevel.currentRoomIndex].objects {
				if pos.X < o.pos.X+o.size.X+(enemySize.X/2.) &&
					pos.X+size.X > o.pos.X-(enemySize.X/2.) &&
					pos.Y < o.pos.Y+o.size.Y+(enemySize.Y/2.) &&
					pos.Y+size.Y > o.pos.Y-(enemySize.Y/2.) {
					if !o.backgroundObject {
						*nodes = append(*nodes, createNode(pos, size, false, pixel.V(float64(j), float64(i)), Node{}))
						break
					}
				}
			}
			*nodes = append(*nodes, createNode(pos, size, true, pixel.V(float64(j), float64(i)), Node{}))
		}
	}
}

func astar(start int, end int, nodes []Node, enemySize pixel.Vec) []Node { // start and end being the position

	startNode := createNode(nodes[start].pos, nodes[start].size, nodes[start].passable, nodes[start].index, Node{})
	endNode := createNode(nodes[end].pos, nodes[end].size, nodes[end].passable, nodes[end].index, Node{})

	totalIterations := 0

	// Initialize open and closed lists
	open := []Node{}
	closed := []Node{}

	// Add start node
	open = append(open, startNode)

	// Loop until end is found
	for len(open) > 0 {
		currentNode := open[0]
		currentIndex := 0
		for i, j := range open {
			for _, o := range currentLevel.rooms[currentLevel.currentRoomIndex].objects {
				if j.pos.X < o.pos.X+o.size.X+(enemySize.X/2.) &&
					j.pos.X+j.size.X > o.pos.X-(enemySize.X/2.) &&
					j.pos.Y < o.pos.Y+o.size.Y+(enemySize.Y/2.) &&
					j.pos.Y+j.size.Y > o.pos.Y-(enemySize.Y/2.) {
					if !o.backgroundObject {
						j.f += 10000
					}
				}
			}
			if j.f < currentNode.f {
				currentNode = j
				currentIndex = i
			}
			totalIterations++
			if totalIterations > 100000 {
				nodes[end].index = currentNode.index
			}
		}

		// Pop current off open list, add to closed list
		open = append(open[:currentIndex], open[currentIndex+1:]...)
		closed = append(closed, currentNode)

		// Found the goal
		if currentNode.index == nodes[end].index {
			fmt.Println("found the goal")
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
			nodePosition := pixel.V(currentNode.index.X+newPosition.X, currentNode.index.Y+newPosition.Y)

			// Make sure within range
			if nodePosition.X < 0 || nodePosition.X > (maxNodePosition.X-1) || nodePosition.Y < 0 || nodePosition.Y > (maxNodePosition.Y-1) {
				continue
			}

			costAddition := 0

			// Make sure walkable terrain
			if !nodes[int(nodePosition.X+(nodePosition.Y*maxNodePosition.X))].passable {
				continue
			}

			// Create new node
			newNode := createNode(pixel.V(nodePosition.X*10, nodePosition.Y*10), pixel.V(10., 10.), nodes[int(nodePosition.X+(nodePosition.Y*maxNodePosition.X))].passable, nodePosition, currentNode)

			// Append
			children = append(children, newNode)

			// Loop through children
			for _, child := range children {
				// Child is on the closed list
				for _, closedChild := range closed {
					if child.index == closedChild.index {
						continue
					}
				}

				if newPosition.X != 0 && nodePosition.Y != 0 {
					costAddition++
				}

				// Create the f, g, and h values
				child.g = currentNode.g + 1
				child.h = int(math.Pow(child.index.X-endNode.index.X, 2) + math.Pow(child.index.Y-endNode.index.Y, 2))
				child.f = child.g + child.h + costAddition

				// Child is already in the open list
				for _, openNode := range open {
					if child.index == openNode.index && child.g > openNode.g {
						continue
					}
				}

				// Add the child to the open list
				open = append(open, child)
			}
		}
	}
	// Didn't work
	fmt.Println("didn't work")
	return []Node{}
}

//this function will replace the current astar function whenever it is completed and working
func newAstar(start int, end int, nodes []Node, enemySize pixel.Vec) []Node {
	startNode := createNode(nodes[start].pos, nodes[start].size, nodes[start].passable, nodes[start].index, Node{})
	endNode := createNode(nodes[end].pos, nodes[end].size, nodes[end].passable, nodes[end].index, Node{})

	totalIterations := 0

	// Initialize open and closed lists
	open := []Node{}
	closed := []Node{}

	// Add start node
	open = append(open, startNode)

	//failure
	return []Node{}
}
