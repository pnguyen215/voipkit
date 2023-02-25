package ami

import (
	"errors"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func NewDictionary() *AMIDictionary {
	return &AMIDictionary{}
}

func (d *AMIDictionary) GetDictionaries() *[]AMIEventDictionary {
	var dictionaries []AMIEventDictionary

	dictionaries = append(dictionaries, AMIEventDictionary{
		EventKey: config.AmiListenerEventCommon,
		Dictionaries: map[string]string{
			"Channel":           "channel",
			"Channelstate":      "channel_state",
			"Channelstatedesc":  "channel_state_description",
			"Connectedlinename": "connected_line_name",
			"Connectedlinenum":  "connected_line_no",
			"Context":           "context",
			"Event":             "event",
			"Exten":             "exten",
			"Language":          "lang",
			"Linkedid":          "linked_id",
			"Priority":          "priority",
			"Privilege":         "privilege",
			"Uniqueid":          "unique_id",
			"Calleridname":      "caller_id_name",
			"Calleridnum":       "caller_id_number",
			"Callerid":          "caller_id",
			"Accountcode":       "account_code",
			"Cause":             "cause",
			"Cause-Txt":         "cause_text",
			"Handler":           "handler",
			"Hint":              "hint",
			"Status":            "status",
			"Statustext":        "status_text",
			"Device":            "device",
			"State":             "State",
			"Callstaken":        "calls_taken",
			"Incall":            "in_call",
			"Interface":         "interface",
			"Lastcall":          "last_call",
			"Lastpause":         "last_pause",
			"Membername":        "member_name",
			"Membership":        "membership",
			"Paused":            "paused",
			"Pausedreason":      "paused_reason",
			"Penalty":           "penalty",
			"Ringinuse":         "ring_in_use",
			"Queue":             "queue",
			"Stateinterface":    "status_interface",
			"Wrapuptime":        "wrap_uptime",
			"Uptime":            "uptime",
			"Lastreload":        "last_reload",
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

	return &dictionaries
}

func (d *AMIDictionary) FindDictionariesByKey(eventKey string) *[]AMIEventDictionary {

	var dictionaries []AMIEventDictionary

	for _, e := range *d.GetDictionaries() {

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
		return value
	}

	return d.TranslateFieldWith(field, *d.GetDictionaries())
}

func (d *AMIDictionary) TranslateFieldWith(field string, dictionaries []AMIEventDictionary) string {
	if len(dictionaries) <= 0 {
		return field
	}

	for _, e := range dictionaries {
		if v, ok := e.Dictionaries[field]; ok {
			return v
		}
	}

	return field
}
