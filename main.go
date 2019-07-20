package dockerlabel

import (
	"context"
	docker "github.com/docker/docker/client"
	ce "github.com/polpetta/docker-label/error"
)

func newDockerClient() *docker.Client {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return cli
}

func GetLabelsFromContainer(containerId string) ([]string, error) {
	var result []string
	cli := newDockerClient()
	containerDetails, err := cli.ContainerInspect(context.Background(),
		containerId)

	if err != nil {
		return nil, ce.NewNoContainerError(containerId)
	}

	for key, value := range containerDetails.Config.Labels {
		if value != "false" {
			result = append(result, key)
		}
	}

	return result, nil
}

func GetLabelsFromService(serviceId string) ([]string, error) {
	return []string{}, nil
}

func GetLabelsFromStack(stackId string) ([]string, error) {
	return []string{}, nil
}
