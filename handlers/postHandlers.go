package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	item := new(Item)
	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal(`{"Bad json"}`)
		w.Write(j)
		return
	}
	strings.Trim(item.Collection, " \n")

	res, err := Client.Collection(item.Collection).InsertOne(ctx, item)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	j, _ := json.Marshal(res)
	log.Printf(
		"%v: [Collection: %s, Title: %s, Details: %s, Assignee: %s]\n", res, item.Collection, item.Title, item.Details, item.Assignee,
	)
	w.Write(j)
}

func MoveItemHandler(w http.ResponseWriter, r *http.Request) {
	mItem := new(MoveItem)
	var newItem struct {
		Collection string `json:"collection"`
		Title      string `json:"title,omitempty"`
		Details    string `json:"details,omitempty"`
		Assignee   string `json:"assignee"`
	}
	err := json.NewDecoder(r.Body).Decode(mItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal(`{"Bad json"}`)
		w.Write(j)
		return
	}
	Client.Collection(mItem.PrevCollection).FindOneAndDelete(
		ctx,
		bson.D{primitive.E{Key: "_id", Value: mItem.Id}},
	).Decode(&newItem)
	res, err := Client.Collection(mItem.NewCollections).InsertOne(ctx, newItem)
	j, _ := json.Marshal(res)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Printf(
		"%v: [Collection: %s, Title: %s, Details: %s, Assignee: %s] -> %s from %s\n", res,
		newItem.Collection,
		newItem.Title,
		newItem.Details,
		newItem.Assignee,
		mItem.NewCollections,
		mItem.PrevCollection,
	)
	w.Write(j)
}
