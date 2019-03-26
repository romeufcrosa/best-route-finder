package tests

import (
	"database/sql"

	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
)

// Edge ...
func Edge(pool *sql.DB, edgeID int) (result domain.EdgeDTO) {
	stmt := `SELECT edge_id, from_id, to_id, duration, cost
                FROM edges
                WHERE edge_id = ?`

	err := pool.QueryRow(stmt, edgeID).Scan(
		&result.ID, &result.FromID,
		&result.ToID, &result.Duration,
		&result.Cost,
	)
	checkError(err)

	return
}

// Node ...
func Node(pool *sql.DB, nodeID int) (result domain.Node) {
	stmt := `SELECT node_id, name
                FROM nodes
                WHERE node_id = ?`

	err := pool.QueryRow(stmt, nodeID).Scan(
		&result.ID, &result.Name,
	)
	checkError(err)

	return
}

// SetupRoutesScenario ...
func SetupRoutesScenario(pool *sql.DB) {
	ClearTables(pool)
	stmt := `
	INSERT INTO nodes(node_id, name) VALUES
		(NULL, 'Node_A'),
		(NULL, 'Node_B'),
		(NULL, 'Node_C'),
		(NULL, 'Node_D')
	`
	_, err := pool.Exec(stmt)
	checkError(err)

	stmt = `
	INSERT INTO edges(edge_id, from_id, to_id, duration, cost) VALUES
		(NULL, 1, 3, 1, 20),
		(NULL, 1, 4, 1, 14),
		(NULL, 3, 2, 1, 12),
		(NULL, 4, 2, 1, 8)
	`
	_, err = pool.Exec(stmt)
	checkError(err)
}
