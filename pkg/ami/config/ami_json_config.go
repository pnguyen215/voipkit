package config

const (
	AmiJsonFieldObjectName = "object_name"
	AmiJsonFieldActionId   = "action_id"
	AmiJsonFieldEvent      = "event"
	AmiJsonFieldLastReload = "last_reload"
	AmiJsonFieldPrivilege  = "privilege"
	AmiJsonFieldStatus     = "status"
	AmiJsonFieldUptime     = "uptime"
	AmiJsonFieldResponse   = "response"
	AmiJsonFieldMessage    = "message"
	AmiJsonFieldUniqueId   = "unique_id"
	AmiJsonFieldPing       = "ping"
)

var (
	AmiJsonIgnoringFieldType map[string]bool = map[string]bool{
		AmiJsonFieldActionId:   true,
		AmiJsonFieldEvent:      true,
		AmiJsonFieldLastReload: true,
		AmiJsonFieldPrivilege:  true,
		AmiJsonFieldStatus:     true,
		AmiJsonFieldUptime:     true,
		AmiJsonFieldResponse:   true,
		AmiJsonFieldMessage:    true,
		AmiJsonFieldUniqueId:   true,
		AmiJsonFieldPing:       true,
	}
)
