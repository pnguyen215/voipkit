package config

const (
	AmiJsonFieldObjectName         = "object_name"
	AmiJsonFieldActionId           = "action_id"
	AmiJsonFieldEvent              = "event"
	AmiJsonFieldLastReload         = "last_reload"
	AmiJsonFieldPrivilege          = "privilege"
	AmiJsonFieldStatus             = "status"
	AmiJsonFieldUptime             = "uptime"
	AmiJsonFieldResponse           = "response"
	AmiJsonFieldMessage            = "message"
	AmiJsonFieldUniqueId           = "unique_id"
	AmiJsonFieldPing               = "ping"
	AmiJsonFieldLinkedId           = "linked_id"
	AmiJsonFieldChannel            = "channel"
	AmiJsonFieldDestChannel        = "destination_channel"
	AmiJsonFieldExten              = "exten"
	AmiJsonFieldContext            = "context"
	AmiJsonFieldPeer               = "peer"
	AmiJsonFieldChannelType        = "channel_type"
	AmiJsonFieldPeerStatus         = "peer_status"
	AmiJsonFieldTime               = "time"
	AmiJsonFieldHint               = "hint"
	AmiJsonFieldStatusText         = "status_text"
	AmiJsonFieldPublishedAt        = "published_at"
	AmiJsonFieldAccountCode        = "account_code"
	AmiJsonFieldAmaFlags           = "ama_flags"
	AmiJsonFieldAnswerTime         = "answer_time"
	AmiJsonFieldBillableSeconds    = "billable_seconds"
	AmiJsonFieldCallerId           = "caller_id"
	AmiJsonFieldDateReceivedAt     = "date_received_at"
	AmiJsonFieldDestination        = "destination"
	AmiJsonFieldDestinationChannel = "destination_channel"
	AmiJsonFieldDestinationContext = "destination_context"
	AmiJsonFieldDisposition        = "disposition"
	AmiJsonFieldDuration           = "duration"
	AmiJsonFieldEndTime            = "end_time"
	AmiJsonFieldLastApplication    = "last_application"
	AmiJsonFieldLastData           = "last_data"
	AmiJsonFieldSource             = "source"
	AmiJsonFieldStartTime          = "start_time"
	AmiJsonFieldUserField          = "user_field"
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
