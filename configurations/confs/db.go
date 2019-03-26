package confs

import (
	"fmt"
)

// DB database configurations
type DB struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	Address            string `json:"address"`
	Database           string `json:"database"`
	MaxConnections     int    `json:"max_conn"`
	MaxIdleConnections int    `json:"max_idle_conns"`
	ConnMaxLifetime    int    `json:"conn_max_lifetime_hours"`
	Timeouts           int    `json:"timeout_ms"`
}

// IsValid returns whether or not the configurations are valid
func (db DB) IsValid() bool {
	return db.Username != "" && db.Password != "" &&
		db.Address != "" && db.Database != ""
}

// DSN returns the DSN configured for the database
func (db DB) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		db.Username, db.Password,
		db.Address, db.Database,
	)
}

// OmittedDSN returns an omitted DSN configured for the database
func (db DB) OmittedDSN() string {
	return fmt.Sprintf(
		"%s:****@tcp(%s)/%s?parseTime=true",
		db.Username, db.Address, db.Database,
	)
}
