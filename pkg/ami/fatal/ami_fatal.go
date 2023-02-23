package fatal

import (
	"fmt"
)

type AMIError struct {
	S string
	E string
}

func AMIErrorNew(message string) *AMIError {
	return &AMIError{S: message}
}

func (e *AMIError) AMIError(message string, args ...interface{}) *AMIError {
	t := fmt.Sprintf(message, args...)
	e.E = fmt.Sprintf(": %s", t)
	return e
}

func (e *AMIError) Error() string {
	return fmt.Sprintf("ami has error ocurred: %s%s", e.S, e.E)
}
