package error

import (
	"fmt"
)

const (
	noContainerErrorMessage = "Container with id %s was not found"
)

// An shorthand error to say that the desired element is not existing
type NoSuchElement struct {
	elementId string
}

// Create a new NoSuchElement error, with the associated id
func NewNoSuchElement(elementId string) *NoSuchElement {
	return &NoSuchElement{elementId: elementId}
}

// Prints the error. See noContainerErrorMessage to get the output of the
// error itself
func (e *NoSuchElement) Error() string {
	return fmt.Sprintf(noContainerErrorMessage, e.elementId)
}
