package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// example from: https://docs.docker.com/develop/sdk/examples/#list-and-manage-containers
// I don't know, what I want to add to this program
func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	// print out every container's id
	for _, container := range containers {
		fmt.Println(container.ID)
	}
}
