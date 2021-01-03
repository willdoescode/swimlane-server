package handlers

import (
	"context"
	"encoding/json"
	u "main/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetManyItemsHandler(w http.ResponseWriter, r *http.Request) {
	var items []ReturnItem
	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := Client.Collection(mux.Vars(r)["id"]).Find(ctx, bson.D{})
	u.Error(err)
	defer cur.Close(context)
	for cur.Next(context) {
		var result bson.M
		err := cur.Decode(&result)
		u.Error(err)
		items = append(
			items,
			ReturnItem{
				result["_id"].(primitive.ObjectID),
				result["title"].(string),
				result["details"].(string),
				result["assignee"].(string),
			},
		)
	}
	res := ManyItems{mux.Vars(r)["id"], items}
	w.Header().Set("Content-type", "application/json")
	j, _ := json.Marshal(res)
	w.Write(j)
}
