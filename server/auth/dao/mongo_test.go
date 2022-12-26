package dao

import (
	"context"
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

func TestMain(m *testing.M) {
	c := context.Background()
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("coolcar")
	m.Run()
}

func TestResolveAccountID(t *testing.T) {
	m := NewMongo(db)
	account, err := m.ResolveAccountID(context.Background(), "123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account)
}
