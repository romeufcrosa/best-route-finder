package entities

import "encoding/json"

// Edges represents a matrix of edges linked by source node id
type Edges map[int][]Edge

// Edge represents the structure of a node's edge, it's arc to another node
type Edge struct {
	EdgeID   int `json:"edge_id"`
	From     Node
	To       Node
	Duration int `json:"duration"`
	Cost     int `json:"cost"`
}

// EdgeDTO holds a transformation representation of edge to be mapped on requests
type EdgeDTO struct {
	ID       int `json:"id"`
	FromID   int `json:"from_id"`
	ToID     int `json:"to_id"`
	Duration int `json:"duration"`
	Cost     int `json:"cost"`
}

// NewEdge creates a new edge from it's JSON payload
func NewEdge(bodyBytes []byte) (edge Edge, err error) {
	if err = json.Unmarshal(bodyBytes, &edge); err != nil {
		return edge, err
	}

	return edge, nil
}

// ToJSON returns a JSON representation of Edge
func (e Edge) ToJSON() (json.RawMessage, error) {
	return json.Marshal(e.ConvertTo())
}

// UnmarshalJSON of Edge returns the domain representation
func (e *Edge) UnmarshalJSON(data []byte) error {
	var temp EdgeDTO

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	e.ConvertFrom(temp)

	return nil
}

// ConvertFrom converts from the DTO to the domain representation
func (e *Edge) ConvertFrom(dto EdgeDTO) {
	e.From = Node{
		ID: dto.FromID,
	}
	e.To = Node{
		ID: dto.ToID,
	}
	e.Duration = dto.Duration
	e.Cost = dto.Cost
}

// ConvertTo converts from the domain representation to DTO
func (e *Edge) ConvertTo() (dto EdgeDTO) {
	return EdgeDTO{
		ID:       e.EdgeID,
		FromID:   e.From.ID,
		ToID:     e.To.ID,
		Duration: e.Duration,
		Cost:     e.Cost,
	}
}
