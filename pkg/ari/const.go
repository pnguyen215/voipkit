package ari

const (
	keyAnyMessage = "ari-key"
)

var (
	// ErrorAsteriskConnTimeout error on connection timeout
	ErrorAsteriskConnTimeout = AsteriskErrorNew("Asterisk Server connection timeout")

	// ErrorAsteriskInvalidPrompt invalid prompt received from AMI server
	ErrorAsteriskInvalidPrompt = AsteriskErrorNew("Asterisk Server invalid prompt command line")

	// ErrorAsteriskNetwork networking errors
	ErrorAsteriskNetwork = AsteriskErrorNew("Network error")

	// ErrorAsteriskLogin AMI server login failed
	ErrorAsteriskLogin = AsteriskErrorNew("Asterisk Server login failed")

	// Error EOF
	ErrorEOF = "EOF"

	// Error I/O
	ErrorIO = "io: read/write on closed pipe"
)
