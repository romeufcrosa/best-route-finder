package mysql

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/repositories/mysql/internal"
)

var (
	insertEdge = `
		INSERT into edges(edge_id, from_id, to_id, duration, cost)
		VALUES(NULL, ?, ?, ?, ?)
	`

	insertNode = `
		INSERT into nodes(node_id, name)
		VALUES(NULL, ?)
	`

	getEdges = `SELECT edge_id, from_id, to_id, duration, cost FROM edges`

	getNodes = `SELECT node_id, name FROM nodes`

	getRoutes = `
		SELECT e.*, n1.Name AS FromName, n2.Name AS ToName FROM edges e
		JOIN nodes n1 ON e.from_id = n1.node_id
		JOIN nodes n2 ON e.to_id = n2.node_id
	`
)

// Routes type holds the mysql connection pool
type Routes struct {
	pool *sql.DB
}

// Route struct that identities a data storage represenation of a route
type Route struct {
	EdgeID   int
	FromID   int
	ToID     int
	Duration int
	Cost     int
	FromName string
	ToName   string
}

// NewRoutes factory method to create a new repository instance
func NewRoutes(pool *sql.DB) Routes {
	return Routes{pool}
}

// AddEdge inserts an edge entity into the database
func (r Routes) AddEdge(ctx context.Context, edge domain.Edge) (result domain.Edge, err error) {
	err = internal.InTransaction(ctx, r.pool, func(ctx context.Context, tx *sql.Tx) error {
		var res sql.Result
		if err != nil {
			return errors.Wrap(err, "could not generate ID")
		}

		if res, err = tx.ExecContext(
			ctx, insertEdge, edge.From.ID, edge.To.ID, edge.Duration, edge.Cost,
		); err != nil {
			return err
		}
		insertedID, err := res.LastInsertId()
		if err != nil {
			return err
		}

		result = edge
		result.EdgeID = int(insertedID)

		return nil
	})

	if err != nil {
		return
	}

	return
}

// AddNode inserts a node entity into the database
func (r Routes) AddNode(ctx context.Context, node domain.Node) (result domain.Node, err error) {
	err = internal.InTransaction(ctx, r.pool, func(ctx context.Context, tx *sql.Tx) error {
		var res sql.Result
		if err != nil {
			return errors.Wrap(err, "could not generate ID")
		}

		if res, err = tx.ExecContext(
			ctx, insertNode, node.Name,
		); err != nil {
			return err
		}
		insertedID, err := res.LastInsertId()
		if err != nil {
			return err
		}

		result = node
		result.ID = int(insertedID)

		return nil
	})

	if err != nil {
		return
	}

	return
}

// GetRoute returns the neighbours and nodes to calculate the optimal route
// TODO: should only return the Route, the rest should be calculated in the interactor
func (r Routes) GetRoute(ctx context.Context, origin, destination int) (near map[int][]domain.Edge, nodes []*domain.Node, err error) {
	var (
		rows         *sql.Rows
		closeErr     error
		visited      bool
		vOriginNodes = make(map[int]bool)
	)
	near = make(map[int][]domain.Edge)

	rows, err = r.pool.QueryContext(ctx, getRoutes)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not perform query")
	}
	defer func() {
		if closeErr = rows.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		var route Route
		if err = rows.Scan(
			&route.EdgeID, &route.FromID,
			&route.ToID, &route.Duration,
			&route.Cost, &route.FromName,
			&route.ToName,
		); err != nil {
			return nil, nil, errors.Wrap(err, "could not scan rule")
		}

		originNode := &domain.Node{
			ID:   route.FromID,
			Name: route.FromName,
		}
		if visited = vOriginNodes[originNode.ID]; !visited {
			vOriginNodes[originNode.ID] = true
			nodes = append(nodes, originNode)
		}

		destinationNode := &domain.Node{
			ID:   route.ToID,
			Name: route.ToName,
		}

		edge := domain.Edge{
			EdgeID:   route.EdgeID,
			From:     *originNode,
			To:       *destinationNode,
			Cost:     route.Cost,
			Duration: route.Duration,
		}
		near[originNode.ID] = append(near[originNode.ID], edge)
	}

	return
}

// GetNodes returns all the nodes
func (r Routes) GetNodes(ctx context.Context) ([]domain.Node, error) {
	var (
		closeErr, err error
		nodes         []domain.Node
		rows          *sql.Rows
	)
	rows, err = r.pool.QueryContext(ctx, getNodes)
	if err != nil {
		return nil, errors.Wrap(err, "could not perform query")
	}
	defer func() {
		if closeErr = rows.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		var node domain.Node
		if err = rows.Scan(
			&node.ID, &node.Name,
		); err != nil {
			return nil, errors.Wrap(err, "could not scan rule")
		}

		nodes = append(nodes, node)
	}

	return nil, nil
}

// GetEdges returns all the edges
func (r Routes) GetEdges(ctx context.Context) ([]domain.EdgeDTO, error) {
	var (
		closeErr, err error
		edges         []domain.EdgeDTO
		rows          *sql.Rows
	)
	rows, err = r.pool.QueryContext(ctx, getEdges)
	if err != nil {
		return nil, errors.Wrap(err, "could not perform query")
	}
	defer func() {
		if closeErr = rows.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		var edge domain.EdgeDTO
		if err = rows.Scan(
			&edge.ID, &edge.FromID,
			&edge.ToID, &edge.Duration,
			&edge.Cost,
		); err != nil {
			return nil, errors.Wrap(err, "could not scan rule")
		}

		edges = append(edges, edge)
	}

	return nil, nil
}
