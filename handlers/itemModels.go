package handlers

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Collection string             `json:"collection,omitempty"`
	Id         primitive.ObjectID `json:"id"`
	Title      string             `json:"title,omitempty"`
	Details    string             `json:"details,omitempty"`
	Assignee   string             `json:"assignee,omitempty"`
}

type ReturnItem struct {
	Id       primitive.ObjectID `json:"id"`
	Title    string             `json:"title,omitempty"`
	Details  string             `json:"details,omitempty"`
	Assignee string             `json:"assignee,omitempty"`
}

type ManyItems struct {
	Collection string       `json:"collection,omitempty"`
	Items      []ReturnItem `json:"items,omitempty"`
}