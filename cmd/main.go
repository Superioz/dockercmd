package main

import (
	"context"
	"github.com/docker/docker/client"
)

func main() {
	_ = context.Background()
	_, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}
