package config

const (
	AmiManagerPerm   = "system,call,all,user"
	AmiSysPerm       = "system"
	AmiCallPerm      = "call"
	AmiAllPerm       = "all"
	AmiUserPerm      = "user"
	AmiLogPerm       = "log"
	AmiVerbosePerm   = "verbose"
	AmiCommandPerm   = "command"
	AmiAgentPerm     = "agent"
	AmiOriginatePerm = "originate"
	AmiConfigPerm    = "config"
	AmiDTMFPerm      = "dtmf"
	AmiReportingPerm = "reporting"
	AmiCdrPerm       = "cdr"
	AmiDialplanPerm  = "dialplan"
)

var (
	AmiPerms map[string]bool = map[string]bool{
		AmiSysPerm:       true,
		AmiCallPerm:      true,
		AmiAllPerm:       true,
		AmiUserPerm:      true,
		AmiLogPerm:       true,
		AmiVerbosePerm:   true,
		AmiCommandPerm:   true,
		AmiAgentPerm:     true,
		AmiOriginatePerm: true,
		AmiConfigPerm:    true,
		AmiDTMFPerm:      true,
		AmiReportingPerm: true,
		AmiCdrPerm:       true,
		AmiDialplanPerm:  true,
	}
)
