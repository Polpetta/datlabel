package dockerlabel

import (
	m "github.com/polpetta/docker-label"
	u "github.com/polpetta/docker-label/test/utils"
	"testing"
)

// Here we need to launch a dummy container with labels and then check
// that GetLabelsFromContainer returns the correct list of labels
func TestItShouldGetLabelsFromContainer(t *testing.T) {
	testLabel := "com.polpetta.test.getLabels"
	u.DockerCli(u.NewDockerCliBuilder(u.DockerRunCommand).
		Flag(u.DockerCliDetachFlag).
		Flag(u.DockerTTYFlag).
		Label(testLabel).
		Image("alpine"), func(containerId string) {
		labels, err := m.GetLabelsFromContainer(containerId)
		if err != nil {
			u.KillContainer(containerId, t)
			t.Fatalf(err.Error())
		}
		if len(labels) != 1 || labels[0] != testLabel {
			u.KillContainer(containerId, t)
			t.Fatalf("Expected 1 element in the array, found %d",
				len(labels))
		}

		t.Log(labels[0])
	}, t)
}

// Here we do the opposite, we check that an empty list is returned when
// a container with no labels is checked
func TestItShouldNotGetLabelsFromEmptyContainer(t *testing.T) {
	u.DockerCli(u.NewDockerCliBuilder(u.DockerRunCommand).
		Flag(u.DockerCliDetachFlag).
		Flag(u.DockerTTYFlag).
		Image("alpine"), func(containerId string) {
		labels, err := m.GetLabelsFromContainer(containerId)
		if err != nil {
			u.KillContainer(containerId, t)
			t.Fatalf(err.Error())
		}
		if len(labels) != 0 {
			u.KillContainer(containerId, t)
			t.Fatalf("Expected 0 elements in the array, found %d",
				len(labels))
		}
	}, t)
}

// Here we test if an error is returned when the label list of a
// non-existing container is asked
func TestItShouldReturnErrorIfContainerIdIsInvalid(t *testing.T) {
	labels, err := m.GetLabelsFromContainer("dummyId")
	if err == nil {
		for _, v := range labels {
			t.Logf("Label value: %s", v)
		}
		t.Fatalf("Expected error to have non-nil value")
	}
}

// We are now testing the service functionality.
func TestItShouldGetLabelsFromService(t *testing.T) {
	testLabel := "com.polpetta.test.getLabels"
	u.DockerCli(u.NewDockerCliBuilder(u.DockerServiceCommand).
		Command(u.DockerCreateCommand).
		Flag(u.DockerCliDetachFlag).
		Label(testLabel).
		Name("testService").
		Image("nginx"), func(serviceId string) {
		labels, err := m.GetLabelsFromService(serviceId)
		if err != nil {
			u.KillService(serviceId, t)
			t.Fatalf(err.Error())
		}
		if len(labels) != 1 || labels[0] != testLabel {
			u.KillService(serviceId, t)
			t.Fatalf("Expected 1 element in the array, found %d",
				len(labels))
		}

		t.Log(labels[0])
	}, t)
}

// Here we do the opposite, we check that an empty list is returned when
// a service with no labels is checked
func TestItShouldNotGetLabelsFromEmptyService(t *testing.T) {
	u.DockerCli(u.NewDockerCliBuilder(u.DockerServiceCommand).
		Command(u.DockerCreateCommand).
		Flag(u.DockerCliDetachFlag).
		Image("nginx"), func(serviceId string) {
		labels, err := m.GetLabelsFromService(serviceId)
		if err != nil {
			u.KillService(serviceId, t)
			t.Fatalf(err.Error())
		}
		if len(labels) != 0 {
			u.KillService(serviceId, t)
			t.Fatalf("Expected 0 elements in the array, found %d",
				len(labels))
		}
	}, t)
}

// Here we test if an error is returned when the label list of a
// non-existing container is asked
func TestItShouldReturnErrorIfServiceIdIsInvalid(t *testing.T) {
	labels, err := m.GetLabelsFromService("dummyId")
	if err == nil {
		for _, v := range labels {
			t.Logf("Label value: %s", v)
		}
		t.Fatalf("Expected error to have non-nil value")
	}
}
