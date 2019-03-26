// The following directive is necessary to make the package coherent:

// +build ignore

// This program prepares the database by flushing it
// and inserting data into the tables of MySQL
package main

import "github.com/romeufcrosa/best-route-finder/tests"

func main() {
	pool := tests.GetPool()

	tests.PrepareTables(pool)
}
