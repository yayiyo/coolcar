package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	c := context.Background()
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
