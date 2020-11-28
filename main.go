package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// create and run container
	resp, err := container_create_and_run(cli, ctx, "lupino/ardb-server")

	// wait for container to stop
	// container_wait(cli, ctx, resp.ID)

	fmt.Println(resp.ID)
}
