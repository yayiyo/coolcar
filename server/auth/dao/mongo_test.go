package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"coolcar/shared/mongotesting"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI string
)

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunMongoInDocker(m, &mongoURI))
}

func TestResolveAccountID(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	m := NewMongo(client.Database("coolcar"))
	account, err := m.ResolveAccountID(context.Background(), "123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account)
}
