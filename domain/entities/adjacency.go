package entities

// CompoundCost holds the two weights of an edge
type CompoundCost struct {
	Cost     int
	Duration int
}

// PathMatrix holds the matrix for all complete paths travelled from source to destination
type PathMatrix map[int][]*Node

// Push adds a full travelled path to the matrix
func (p PathMatrix) Push(key int, nodes []*Node) {
	for _, node := range nodes {
		p[key] = append(p[key], node)
	}
}

// CostMatrix holds all costs for the complete paths travelled
// keys of both matrixes are in sync
type CostMatrix map[int]CompoundCost

// Push adds the cost of a full travelled path to matrix
// TODO: this is currently having to traverse the path matrix again to
// ensure keys are in sync, this should be changed to use a single matrix
// to save on cycles
func (c CostMatrix) Push(key int, pMatrix PathMatrix, edges Edges) {
	var (
		cost, duration int
	)

	route := pMatrix[key]

	for i := 0; i < len(route); i++ {
		source := route[i]
		for _, next := range edges[source.ID] {
			if next.To.ID == route[i+1].ID {
				cost += next.Cost
				duration += next.Duration
			}
		}
	}

	c[key] = CompoundCost{
		Cost:     cost,
		Duration: duration,
	}
}
