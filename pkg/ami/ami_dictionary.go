package ami

import (
	"errors"
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

var overlapDictionaries *[]AMIEventDictionary = &[]AMIEventDictionary{}

func NewDictionary() *AMIDictionary {
	d := &AMIDictionary{}
	d.GetDictionaries()
	return d
}

func (d *AMIDictionary) SetAllowForceTranslate(allow bool) *AMIDictionary {
	d.AllowForceTranslate = allow
	return d
}

func (d *AMIDictionary) GetDictionaries() *[]AMIEventDictionary {

	if d.Length() > 0 {
		return overlapDictionaries
	}

	var dictionaries []AMIEventDictionary

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey: config.AmiListenerEventCommon,
		Dictionaries: map[string]string{
			"Channel":               "channel",
			"Channelstate":          "channel_state",
			"Channelstatedesc":      "channel_state_description",
			"Connectedlinename":     "connected_line_name",
			"Connectedlinenum":      "connected_line_number",
			"Context":               "context",
			"Event":                 "event",
			"Exten":                 "exten",
			"Language":              "lang",
			"Linkedid":              "linked_id",
			"Priority":              "priority",
			"Privilege":             "privilege",
			"Uniqueid":              "unique_id",
			"Calleridname":          "caller_id_name",
			"Calleridnum":           "caller_id_number",
			"Callerid":              "caller_id",
			"Accountcode":           "account_code",
			"Cause":                 "cause",
			"Cause-Txt":             "cause_text",
			"Handler":               "handler",
			"Hint":                  "hint",
			"Status":                "status",
			"Statustext":            "status_text",
			"Device":                "device",
			"State":                 "state",
			"Callstaken":            "calls_taken",
			"Incall":                "in_call",
			"Interface":             "interface",
			"Lastcall":              "last_call",
			"Lastpause":             "last_pause",
			"Membername":            "member_name",
			"Membership":            "membership",
			"Paused":                "paused",
			"Pausedreason":          "paused_reason",
			"Penalty":               "penalty",
			"Ringinuse":             "ring_in_use",
			"Queue":                 "queue",
			"Stateinterface":        "status_interface",
			"Wrapuptime":            "wrap_uptime",
			"Uptime":                "uptime",
			"Lastreload":            "last_reload",
			"Channeltype":           "channel_type",
			"Peer":                  "peer",
			"Peerstatus":            "peer_status",
			"Address":               "address",
			"Destaccountcode":       "destination_account_code",
			"Destcalleridname":      "destination_caller_id_name",
			"Destcalleridnum":       "destination_caller_id_number",
			"Variable":              "variable",
			"Value":                 "value",
			"Destchannel":           "destination_channel",
			"Destchannelstate":      "destination_channel_state",
			"Destchannelstatedesc":  "destination_channel_state_description",
			"Destconnectedlinename": "destination_connected_line_name",
			"Destconnectedlinenum":  "destination_connected_line_number",
			"Destcontext":           "destination_context",
			"Destexten":             "destination_exten",
			"Destlanguage":          "destination_lang",
			"Destlinkedid":          "destination_linked_id",
			"Destpriority":          "destination_priority",
			"Destuniqueid":          "destination_unique_id",
			"Dialstring":            "dial_string",
			"Actionid":              "action_id",
			"Message":               "message",
			"Output":                "output",
			"Response":              "response",
			"Header":                "header",
			"Mutex":                 "mutex",
			"Reason":                "reason",
			"Datereceivedat":        "date_received_at",
			"Time":                  "time",
			"Timestamp":             "timestamp",
			"Ping":                  "ping",
			"AMIversion":            "ami_version",
			"AsteriskVersion":       "asterisk_version",
			"CoreCDRenabled":        "core_cdr_enabled",
			"CoreHTTPenabled":       "core_http_enabled",
			"CoreMaxCalls":          "core_max_calls",
			"CoreMaxFilehandles":    "core_max_file_handles",
			"CoreMaxLoadAvg":        "core_max_load_avg",
			"CoreRealTimeEnabled":   "core_realtime_enabled",
			"CoreRunGroup":          "core_run_group",
			"CoreRunUser":           "core_run_user",
			"SystemName":            "system_name",
			"CoreCurrentCalls":      "core_current_calls",
			"CoreReloadDate":        "core_reload_date",
			"CoreReloadTime":        "core_reload_time",
			"CoreStartupDate":       "core_startup_date",
			"CoreStartupTime":       "core_startup_time",
			"Challenge":             "challenge",
			"JSON":                  "json",
			"Available":             "available",
			"Callers":               "callers",
			"LoggedIn":              "logged_in",
			"LongestHoldTime":       "longest_hold_time",
			"AttachMessage":         "attach_message",
			"AttachmentFormat":      "attach_format",
			"CallOperator":          "call_operator",
			"Callback":              "callback",
			"CanReview":             "can_review",
			"DeleteMessage":         "delete_message",
			"Dialout":               "dial_out",
			"Email":                 "email",
			"ExitContext":           "exit_context",
			"FromString":            "from_string",
			"Fullname":              "full_name",
			"MailCommand":           "mail_command",
			"MaxMessageCount":       "max_message_count",
			"MaxMessageLength":      "max_message_length",
			"NewMessageCount":       "new_message_count",
			"OldMessageCount":       "old_message_count",
			"Pager":                 "pager",
			"SayCID":                "say_cid",
			"SayDurationMinimum":    "say_duration_minimum",
			"SayEnvelope":           "say_envelope",
			"ServerEmail":           "server_email",
			"TimeZone":              "timezone",
			"VMContext":             "vm_context",
			"VolumeGain":            "volume_gain",
			"BridgePriority":        "bridge_priority",
			"BridgeSuspended":       "bridge_suspended",
			"CompletedFAXes":        "completed_faxes",
			"CurrentSessions":       "current_sessions",
			"FailedFAXes":           "failed_faxes",
			"ReceiveAttempts":       "received_attempts",
			"ReservedSessions":      "reserved_sessions",
			"TransmitAttempts":      "transmit_attempts",
			"Data":                  "data",
		},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey: config.AmiListenerEventCdr,
		Dictionaries: map[string]string{
			"Amaflags":           "ama_flags",
			"Answertime":         "answer_time",
			"Billableseconds":    "billable_seconds",
			"Destination":        "destination",
			"Destinationchannel": "destination_channel",
			"Destinationcontext": "destination_context",
			"Disposition":        "disposition",
			"Duration":           "duration",
			"Endtime":            "end_time",
			"Lastapplication":    "last_application",
			"Lastdata":           "last_data",
			"Source":             "source",
			"Starttime":          "start_time",
			"Userfield":          "user_field",
		},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey: config.AmiListenerEventBridgeEnter,
		Dictionaries: map[string]string{
			"Bridgecreator":         "bridge_creator",
			"Bridgename":            "bridge_name",
			"Bridgenumchannels":     "bridge_no_channels",
			"Bridgetechnology":      "bridge_technology",
			"Bridgetype":            "bridge_type",
			"Bridgeuniqueid":        "bridge_unique_id",
			"Bridgevideosourcemode": "bridge_video_source_mode",
		},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey:     config.AmiListenerEventHangupRequest,
		Dictionaries: map[string]string{},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey:     config.AmiListenerEventHangupHandlerPush,
		Dictionaries: map[string]string{},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey:     config.AmiListenerEventHangup,
		Dictionaries: map[string]string{},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey:     config.AmiListenerEventSoftHangupRequest,
		Dictionaries: map[string]string{},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey:     config.AmiListenerEventHangupHandlerRun,
		Dictionaries: map[string]string{},
	})

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey:     config.AmiListenerEventExtensionStatus,
		Dictionaries: map[string]string{},
	})

	overlapDictionaries = &dictionaries
	dictionary, _ := d.FindDictionaryByKey(config.AmiListenerEventCommon)
	log.Printf("Dictionaries (common) was initialized, len = %d", len(dictionary.Dictionaries))
	return overlapDictionaries
}

