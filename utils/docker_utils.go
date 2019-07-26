package utils

import (
	docker "github.com/docker/docker/client"
)

func NewDockerClient() *docker.Client {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return cli
}
