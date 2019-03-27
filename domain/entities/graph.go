package entities

import "sync"

// Graph holds the domain representation of the graph
type Graph struct {
	nodes []*Node
	edges Edges
	lock  sync.RWMutex
}

// SetNodes writes the nodes to the graph
func (g *Graph) SetNodes(nodes []*Node) {
	g.lock.Lock()
	g.nodes = nodes
	g.lock.Unlock()
}

// SetEdges writes the edges to the graph
func (g *Graph) SetEdges(edges map[int][]Edge) {
	g.lock.Lock()
	g.edges = edges
	g.lock.Unlock()
}

// Traverse does a BFS based traversal of the graph from source to destination
// TODO: this currently uses two adjacent matrixes (one for cost, another for travelled paths)
// in a future iteration these could most likely be merged into a single one
func (g *Graph) Traverse(source, destination *Node) (PathMatrix, CostMatrix) {
	traversal := g.edges

	queue := [][]*Node{}
	path := []*Node{}

	currentPath := 0

	pMatrix := make(PathMatrix)
	cMatrix := make(CostMatrix)

	path = append(path, source)
	queue = append(queue, path)

	for {
		if len(queue) == 0 {
			break
		}

		path = queue[0]
		queue = queue[1:len(queue)]
		last := path[len(path)-1]

		if last.ID == destination.ID {
			if len(path) > 2 { // skip direct paths
				pMatrix.Push(currentPath, path)
				cMatrix.Push(currentPath, pMatrix, traversal)
			}
			currentPath++
		}

		for i := 0; i < len(traversal[last.ID]); i++ {
			next := traversal[last.ID][i]
			if isNotVisited(next.To.ID, path) {
				newpath := path
				newpath = append(newpath, &next.To)
				queue = append(queue, newpath)
			}
		}

	}

	return pMatrix, cMatrix
}

func isNotVisited(i int, path []*Node) bool {
	n := len(path)

	for j := 0; j < n; j++ {
		if path[j].ID == i {
			return false
		}
	}

	return true

}
