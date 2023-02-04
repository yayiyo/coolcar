package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"coolcar/shared/mongotesting"
)

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunMongoInDocker(m))
}

func TestResolveAccountID(t *testing.T) {
	client, err := mongotesting.NewClient(context.Background())
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
