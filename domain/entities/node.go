package entities

import (
	"encoding/json"
	"fmt"
)

// Node represents the data structure of a node, a step in a route
type Node struct {
	ID   int    `json:"node_id"`
	Name string `json:"name"`
}

// NewNode creates a Node from it's JSON payload
func NewNode(bodyBytes []byte) (node Node, err error) {
	if err = json.Unmarshal(bodyBytes, &node); err != nil {
		return node, err
	}

	return node, nil
}

// ToJSON returns a JSON representation of Node
func (n Node) ToJSON() (json.RawMessage, error) {
	return json.Marshal(n)
}

// String satisfies the interface to output the name of the node
// when print methods are called
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Name)
}
