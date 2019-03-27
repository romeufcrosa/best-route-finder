package tests

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	// tests only
	_ "github.com/go-sql-driver/mysql"
	"github.com/romeufcrosa/best-route-finder/tests/fixtures"
)

var (
	pool       *sql.DB
	dbOnce     sync.Once
	clearEdges = `TRUNCATE edges`
	clearNodes = `TRUNCATE nodes`
)

// GetPool returns a connection for MySQL
func GetPool() *sql.DB {
	dbOnce.Do(func() {
		var err error

		address := fmt.Sprintf("demo:demo@tcp(%s)/routes?parseTime=true", os.Getenv("MYSQL_ADDRESS"))
		pool, err = sql.Open("mysql", address)
		checkError(err)

		err = pool.Ping()
		checkError(err)
	})

	return pool
}

// ClearTables ...
func ClearTables(db *sql.DB) {
	tx, err := db.Begin()
	checkError(err)

	_, err = tx.Exec(clearEdges)
	checkError(err)

	_, err = tx.Exec(clearNodes)
	checkError(err)

	err = tx.Commit()
	checkError(err)
}

// PrepareTables prepares the tables for freyr
func PrepareTables(db *sql.DB) {
	Cleanup(db)
	path := "assets/sql"
	if os.Getenv("ENV") == "docker" {
		path = "/assets/sql"
	}
	ExecuteSQL(db, fmt.Sprintf("%s/02_content.sql", path))
}

// Cleanup does a cleanup of the database
func Cleanup(db *sql.DB) {
	path := "assets/sql"
	if os.Getenv("ENV") == "docker" {
		path = "/assets/sql"
	}
	ExecuteSQL(db, fmt.Sprintf("%s/01_cleanup.sql", path))
	ExecuteSQL(db, fmt.Sprintf("%s/00_tables.sql", path))
}

// ExecuteSQL loads the SQL from the given file and executes them
func ExecuteSQL(db *sql.DB, file string) {
	queries := fixtures.LoadSQL(file)

	for _, query := range queries {
		_, err := db.Exec(query)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
