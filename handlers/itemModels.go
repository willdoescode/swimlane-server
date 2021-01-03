package handlers

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Collection string `json:"collection,omitempty"`
	Title      string `json:"title,omitempty"`
	Details    string `json:"details,omitempty"`
	Assignee   string `json:"assignee,omitempty"`
}

// primitive.ObjectID is the type for
// the auto generated mongo ids

type ReturnItem struct {
	Id       primitive.ObjectID `json:"id,_id"`
	Title    string             `json:"title,omitempty"`
	Details  string             `json:"details,omitempty"`
	Assignee string             `json:"assignee,omitempty"`
}

type ManyItems struct {
	Collection string       `json:"collection,omitempty"`
	Items      []ReturnItem `json:"items,omitempty"`
}

type MoveItem struct {
	PrevCollection string             `json:"prevcollection,omitempty"`
	NewCollections string             `json:"newcollection,omitempty"`
	Id             primitive.ObjectID `json:"id,omitempty"`
}
