package fixtures

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
)

// LoadSQL loads a given SQL file and returns all the inserted queries
func LoadSQL(file string) []string {
	commentBlock := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)

	data, err := ioutil.ReadFile(file) // nolint
	checkError(err)

	result := bytes.Split(data, []byte(";"))
	result = result[:len(result)-1] // remove the empty one
	queries := make([]string, len(result))
	for index, value := range result {
		queries[index] = commentBlock.ReplaceAllString(strings.TrimSpace(string(value)+";"), "")
	}

	return queries
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
