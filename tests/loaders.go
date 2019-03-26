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
	ExecuteSQL(pool, "../../../assets/sql/03_route_scenario.sql")
}
