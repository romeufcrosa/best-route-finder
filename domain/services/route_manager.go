package services

import (
	"context"

	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
)

// RouteStorage provides an interface to the repository
type RouteStorage interface {
	AddEdge(context.Context, domain.Edge) (domain.Edge, error)
	AddNode(context.Context, domain.Node) (domain.Node, error)
	GetRoute(context.Context, int, int) (map[int][]domain.Edge, []*domain.Node, error)
}

// Routes holds the representation of this interactor
type Routes struct {
	storage RouteStorage
}

// NewRoutesInteractor initializes a new interactor
func NewRoutesInteractor(storage RouteStorage) Routes {
	return Routes{storage}

}

// AddNode inserts a node into the storage
func (r Routes) AddNode(ctx context.Context, node domain.Node) (domain.Node, error) {
	result, err := r.storage.AddNode(ctx, node)
	if err != nil {
		return node, err
	}

	return result, nil
}

// AddEdge inserts an edge into the storage
func (r Routes) AddEdge(ctx context.Context, edge domain.Edge) (domain.Edge, error) {
	result, err := r.storage.AddEdge(ctx, edge)
	if err != nil {
		return edge, err
	}

	return result, nil
}

// GetRoute returns the optimal route for origin and destination
func (r Routes) GetRoute(ctx context.Context, origin, destination int) (domain.Route, error) {
	graph := domain.Graph{}
	var sourceNode, destNode *domain.Node
	edges, nodes, err := r.storage.GetRoute(ctx, origin, destination)
	if err != nil {
		return domain.Route{}, err
	}

	graph.SetEdges(edges)
	graph.SetNodes(nodes)

	for _, node := range nodes {
		if node.ID == origin {
			sourceNode = node
		}
	}
	destNode = &domain.Node{
		ID: destination,
	}

	pathMatrix, costMatrix := graph.Traverse(sourceNode, destNode)

	return findOptimalRoute(pathMatrix, costMatrix), nil
}

func findOptimalRoute(pMatrix domain.PathMatrix, cMatrix domain.CostMatrix) domain.Route {
	previousCost := cMatrix[0].Cost
	previousDuration := cMatrix[0].Duration
	chosenPath := 0

	for i := 1; i < len(pMatrix); i++ {
		if cMatrix[i].Cost <= previousCost &&
			cMatrix[i].Duration <= previousDuration {
			chosenPath = i
			previousCost = cMatrix[i].Cost
			previousDuration = cMatrix[i].Duration
		}
	}

	return domain.Route{
		Voyage:   pMatrix[chosenPath],
		Cost:     previousCost,
		Duration: previousDuration,
	}
}
