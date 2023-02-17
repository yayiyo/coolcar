package mongotesting

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	imageName     = "dao:4.4"
	containerPort = "27017/tcp"
)

var mongoURI string

func RunMongoInDocker(m *testing.M) int {
	c, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	container, err := c.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		ExposedPorts: map[nat.Port]struct{}{
			containerPort: {},
		},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			containerPort: {
				{
					HostIP:   "127.0.0.1",
					HostPort: "0",
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	defer func() {
		err = c.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
			Force: true,
		})

		if err != nil {
			panic(err)
		}
		fmt.Println("mongodb container stopped")
	}()

	err = c.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("mongodb container started")
	inspect, err := c.ContainerInspect(context.Background(), container.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("mongodb container listening on port %s\n", inspect.NetworkSettings.Ports[containerPort][0].HostPort)
	mongoURI = fmt.Sprintf("mongodb://%s:%s", inspect.NetworkSettings.Ports[containerPort][0].HostIP,
		inspect.NetworkSettings.Ports[containerPort][0].HostPort)
	return m.Run()
}

// NewClient creates a new MongoDB client with the given configuration of dao container.
func NewClient(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
}

// NewDefaultClient creates a new default dao client with localhost and default port.
func NewDefaultClient(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
}

// SetupIndexes sets up indexes with given database.
func SetupIndexes(ctx context.Context, db *mongo.Database) (err error) {
	_, err = db.Collection("account").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "open_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		return err
	}

	_, err = db.Collection("trip").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "trip.accountid", Value: 1},
			{Key: "trip.status", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{"trip.status": 1}),
	})
	if err != nil {
		return err
	}

	_, err = db.Collection("profile").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "account_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	return nil
}
