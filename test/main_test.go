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
		Label(testLabel).
		Image("alpine"), func(containerId string) {
		labels, err := m.GetLabelsFromContainer(containerId)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if len(labels) != 1 && labels[0] == testLabel {
			t.Fatalf("Expected 1 element in the array, found %d",
				len(labels))
		}
	}, t)
}

// Here we do the opposite, we check that an empty list is returned when
// a container with no labels is checked
func TestItShouldNotGetLabelsFromEmptyContainer(t *testing.T) {
	u.DockerCli(u.NewDockerCliBuilder(u.DockerRunCommand).
		Image("alpine"), func(containerId string) {
		labels, err := m.GetLabelsFromContainer(containerId)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if len(labels) != 0 {
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
