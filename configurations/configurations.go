package configurations

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" // use driver for db pool
	"github.com/romeufcrosa/best-route-finder/configurations/confs"
	"github.com/romeufcrosa/best-route-finder/configurations/readers"
)

// vars basically constants
var (
	conf                    Config
	ErrLoadingConfiguration = errors.New("error loading configuration")
	pool                    *sql.DB
	confOnce, dbOnce        sync.Once
)

// Config struct that holds all configurations
type Config struct {
	Env          string   `json:"env"`
	IsConfigured bool     `json:"-"`
	ListenAddr   string   `json:"listen_addr"`
	DB           confs.DB `json:"db"`
}

// IsConfigured returns if the service is configured
func IsConfigured() bool {
	return conf.IsConfigured
}

// Updated returns whether or not the configurations were updated
func Updated() bool {
	return conf.IsUpdated()
}

// IsValid indicates whether or not the configurations are valid
func (c Config) IsValid() bool {
	return true
}

// IsUpdated for now just return false, since we don't have
// any live reloads yet...
func (c Config) IsUpdated() bool {
	return false
}

// Env loads the configured environment
func Env() string {
	if conf.Env == "" {
		return "tests"
	}

	return conf.Env
}

// ListenAddr returns the host:port the server listens to
func ListenAddr() string {
	if conf.ListenAddr == "" {
		conf.ListenAddr = ":8080"
	}

	return conf.ListenAddr
}

// Load loads the configurations using a file reader
func Load(ctx context.Context, reader readers.FileReader) (err error) {
	if IsConfigured() {
		return
	}

	confOnce.Do(func() {
		loader := Loader{reader}
		conf, err = loader.LoadConfigurations()
	})
	if err != nil {
		log.Fatalln("could not load configurations")
	}

	return
}

// Pool returns the connection pool to a given database
func Pool() (p *sql.DB, err error) {
	dbOnce.Do(func() {
		pool, err = sql.Open("mysql", conf.DB.DSN())
		if err != nil {
			return
		}

		log.Println("connecting to database at", conf.DB.OmittedDSN())
		if err = pool.Ping(); err != nil {
			return
		}

		pool.SetMaxOpenConns(conf.DB.MaxConnections)
		pool.SetMaxIdleConns(conf.DB.MaxIdleConnections)
		pool.SetConnMaxLifetime(time.Duration(conf.DB.ConnMaxLifetime) * time.Hour)
	})

	return pool, err
}
