package mysql

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/tests"
)

var (
	ctx  = context.TODO()
	pool = tests.GetPool()
)

func TestEdgeStore(t *testing.T) {
	RegisterTestingT(t)

	repository := NewRoutes(pool)

	edge := domain.Edge{
		From:     domain.Node{},
		To:       domain.Node{},
		Cost:     123,
		Duration: 321,
	}

	inserted, err := repository.AddEdge(ctx, edge)
	Expect(err).To(BeNil(), "Should return no error")

	read := tests.Edge(pool, inserted.EdgeID)

	Expect(read.ID).To(Equal(inserted.EdgeID), "should match inserted edge")
}

func TestNodeStore(t *testing.T) {
	RegisterTestingT(t)

	repository := NewRoutes(pool)

	node := domain.Node{
		Name: "Node_A",
	}

	inserted, err := repository.AddNode(ctx, node)
	Expect(err).To(BeNil(), "Should return no error")

	read := tests.Node(pool, inserted.ID)

	Expect(read.ID).To(Equal(inserted.ID), "should match inserted node")
}

func TestGetRoutes(t *testing.T) {
	RegisterTestingT(t)

	repository := NewRoutes(pool)

	NodeA := domain.Node{
		Name: "Node_A",
	}
	insNodeA, err := repository.AddNode(ctx, NodeA)
	Expect(err).To(BeNil(), "Should return no error")

	NodeB := domain.Node{
		Name: "Node_B",
	}
	insNodeB, err := repository.AddNode(ctx, NodeB)
	Expect(err).To(BeNil(), "Should return no error")

	NodeC := domain.Node{
		Name: "Node_C",
	}
	insNodeC, err := repository.AddNode(ctx, NodeC)
	Expect(err).To(BeNil(), "Should return no error")

	EdgeAtoB := domain.Edge{
		From:     insNodeA,
		To:       insNodeB,
		Cost:     1,
		Duration: 1,
	}
	insEdgeAtoB, err := repository.AddEdge(ctx, EdgeAtoB)
	Expect(err).To(BeNil(), "Should return no error")

	EdgeBtoC := domain.Edge{
		From:     insNodeB,
		To:       insNodeC,
		Cost:     2,
		Duration: 2,
	}
	insEdgeBtoC, err := repository.AddEdge(ctx, EdgeBtoC)
	Expect(err).To(BeNil(), "Should return no error")

	near, nodes, err := repository.GetRoute(ctx, NodeA.ID, NodeC.ID)
	Expect(err).To(BeNil(), "Should return no error")

	expectedNear := make(map[int][]domain.Edge)
	expectedNear[insNodeA.ID] = append(expectedNear[insNodeA.ID], insEdgeAtoB)
	expectedNear[insNodeB.ID] = append(expectedNear[insNodeB.ID], insEdgeBtoC)

	Expect(near).To(Equal(expectedNear), "should match returned edges")

	expectedNodes := []*domain.Node{}
	expectedNodes = append(expectedNodes, &insNodeA)
	expectedNodes = append(expectedNodes, &insNodeB)

	Expect(nodes).To(Equal(expectedNodes), "should match returned nodes")
}

func init() {
	tests.ClearTables(pool)
}