func (d *AMIDictionary) FindDictionariesByKey(eventKey string) *[]AMIEventDictionary {

	if d.Length() <= 0 {
		d.GetDictionaries()
	}

	var dictionaries []AMIEventDictionary

	for _, e := range *overlapDictionaries {

		if strings.EqualFold(eventKey, e.EventKey) {
			dictionaries = append(dictionaries, e)
		}
	}

	return &dictionaries
}

func (d *AMIDictionary) FindDictionaryByKey(eventKey string) (*AMIEventDictionary, error) {
	var dictionaries []AMIEventDictionary = *d.FindDictionariesByKey(eventKey)

	if len(dictionaries) <= 0 {
		return &AMIEventDictionary{}, errors.New("Dictionary not found.")
	}

	return &dictionaries[0], nil
}

func (d *AMIDictionary) TranslateField(field string) string {
	dictionary, _ := d.FindDictionaryByKey(config.AmiListenerEventCommon)
	value, ok := dictionary.Dictionaries[field]

	if ok {
		return strings.ToLower(value)
	} else {
		if d.AllowForceTranslate {
			value = GetValByKey(dictionary.Dictionaries, field)
			if len(value) > 0 {
				return value
			}
		}
	}

	return d.TranslateFieldWith(field, *overlapDictionaries)
}

