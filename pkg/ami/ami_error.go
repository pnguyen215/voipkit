package ami

import "fmt"

var (
	// ErrorAsteriskConnTimeout error on connection timeout
	ErrorAsteriskConnTimeout = AMIErrorNew("Asterisk Server connection timeout")

	// ErrorAsteriskInvalidPrompt invalid prompt received from AMI server
	ErrorAsteriskInvalidPrompt = AMIErrorNew("Asterisk Server invalid prompt command line")

	// ErrorAsteriskNetwork networking errors
	ErrorAsteriskNetwork = AMIErrorNew("Network error")

	// ErrorAsteriskLogin AMI server login failed
	ErrorAsteriskLogin = AMIErrorNew("Asterisk Server login failed")

	// Error EOF
	ErrorEOF = "EOF"

	// Error I/O
	ErrorIO          = "io: read/write on closed pipe"
	ErrorLoginFailed = "Failed login"
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
