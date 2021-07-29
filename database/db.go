package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateConnection() (*mongo.Database, context.Context, context.CancelFunc) {
	atlasConnString := "mongodb+srv://Shaiful:123545@mycluster.gv2ka.mongodb.net/Store?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(atlasConnString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("Store"), ctx, cancel
}
