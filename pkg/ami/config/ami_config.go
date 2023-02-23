package config

import (
	"time"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/fatal"
)

const (
	AmiActionKey      = "Action"
	AmiEventKey       = "Event"
	AmiResponseKey    = "Response"
	AmiActionIdKey    = "ActionID"
	AmiLoginKey       = "Login"
	AmiCallManagerKey = "Asterisk Call Manager"
)

const (
	AmiStatusSuccessKey = "success"
	AmiStatusFailedKey  = "failed"
)

const (
	AmiUsernameField = "Username"
	AmiSecretField   = "Secret"
	AmiPasswordField = "Password"
)

const (
	AmiPubSubKeyRef = "ami-key"
)

var (
	// ErrorAsteriskConnTimeout error on connection timeout
	ErrorAsteriskConnTimeout = fatal.AMIErrorNew("Asterisk Server connection timeout")

	// ErrorAsteriskInvalidPrompt invalid prompt received from AMI server
	ErrorAsteriskInvalidPrompt = fatal.AMIErrorNew("Asterisk Server invalid prompt command line")

	// ErrorAsteriskNetwork networking errors
	ErrorAsteriskNetwork = fatal.AMIErrorNew("Network error")

	// ErrorAsteriskLogin AMI server login failed
	ErrorAsteriskLogin = fatal.AMIErrorNew("Asterisk Server login failed")

	// Error EOF
	ErrorEOF = "EOF"

	// Error I/O
	ErrorIO = "io: read/write on closed pipe"
	ErrorLoginFailed = "Failed login"
)

var (
	NetworkTimeoutAfterSeconds = time.Second * 3 // default is 3 seconds
)
