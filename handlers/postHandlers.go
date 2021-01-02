package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
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
	fmt.Printf("%v: [Collection: %s, Title: %s, Details: %s, Assignee: %s]\n", res, item.Collection, item.Title, item.Details, item.Assignee)
	w.Write(j)
}
