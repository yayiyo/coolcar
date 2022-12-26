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
)

const (
	imageName     = "mongo:4.4"
	containerPort = "27017/tcp"
)

func RunMongoInDocker(m *testing.M, mongoURI *string) int {
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
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", inspect.NetworkSettings.Ports[containerPort][0].HostIP,
		inspect.NetworkSettings.Ports[containerPort][0].HostPort)
	return m.Run()
}
