package main

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func container_wait(
	cli *client.Client,
	ctx context.Context,
	id string,
) {
	statusCh, errCh := cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}
}

func container_create_and_run(
	cli *client.Client,
	ctx context.Context,
	image string,
) (
	container.ContainerCreateCreatedBody,
	error,
) {
	// fetch image
	_, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	// create container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
	}, nil, nil, nil, "")
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	// start container
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	// return result
	return resp, nil
}
