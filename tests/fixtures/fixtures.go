package fixtures

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// LoadJSON loads a JSON file name from a file into a json.RawMessage
func LoadJSON(fname string) (json.RawMessage, error) {
	raw, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(raw), nil
}

// LoadAndCompressJSON loads the content of a JSON file, compacts it
// and returns it as json.RawMessage
func LoadAndCompressJSON(fname string) (json.RawMessage, error) {
	raw, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, raw)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(buffer.String()), nil
}

// LoadJSONInto loads a JSON file name from a file into the supplied target
func LoadJSONInto(fname string, target interface{}) error {
	raw, err := LoadJSON(fname)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &target)
	if err != nil {
		return err
	}

	return nil
}
