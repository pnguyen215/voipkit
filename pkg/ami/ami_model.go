package ami

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"sync"
)

type PubChannel chan *AMIMessage
type MessageChannel map[string]PubChannel

type AMI struct {
	Mutex  sync.RWMutex       `json:"-"`
	Conn   net.Conn           `json:"-"`
	Cancel context.CancelFunc `json:"-"`
	Reader *textproto.Reader  `json:"-"`
	Writer *bufio.Writer      `json:"-"`
	Subs   *AMIPubSubQueue    `json:"-"`
	Raw    *AMIMessage        `json:"-"`
	Socket *AMISocket         `json:"-"`
	Err    chan error         `json:"-"`
}

type AMIPubSubQueue struct {
	Message MessageChannel
	Mutex   sync.RWMutex
	Off     bool
}

type AMIMessage struct {
	Header textproto.MIMEHeader
	Mutex  sync.RWMutex
}

type AMIEvent struct {
}

type AMIDictionary struct {
	AllowForceTranslate bool `json:"allow_force_translate"`
}

type AMIEventDictionary struct {
	EventKey     string            `json:"event_key,omitempty"`
	Dictionaries map[string]string `json:"dictionaries,omitempty"`
}

type AMIGrouping struct {
}

type AMIAction struct {
	ActionCmd string `json:"action_cmd"`
	Timeout   int    `json:"timeout,omitempty"`
}

