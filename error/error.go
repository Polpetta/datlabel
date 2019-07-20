package error

import (
	"fmt"
)

const (
	noContainerErrorMessage = "Container with id %s was not found"
)

type NoContainerError struct {
	containerId string
}

func NewNoContainerError(containerId string) *NoContainerError {
	return &NoContainerError{containerId: containerId}
}

func (e *NoContainerError) Error() string {
	return fmt.Sprintf(noContainerErrorMessage, e.containerId)
}
