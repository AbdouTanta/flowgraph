package db

import "go.mongodb.org/mongo-driver/v2/bson"

type Flow struct {
	Id          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // MongoDB uses _id as the primary key
	Description string        `json:"description" bson:"description"`
	Name        string        `json:"name" bson:"name"`
	Nodes       []Node        `json:"nodes" bson:"nodes"`
	Edges       []Edge        `json:"edges" bson:"edges"`
}

type Node struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Position struct {
		X float64 `json:"x" bson:"x"`
		Y float64 `json:"y" bson:"y"`
	} `json:"position" bson:"position"`
}

type Edge struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Source string `json:"source" bson:"source"` // ID of the source node
	Target string `json:"target" bson:"target"` // ID of the target node
}
