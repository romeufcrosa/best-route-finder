package tests

import (
	"net/http"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/romeufcrosa/best-route-finder/configurations/readers"
)

var (
	httpPool = &http.Client{
		Transport: http.DefaultTransport,
	}
	once sync.Once
)

// ConfsReader prepares the given configurations
func ConfsReader() readers.FileReader {
	PrepareEnv()

	_, filename, _, _ := runtime.Caller(0)
	file := path.Join(path.Dir(filename), "..")

	return readers.NewFileReader(file + "/configurations/tests.json")
}

// PrepareEnv prepares the needed env vars if needed
// nolint
func PrepareEnv() {
	os.Setenv("ENV", "tests")
}

// SetTransport function to set the default http pool.
// Only usable during tests
func SetTransport(transport http.RoundTripper) {
	if os.Getenv("ENV") != "tests" {
		return
	}

	httpPool = &http.Client{
		Transport: transport,
	}
}
