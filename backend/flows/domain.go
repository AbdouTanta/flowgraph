package flows

import "go.mongodb.org/mongo-driver/v2/bson"

// For now, I'm using the domain model in the REST layer too.
// Not worth the complexity of separating them yet.

type Flow struct {
	Id          bson.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB uses _id as the primary key
	Description string        `bson:"description" json:"description"`
	Name        string        `bson:"name" json:"name"`
	Nodes       []Node        `bson:"nodes" json:"nodes"`
	Edges       []Edge        `bson:"edges" json:"edges"`
	UserId      string        `bson:"user_id" json:"user_id"`
}

type NodeData struct {
	Label string `bson:"label" json:"label"` // Label for the node
}

type Node struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Position struct {
		X float64 `bson:"x" json:"x"`
		Y float64 `bson:"y" json:"y"`
	} `bson:"position" json:"position"`
	Type string   `bson:"type" json:"type"` // Type of the node (e.g., "input", "output", "process")
	Data NodeData `bson:"data" json:"data"` // Data associated with the node
}

type Edge struct {
	ID     string `bson:"_id,omitempty" json:"id"`
	Source string `bson:"source" json:"source"` // ID of the source node
	Target string `bson:"target" json:"target"` // ID of the target node
}
