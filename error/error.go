package error

import (
	"fmt"
)

const (
	noContainerErrorMessage = "Container with id %s was not found"
)

type NoSuchElement struct {
	elementId string
}

func NewNoSuchElement(elementId string) *NoSuchElement {
	return &NoSuchElement{elementId: elementId}
}

func (e *NoSuchElement) Error() string {
	return fmt.Sprintf(noContainerErrorMessage, e.elementId)
}
