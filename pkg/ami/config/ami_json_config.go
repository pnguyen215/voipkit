package config

const (
	AmiJsonFieldObjectName  = "object_name"
	AmiJsonFieldActionId    = "action_id"
	AmiJsonFieldEvent       = "event"
	AmiJsonFieldLastReload  = "last_reload"
	AmiJsonFieldPrivilege   = "privilege"
	AmiJsonFieldStatus      = "status"
	AmiJsonFieldUptime      = "uptime"
	AmiJsonFieldResponse    = "response"
	AmiJsonFieldMessage     = "message"
	AmiJsonFieldUniqueId    = "unique_id"
	AmiJsonFieldPing        = "ping"
	AmiJsonFieldLinkedId    = "linked_id"
	AmiJsonFieldChannel     = "channel"
	AmiJsonFieldDestChannel = "destination_channel"
	AmiJsonFieldExten       = "exten"
	AmiJsonFieldContext     = "context"
	AmiJsonFieldPeer        = "peer"
	AmiJsonFieldChannelType = "channel_type"
	AmiJsonFieldPeerStatus  = "peer_status"
	AmiJsonFieldTime        = "time"
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
