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

// This test launch two containers and then it check that
// containerResultImpl struct has two Containers in its list
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

// This test launch a container with two labels,
// and then it checks that the Container struct has the same labels
// set before
func TestItShouldReturnRightContainers(t *testing.T) {

	// FIXME sometime this test fails because nameLabel1 and nameLabel2 are
	//  swapped by Docker when the container starts!
	valueLabel1 := "value1"
	valueLabel2 := "value2"
	nameLabel1 := "com.polpetta.test1"
	nameLabel2 := "com.polpetta.test2"
	docker := utils.NewDockerClient()

	u.DockerCli(u.NewDockerCliBuilder(u.DockerRunCommand).
		Flag(u.DockerCliDetachFlag).
		Flag(u.DockerTTYFlag).
		Label(nameLabel1 + "=" + valueLabel1).
		Label(nameLabel2 + "=" + valueLabel2).
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

		if len(containers) != 1 {
			u.KillContainer(containerId, t)
			t.Fatalf("Number of containers unexpected")
		}

		t.Logf("Returned from daemon: %+v", containers[0])

		var labels []Label
		for key, value := range containers[0].Labels {
			labels = append(labels, Label{
				name:  key,
				value: value,
			})
		}

		container := Container{
			&containers[0],
			labels,
		}

		if container.Id() != containerId {
			u.KillContainer(containerId, t)
			t.Fatalf("The expected id was %s but instead got %s",
				containerId,
				container.Id())
		}

		// Check that the labels name are correct
		if container.Labels()[0].Name() != nameLabel1 {
			u.KillContainer(containerId, t)
			t.Fatalf("Label name was %s, but got %s",
				nameLabel1,
				container.Labels()[0].Name())
		}

		if container.Labels()[1].Name() != nameLabel2 {
			u.KillContainer(containerId, t)
			t.Fatalf("Label name was %s, but got %s",
				nameLabel2,
				container.Labels()[1].Name())
		}

		// Check that the labels value are correct
		if container.Labels()[0].Value() != valueLabel1 {
			u.KillContainer(containerId, t)
			t.Fatalf("Label value was %s, but got %s",
				valueLabel1,
				container.Labels()[0].Value())
		}

		if container.Labels()[1].Value() != valueLabel2 {
			u.KillContainer(containerId, t)
			t.Fatalf("Label value was %s, but got %s",
				valueLabel1,
				container.Labels()[1].Value())
		}
	}, t)
}

// This test checks the NewContainerResult function,
// assuring that the conversion between raw Docker function result
// (ContainerList) and the containerResultImpl is correctly performed
func TestNewContainerResult(t *testing.T) {
	valueLabel1 := "value1"
	valueLabel2 := "value2"
	nameLabel1 := "com.polpetta.test1"
	nameLabel2 := "com.polpetta.test2"
	docker := utils.NewDockerClient()

	u.DockerCli(u.NewDockerCliBuilder(u.DockerRunCommand).
		Flag(u.DockerCliDetachFlag).
		Flag(u.DockerTTYFlag).
		Label(nameLabel1 + "=" + valueLabel1).
		Label(nameLabel2 + "=" + valueLabel2).
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

		if len(containers) != 1 {
			u.KillContainer(containerId, t)
			t.Fatalf("Expected containers length to be 1, but got %d",
				len(containers))
		}

		newContainerResultList := NewContainerResult(containers)

		if len(newContainerResultList.Containers()) != 1 {
			u.KillContainer(containerId, t)
			t.Fatalf("Expected container list to be of size 1,"+
				" but got %d",
				len(newContainerResultList.Containers()))
		}

		if len(newContainerResultList.Containers()[0].Labels()) != 2 {
			u.KillContainer(containerId, t)
			t.Fatalf("Expected to have 2 labels, but got %d",
				len(newContainerResultList.Containers()[0].Labels()))
		}
	}, t)
}