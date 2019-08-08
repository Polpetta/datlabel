// +build unit

package result

import (
	"context"
	u "github.com/Polpetta/datlabel/test/utils"
	"github.com/Polpetta/datlabel/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"testing"
)

func TestItShouldReturnRightNumberOfContainers(t *testing.T) {

	var listOfContainers []Container
	docker := utils.NewDockerClient()

	u.DockerCli(u.NewDockerCliBuilder(u.DockerRunCommand).
		Flag(u.DockerCliDetachFlag).
		Flag(u.DockerTTYFlag).
		Image("alpine"), func(containerId string) {
		t.Logf("Got from launcher containerId: %s", containerId)
		searchFilter := filters.NewArgs()
		searchFilter.Add("id", containerId)
		containers, err := docker.ContainerList(context.Background(),
			types.ContainerListOptions{
				Filters: searchFilter,
			})

		if err != nil {
			u.KillContainer(containerId, t)
			t.Fatalf("Impossible to query the Docker daemon")
		}

		t.Logf("Returned from daemon: %+v", containers[0])
		listOfContainers = append(listOfContainers,
			Container{rawContainerDefinition: &containers[0]})
	}, t)

	u.DockerCli(u.NewDockerCliBuilder(u.DockerRunCommand).
		Flag(u.DockerCliDetachFlag).
		Flag(u.DockerTTYFlag).
		Image("alpine"), func(containerId string) {
		t.Logf("Got from launcher containerId: %s", containerId)
		searchFilter := filters.NewArgs()
		searchFilter.Add("id", containerId)
		containers, err := docker.ContainerList(context.Background(),
			types.ContainerListOptions{
				Filters: searchFilter,
			})

		if err != nil {
			u.KillContainer(containerId, t)
			t.Fatalf("Impossible to query the Docker daemon")
		}

		t.Logf("Returned from daemon: %+v", containers[0])
		listOfContainers = append(listOfContainers,
			Container{rawContainerDefinition: &containers[0]})
	}, t)

	containerListImpl := &containerResultImpl{
		containers: listOfContainers,
	}

	if len(containerListImpl.Containers()) != len(listOfContainers) {
		t.Fatalf("Size of containers launched is %d, "+
			"but got %d from ContainerResult",
			len(listOfContainers),
			len(containerListImpl.Containers()))
	}
}
