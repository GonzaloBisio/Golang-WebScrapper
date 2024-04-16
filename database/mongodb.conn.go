package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	usr      = "gonzalobisio5"
	pwd      = "9HtF2JnfgolEQmaG"
	host     = "cluster0.brtfhbg.mongodb.net"
	database = "user"
)

func GetCollection(collectionName string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", usr, pwd, host, database)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	return client.Database(database).Collection(collectionName)
}
