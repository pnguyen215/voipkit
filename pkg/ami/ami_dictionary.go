package ami

import (
	"errors"
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

var overlapDictionaries *[]AMIEventDictionary = &[]AMIEventDictionary{}

func NewDictionary() *AMIDictionary {
	d := &AMIDictionary{}
	d.GetDictionaries()
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
	log.Printf("Dictionaries was initialized, len = %d", d.Length())
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
		}
	}

	return field
}

func (d *AMIDictionary) TranslateKey(value string) string {
	dictionary, _ := d.FindDictionaryByKey(config.AmiListenerEventCommon)

	_key := utils.TakeKeyFromValue(dictionary.Dictionaries, value)

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
		_key := utils.TakeKeyFromValue(e.Dictionaries, value)
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
	return utils.ToJson(*overlapDictionaries)
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
