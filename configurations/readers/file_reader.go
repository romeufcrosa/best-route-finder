package readers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// FileReader reads configurations from a given file
type FileReader struct {
	main string
}

// NewFileReader returns a new FileReader
func NewFileReader(file string) FileReader {
	return FileReader{file}
}

// ReadConfigurations reads a configuration from the passed file
func (reader FileReader) ReadConfigurations() ([]byte, error) {
	data, err := ioutil.ReadFile(reader.main)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	if err = json.Compact(&buffer, data); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