func (d *AMIDictionary) TranslateFieldWith(field string, dictionaries []AMIEventDictionary) string {
	if len(dictionaries) <= 0 {
		return field
	}

	for _, e := range dictionaries {
		if v, ok := e.Dictionaries[field]; ok {
			return strings.ToLower(v)
		} else {
			if d.AllowForceTranslate {
				value := GetValByKey(e.Dictionaries, field)
				if len(value) > 0 {
					return value
				}
			}
		}
	}

	return field
}

func (d *AMIDictionary) TranslateKey(value string) string {
	dictionary, _ := d.FindDictionaryByKey(config.AmiListenerEventCommon)

	_key := GetKeyByVal(dictionary.Dictionaries, value)

	if !strings.EqualFold(_key, value) {
		return _key
	}

	return d.TranslateKeyWith(value, *overlapDictionaries)
}

func (d *AMIDictionary) TranslateKeyWith(value string, dictionaries []AMIEventDictionary) string {
	if len(dictionaries) <= 0 {
		return value
	}

	for _, e := range dictionaries {
		_key := GetKeyByVal(e.Dictionaries, value)
		if !strings.EqualFold(_key, value) {
			return _key
		}
	}

	return value
}

func (d *AMIDictionary) Length() int {
	return len(*overlapDictionaries)
}

func (d *AMIDictionary) Reset() {
	*overlapDictionaries = nil // overlapDictionaries = &[]AMIEventDictionary{}
	overlapDictionaries = d.GetDictionaries()
}

func (d *AMIDictionary) Json() string {
	return JsonString(*overlapDictionaries)
}

func (d *AMIDictionary) LenTranslatorCommon() int {
	dictionary, _ := d.FindDictionaryByKey(config.AmiListenerEventCommon)
	_len := len(dictionary.Dictionaries)
	log.Printf("Translator dictionaries (common) was added, len = %d", _len)
	return _len
}

func (d *AMIDictionary) AddKeyTranslator(key, value string) *AMIDictionary {
	dictionary, _ := d.FindDictionaryByKey(config.AmiListenerEventCommon)
	var pairs map[string]string = dictionary.Dictionaries
	pairs[key] = value
	dictionary.Dictionaries = pairs
	p := *overlapDictionaries

	for idx, v := range p {
		if strings.EqualFold(v.EventKey, dictionary.EventKey) {
			p[idx] = *dictionary
			break
		}
	}

	*overlapDictionaries = p
	return d
}

func (d *AMIDictionary) AddKeysTranslator(script map[string]string) *AMIDictionary {
	if len(script) > 0 {
		for k, v := range script {
			d.AddKeyTranslator(k, v)
		}
	}
	return d
}

// AddKeyLinkTranslator
// Add translator from link github
// Example:
// https://raw.githubusercontent.com/pnguyen215/gear-insights-free/master/ami.dictionaries.json
func (d *AMIDictionary) AddKeyLinkTranslator(link string) *AMIDictionary {
	keys, err := ForkDictionaryFromLink(link, false)

	if err != nil {
		return d
	}

	d.AddKeysTranslator(*keys)
	return d
}