type AMIResponse struct {
	Event        *AMIMessage `json:"message,omitempty"`
	IsSuccess    bool        `json:"success"`
	RawJson      string      `json:"raw_json,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

type AMIChannel struct {
	ChannelProtocol  string `json:"channel_protocol"`
	NoDigitExtension int    `json:"no_digit_extension"`
	Extension        string `json:"extension"`
}

type AMIOriginateCall struct {
	Channel   string `json:"channel" binding:"required"`
	Context   string `json:"context" binding:"required"`
	Extension string `json:"extension" binding:"required"`
	Priority  int    `json:"priority"`
	Timeout   int    `json:"timeout"`
	Variable  string `json:"variable"`
}

type AMICall struct {
	AMIOriginateCall
	NoTarget        string `json:"no_target" binding:"required"`        // like customer phone number or no. extension internal from all routes
	ChannelProtocol string `json:"channel_protocol" binding:"required"` // protocols include: SIP, H323, IAX...
	NoExtension     int    `json:"no_extension" binding:"required"`     // like current user using no. extension
}

type AMISocket struct {
	Conn                 net.Conn       `json:"-"`
	Incoming             chan string    `json:"-"`
	Shutdown             chan struct{}  `json:"-"`
	Errors               chan error     `json:"-"`
	Dictionary           *AMIDictionary `json:"-"`
	UUID                 string         `json:"uuid" binding:"required"`
	IsUsedDictionary     bool           `json:"is_used_dictionary"`
	Retry                bool           `json:"retry"`
	MaxRetries           int            `json:"max_retries"`
	AllowTrace           bool           `json:"allow_trace"`
	MaxConcurrencyMillis int64          `json:"max_concurrency_millis"`
}

type AMIResultRaw map[string]string
type AMIResultRawLevel map[string][]string

// AMICommand
// Do not set tags json on field V
type AMICommand struct {
	Action string `ami:"Action" json:"action"`
	ID     string `ami:"ActionID" json:"action_id"`
	V      []interface{}
}

type AMIAuth struct {
	Username string `ami:"Username" json:"username" binding:"required"`
	Secret   string `ami:"Secret" json:"secret" binding:"required"`
	Events   string `ami:"Events,omitempty" json:"events" binding:"required"`
}

type AMICore struct {
	Socket     *AMISocket        `json:"-"`
	Event      chan AMIResultRaw `json:"-"`
	Stop       chan struct{}     `json:"-"`
	Wg         sync.WaitGroup    `json:"-"`
	Dictionary *AMIDictionary    `json:"-"`
	UUID       string            `json:"uuid,omitempty"`
}

type AMICallbackHandler struct {
	Ctx            context.Context `json:"-"`
	Socket         AMISocket       `json:"-"`
	Command        *AMICommand     `json:"-"`
	AcceptedEvents []string        `json:"accepted_events,omitempty"`
	IgnoreEvents   []string        `json:"ignored_events,omitempty"`
}

// AMIUpdateConfigAction holds the params for an action in UpdateConfig AMI command.
//
// example
//
//	actions := make([]ami.UpdateConfigAction, 0)
//	actions = append(actions, ami.AMIUpdateConfigAction{
//		Action:   "EmptyCat",
//		Category: "test01",
//	})
//	actions = append(actions, ami.AMIUpdateConfigAction{
//		Action:   "Append",
//		Category: "test01",
//		Var:      "type",
//		Value:    "peer",
//	})
type AMIUpdateConfigAction struct {
	Action   string `ami:"Action" json:"action"`
	Category string `ami:"Category" json:"category"`
	Var      string `ami:"Var,omitempty" json:"var,omitempty"`
	Value    string `ami:"Value,omitempty" json:"value,omitempty"`
}

// AMIQueueData holds to queue calls.
// used in functions:
//
//	QueueAdd, QueueLog, QueuePause,
//	QueuePenalty, QueueReload, QueueRemove,
//	QueueReset
type AMIQueueData struct {
	Queue          string `ami:"Queue,omitempty"`
	Interface      string `ami:"Interface,omitempty"`
	Penalty        string `ami:"Penalty,omitempty"`
	Paused         string `ami:"Paused,omitempty"`
	MemberName     string `ami:"MemberName,omitempty"`
	StateInterface string `ami:"StateInterface,omitempty"`
	Event          string `ami:"Event,omitempty"`
	UniqueID       string `ami:"UniqueID,omitempty"`
	Message        string `ami:"Message,omitempty"`
	Reason         string `ami:"Reason,omitempty"`
	Members        string `ami:"Members,omitempty"`
	Rules          string `ami:"Rules,omitempty"`
	Parameters     string `ami:"Parameters,omitempty"`
}

// AMIOriginateData holds information used to originate outgoing calls.
//
//	Channel - Channel name to call.
//	Exten - Extension to use (requires Context and Priority)
//	Context - Context to use (requires Exten and Priority)
//	Priority - Priority to use (requires Exten and Context)
//	Application - Application to execute.
//	Data - Data to use (requires Application).
//	Timeout - How long to wait for call to be answered (in ms.).
//	CallerID - Caller ID to be set on the outgoing channel.
//	Variable - Channel variable to set, multiple Variable: headers are allowed.
//	Account - Account code.
//	EarlyMedia - Set to true to force call bridge on early media.
//	Async - Set to true for fast origination.
//	Codecs - Comma-separated list of codecs to use for this call.
//	ChannelId - Channel UniqueId to be set on the channel.
//	OtherChannelId - Channel UniqueId to be set on the second local channel.
type AMIOriginateData struct {
	Channel        string   `ami:"Channel,omitempty"`
	Exten          string   `ami:"Exten,omitempty"`
	Context        string   `ami:"Context,omitempty"`
	Priority       int      `ami:"Priority,omitempty"`
	Application    string   `ami:"Application,omitempty"`
	Data           string   `ami:"Data,omitempty"`
	Timeout        int      `ami:"Timeout,omitempty"`
	CallerID       string   `ami:"CallerID,omitempty"`
	Variable       []string `ami:"Variable,omitempty"`
	Account        string   `ami:"Account,omitempty"`
	EarlyMedia     string   `ami:"EarlyMedia,omitempty"`
	Async          string   `ami:"Async,omitempty"`
	Codecs         string   `ami:"Codecs,omitempty"`
	ChannelID      string   `ami:"ChannelId,omitempty"`
	OtherChannelID string   `ami:"OtherChannelId,omitempty"`
}
