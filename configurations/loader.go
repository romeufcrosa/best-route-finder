package configurations

import (
	"encoding/json"
	"errors"
)

var (
	// ErrInvalidConfigurations error sent when the configuration is invalid
	ErrInvalidConfigurations = errors.New("configurations are invalid")
)

// Reader reads a set of configurations
type Reader interface {
	ReadConfigurations() ([]byte, error)
}

// Loader loads a set of configurations and stores them locally
type Loader struct {
	reader Reader
}

// LoadConfigurations loads all configurations for the service to work
func (loader Loader) LoadConfigurations() (configurations Config, err error) {
	var data []byte

	data, err = loader.reader.ReadConfigurations()
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &configurations); err != nil {
		return
	}

	if !configurations.IsValid() {
		return configurations, ErrInvalidConfigurations
	}

	configurations.IsConfigured = true

	return
}
