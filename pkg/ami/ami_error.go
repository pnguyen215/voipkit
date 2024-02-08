package ami

import "fmt"

var (
	// ErrorAsteriskConnTimeout error on connection timeout
	ErrorAsteriskConnTimeout = AmiErrorWrap("Asterisk Server connection timeout")

	// ErrorAsteriskInvalidPrompt invalid prompt received from AMI server
	ErrorAsteriskInvalidPrompt = AmiErrorWrap("Asterisk Server invalid prompt command line")

	// ErrorAsteriskNetwork networking errors
	ErrorAsteriskNetwork = AmiErrorWrap("Network error")

	// ErrorAsteriskAuthenticated AMI server authenticated unsuccessful
	ErrorAsteriskAuthenticated = AmiErrorWrap("Asterisk Server authenticated unsuccessful")

	// Error messages
	ErrorEOF                         = "EOF"
	ErrorIO                          = "io: read/write on closed pipe"
	ErrorAuthenticatedUnsuccessfully = "Authenticated unsuccessful"
)

type AmiError struct {
	S string
	E string
}

func AmiErrorWrap(message string) *AmiError {
	return &AmiError{S: message}
}

func (e *AmiError) ErrorWrap(message string, args ...interface{}) *AmiError {
	t := fmt.Sprintf(message, args...)
	e.E = fmt.Sprintf(": %s", t)
	return e
}

func (e *AmiError) Error() string {
	return fmt.Sprintf("AMI_ERR: %s%s", e.S, e.E)
}
