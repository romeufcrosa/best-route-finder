package entities

import (
	"encoding/json"
)

// Route holds the representation of a route
// includes all nodes visited, cost and duration
type Route struct {
	Voyage   []*Node `json:"voyage"`
	Cost     int     `json:"cost"`
	Duration int     `json:"duration"`
}

// ToJSON returns the JSON representation of a Route
func (r Route) ToJSON() (json.RawMessage, error) {
	return json.Marshal(r)
}
