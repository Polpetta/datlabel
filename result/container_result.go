package result

import (
	"github.com/docker/docker/api/types"
)

type Container struct {
	Id     string
	Docker *types.Container
}

type ContainerResult interface {
	Result
	Containers() []Container
	Filter(filter func(container *Container) *Container) (ContainerResult, error)
}
