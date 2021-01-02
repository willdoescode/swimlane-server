package handlers

import (
	"context"
	u "main/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx    = context.TODO()
	C, e   = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	Client = C.Database("swimlanes")
)

func init() {
	u.Error(e)
}
