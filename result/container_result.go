package result

import (
	"github.com/docker/docker/api/types"
)

// Struct that represents a container. It contains a pointer for the
// docker-defined structure. This struct act like a proxy
type Container struct {
	rawContainerDefinition *types.Container
}

// Getter method to return the original docker container structure
func (c *Container) RawContainerDefinition() *types.Container {
	return c.rawContainerDefinition
}

// Getter method to return a list of labels
func (c *Container) Labels() []Label {
	var labels []Label
	for key, value := range c.rawContainerDefinition.Labels {
		labels = append(labels, Label{
			name:  key,
			value: value,
		})
	}
	return labels
}

// Getter method to return the container id
func (c *Container) Id() string {
	return c.rawContainerDefinition.ID
}

// Represents the result for a container search.
// It allows to get the list of containers found and to filter them
type ContainerResult interface {
	Result
	Containers() []Container
	Filter(
		filter func(container *Container) *Container) (ContainerResult, error)
}

// Real ContainerResult interface implementation
type containerResultImpl struct {
	ContainerResult
	containers []Container
}

// Getter method to obtain the list of containers
func (c *containerResultImpl) Containers() []Container {
	return c.containers
}

// Performs filtering operation on all the containers.
// A new ContainerResult is returned at the end of the operation,
// enabling the possibility to perform additional filtering.
func (c *containerResultImpl) Filter(
	filter func(container *Container) *Container) (ContainerResult, error) {
	var result []Container
	for _, value := range c.containers {
		filterResult := filter(&value)
		if filterResult != nil {
			result = append(result, *filterResult)
		}
	}

	return &containerResultImpl{
		containers: result,
	}, nil
}

// Returns a new ContainerResult object from a list of Docker Container types
func NewContainerResult(toEncapsulate []types.Container) ContainerResult {
	var containers []Container
	for _, value := range toEncapsulate {
		containers = append(containers, Container{
			rawContainerDefinition: &value,
		})
	}

	return &containerResultImpl{
		containers: containers,
	}
}