package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	c, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	container, err := c.ContainerCreate(context.Background(), &container.Config{
		Image: "dao:4.4",
		ExposedPorts: map[nat.Port]struct{}{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			"27017/tcp": {
				{
					HostIP:   "127.0.0.1",
					HostPort: "0",
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	err = c.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("mongodb container started")
	inspect, err := c.ContainerInspect(context.Background(), container.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("mongodb container listening on port %s\n", inspect.NetworkSettings.Ports["27017/tcp"][0].HostPort)
	err = c.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
		Force: true,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb container stopped")
}
