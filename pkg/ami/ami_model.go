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

// AMIPayloadQueue holds to queue calls.
// used in functions:
//
//	QueueAdd, QueueLog, QueuePause,
//	QueuePenalty, QueueReload, QueueRemove,
//	QueueReset
type AMIPayloadQueue struct {
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

// AMIPayloadOriginate holds information used to originate outgoing calls.
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
type AMIPayloadOriginate struct {
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

// AMIPayloadCall holds the call data to transfer.
type AMIPayloadCall struct {
	Channel       string `ami:"Channel"`
	ExtraChannel  string `ami:"ExtraChannel,omitempty"`
	Exten         string `ami:"Exten"`
	ExtraExten    string `ami:"ExtraExten,omitempty"`
	Context       string `ami:"Context"`
	ExtraContext  string `ami:"ExtraContext,omitempty"`
	Priority      string `ami:"Priority"`
	ExtraPriority string `ami:"ExtraPriority,omitempty"`
}

// AMIPayloadAOC holds the information to generate an Advice of Charge message on a channel.
//
//	Channel - Channel name to generate the AOC message on.
//	ChannelPrefix -	Partial channel prefix. By using this option one can match the beginning part of a channel name without having to put the entire name in.
//					For example if a channel name is SIP/snom-00000001 and this value is set to SIP/snom, then that channel matches and the message will be sent.
//					Note however that only the first matched channel has the message sent on it.
//
//	MsgType - Defines what type of AOC message to create, AOC-D or AOC-E
//		D
//		E
//
//	ChargeType - Defines what kind of charge this message represents.
//		NA
//		FREE
//		Currency
//		Unit
//
//	UnitAmount(0) -	This represents the amount of units charged. The ETSI AOC standard specifies that this value along with the optional UnitType value are entries in a list.
//					To accommodate this these values take an index value starting at 0 which can be used to generate this list of unit entries.
//					For Example, If two unit entires were required this could be achieved by setting the paramter UnitAmount(0)=1234 and UnitAmount(1)=5678.
//					Note that UnitAmount at index 0 is required when ChargeType=Unit, all other entries in the list are optional.
//
//	UnitType(0) -	Defines the type of unit. ETSI AOC standard specifies this as an integer value between 1 and 16, but this value is left open to accept any positive integer.
//					Like the UnitAmount parameter, this value represents a list entry and has an index parameter that starts at 0.
//	CurrencyName - Specifies the currency's name. Note that this value is truncated after 10 characters.
//	CurrencyAmount - Specifies the charge unit amount as a positive integer. This value is required when ChargeType==Currency.
//
//	CurrencyMultiplier - Specifies the currency multiplier. This value is required when ChargeType==Currency.
//		OneThousandth
//		OneHundredth
//		OneTenth
//		One
//		Ten
//		Hundred
//		Thousand
//
//	TotalType - Defines what kind of AOC-D total is represented.
//		Total
//		SubTotal
//
//	AOCBillingId - Represents a billing ID associated with an AOC-D or AOC-E message. Note that only the first 3 items of the enum are valid AOC-D billing IDs
//		Normal
//		ReverseCharge
//		CreditCard
//		CallFwdUnconditional
//		CallFwdBusy
//		CallFwdNoReply
//		CallDeflection
//		CallTransfer
//
//	ChargingAssociationId - 	Charging association identifier. This is optional for AOC-E and can be set to any value between -32768 and 32767
//	ChargingAssociationNumber -	Represents the charging association party number. This value is optional for AOC-E.
//	ChargingAssociationPlan - 	Integer representing the charging plan associated with the ChargingAssociationNumber.
//								The value is bits 7 through 1 of the Q.931 octet containing the type-of-number and numbering-plan-identification fields.
type AMIPayloadAOC struct {
	Channel                   string `ami:"Channel"`
	ChannelPrefix             string `ami:"ChannelPrefix"`
	MsgType                   string `ami:"MsgType"`
	ChargeType                string `ami:"ChargeType"`
	UnitAmount                string `ami:"UnitAmount(0)"`
	UnitType                  string `ami:"UnitType(0)"`
	CurrencyName              string `ami:"CurrencyName"`
	CurrencyAmount            string `ami:"CurrencyAmount"`
	CurrencyMultiplier        string `ami:"CurrencyMultiplier"`
	TotalType                 string `ami:"TotalType"`
	AOCBillingID              string `ami:"AOCBillingId"`
	ChargingAssociationID     string `ami:"ChargingAssociationId"`
	ChargingAssociationNumber string `ami:"ChargingAssociationNumber"`
	ChargingAssociationPlan   string `ami:"ChargingAssociationPlan"`
}

type AMIPayloadMonitor struct {
	Channel      string `ami:"Channel"`
	Direction    string `ami:"Direction,omitempty"`
	State        string `ami:"State,omitempty"`
	File         string `ami:"File, omitempty"`
	Format       string `ami:"Format,omitempty"`
	Mix          bool   `ami:"Mix,omitempty"`
	MixMonitorID string `ami:"MixMonitorID,omitempty"`
}

// AMIPayloadMessage holds the message data to message send command.
type AMIPayloadMessage struct {
	To         string `ami:"To"`
	From       string `ami:"From"`
	Body       string `ami:"Body"`
	Base64Body string `ami:"Base64Body,omitempty"`
	Variable   string `ami:"Variable"`
}

// AMIPayloadKhompSMS holds the Khomp SMS information.
type AMIPayloadKhompSMS struct {
	Device       string `ami:"Device"`
	Destination  string `ami:"Destination"`
	Confirmation bool   `ami:"Confirmation"`
	Message      string `ami:"Message"`
}
