package dockerlabel

import (
	"context"

	//"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
)

func newDockerClient() *docker.Client {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return cli
}

func GetLabelsFromContainer(containerId string) []string {
	result := []string{}
	cli := newDockerClient()
	containerDetails, err := cli.ContainerInspect(context.Background(),
		containerId)

	if err != nil {
		panic(err)
	}

	for key, value := range containerDetails.Config.Labels {
		if value != "false" {
			result = append(result, key)
		}
	}

	return result
}

func GetLabelsFromService(serviceId string) []string {
	return []string{}
}

func GetLabelsFromStack(stackId string) []string {
	return []string{}
}
