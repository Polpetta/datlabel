package datlabel

import (
	"context"
	ce "github.com/polpetta/datlabel/error"
	"github.com/polpetta/datlabel/utils"
)

// Given a container id, the function returns the current labels only, without
// any field description.
func GetLabelsFromContainer(containerId string) ([]string, error) {
	cli := utils.NewDockerClient()
	containerDetails, err := cli.ContainerInspect(context.Background(),
		containerId)

	if err != nil {
		return nil, ce.NewNoSuchElement(containerId)
	}

	return utils.FilterLabelsByString(
		containerDetails.Config.Labels,
		"false"), nil
}

// Given a service id, the function returns the service labels without any filed
// description
func GetLabelsFromService(serviceId string) ([]string, error) {
	cli := utils.NewDockerClient()
	serviceDetails, _, err := cli.ServiceInspectWithRaw(context.Background(),
		serviceId)

	if err != nil {
		return nil, ce.NewNoSuchElement(serviceId)
	}

	return utils.FilterLabelsByString(
		serviceDetails.Spec.Labels,
		"false"), nil
}

// The idea here is to return all the labels a stack has, in order to collect
// them in a list
func GetLabelsFromStack(stackId string) ([]string, error) {
	return []string{}, nil
}
