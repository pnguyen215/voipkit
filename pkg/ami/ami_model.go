package ami

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"sync"
	"time"
)

type PubChannel chan *AMIMessage
type MessageChannel map[string]PubChannel
type tcpAmiFactory struct{}
type udpAmiFactory struct{}

type AmiClient struct {
	enabled  bool
	host     string
	port     int
	username string
	password string
	timeout  time.Duration
}

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
	Message MessageChannel `json:"-"`
	Mutex   sync.RWMutex   `json:"-"`
	Off     bool           `json:"off"`
}

type AMIMessage struct {
	Header         textproto.MIMEHeader `json:"header,omitempty"`
	Mutex          sync.RWMutex         `json:"-"`
	DateTimeLayout string               `json:"date_time_layout,omitempty"`
	PhonePrefix    []string             `json:"phone_prefix,omitempty"`
	Region         string               `json:"region,omitempty"`
}

type AMIEvent struct {
	DateTimeLayout string   `json:"date_time_layout,omitempty"`
	PhonePrefix    []string `json:"phone_prefix,omitempty"`
	Region         string   `json:"region,omitempty"`
}

type AMIDictionary struct {
	EnabledForceTranslate bool `json:"enabled_force_translate"`
}

type AMIEventDictionary struct {
	EventKey     string            `json:"event_key,omitempty"`
	Dictionaries map[string]string `json:"dictionaries,omitempty"`
}

type AMIGrouping struct {
}

type AMIAction struct {
	Name    string `json:"name"`
	Timeout int    `json:"timeout,omitempty"`
}

type AMIResponse struct {
	Event        *AMIMessage `json:"-"`
	IsSuccess    bool        `json:"success"`
	RawJson      string      `json:"raw_json,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Raw          interface{} `json:"raw,omitempty"`
}

type AMIChannel struct {
	ChannelProtocol  string `json:"channel_protocol"`
	NoDigitExtension int    `json:"no_digit_extension"`
	Extension        string `json:"extension"`
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
	DebugMode            bool           `json:"debug_mode"`
	MaxConcurrencyMillis int64          `json:"max_concurrency_millis"`
}

type AmiReply map[string]string
type AmiReplies map[string][]string

// AMICommand
// Do not set tags json on field V
type AMICommand struct {
	Action string `ami:"Action" json:"action"`
	ID     string `ami:"ActionID" json:"action_id"`
	V      []interface{}
}

type AMIAuth struct {
	Username string `ami:"Username" json:"username" binding:"required"`
	Secret   string `ami:"Secret" json:"-" binding:"required"`
	Events   string `ami:"Events,omitempty" json:"events" binding:"required"`
}

type AMICore struct {
	Socket     *AMISocket     `json:"-"`
	Event      chan AmiReply  `json:"-"`
	Stop       chan struct{}  `json:"-"`
	Wg         sync.WaitGroup `json:"-"`
	Dictionary *AMIDictionary `json:"-"`
	UUID       string         `json:"uuid,omitempty"`
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

// AMIOriginate holds information used to originate outgoing calls.
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
type AMIOriginate struct {
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
	File         string `ami:"File,omitempty"`
	Format       string `ami:"Format,omitempty"`
	Mix          bool   `ami:"Mix,omitempty"`
	MixMonitorId string `ami:"MixMonitorID,omitempty"`
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

// AMIPayloadExtension holds the extension data to dialplan.
type AMIPayloadExtension struct {
	Context         string `ami:"Context"`
	Extension       string `ami:"Extension"`
	Priority        string `ami:"Priority,omitempty"`
	Application     string `ami:"Application,omitempty"`
	ApplicationData string `ami:"ApplicationData,omitempty"`
	Replace         string `ami:"Replace,omitempty"`
}

type AMIPayloadDb struct {
	Family string `ami:"Family"`
	Key    string `ami:"Key"`
	Value  string `ami:"Val,omitempty"`
}

type AMIDialCall struct {
	Telephone       string `json:"telephone" binding:"required"`        // like customer phone number or no. extension internal from all routes
	ChannelProtocol string `json:"channel_protocol" binding:"required"` // protocols include: SIP, H323, IAX...
	Extension       int    `json:"extension" binding:"required"`        // like current user using no. extension
	DebugMode       bool   `json:"debug_mode"`                          // allow to trace log
	Timeout         int    `json:"timeout"`                             // set timeout while calling
	ExtensionExists bool   `json:"extension_exists"`                    // allow validate the extension and channel
}

type AMIExtensionStatesConf struct {
	ExtensionState int    `json:"extension_state"`
	Text           string `json:"text"`
}

type AMIConf struct {
}

type AMIBase struct {
	ActionId string `json:"action_id,omitempty"`
	Response string `json:"response,omitempty"`
	Message  string `json:"message,omitempty"`
}

type AMIExtensionStatus struct {
	AMIBase
	Context    string `json:"context"`
	Extension  string `json:"extension"`
	Hint       string `json:"hint"`
	Status     int    `json:"status"`
	StatusText string `json:"status_text"`
}

type AMIExtensionGuard struct {
	AllowExtensionNumeric bool     `json:"allow_extension_numeric"`
	Context               []string `json:"ctx"`
	StatusesText          []string `json:"statuses_text"`
}

type AMIPeerStatus struct {
	AMIBase
	ChannelType    string    `json:"channel_type"`
	Event          string    `json:"event"`
	Peer           string    `json:"peer"`
	PeerStatus     string    `json:"peer_status"`
	Privilege      string    `json:"privilege,omitempty"`
	TimeInMs       int       `json:"time_ms,omitempty"`
	PrePublishedAt string    `json:"pre_published_at"`
	PublishedAt    time.Time `json:"published_at"`
}

type AMIPeerStatusGuard struct {
	DateTimeLayout string `json:"date_time_layout"`
	Timezone       string `json:"timezone"`
}

type AMICdr struct {
	Event string `json:"event"`
	// The account code of the Party A channel.
	AccountCode string `json:"account_code,omitempty"`
	// The Caller ID number associated with the Party A in the CDR.
	Source string `json:"source,omitempty"`
	// The dialplan extension the Party A was executing.
	Destination string `json:"destination,omitempty"`
	// The dialplan context the Party A was executing.
	DestinationContext string `json:"destination_context,omitempty"`
	// The Caller ID name associated with the Party A in the CDR.
	CallerId string `json:"caller_id,omitempty"`
	// The channel name of the Party A.
	Channel string `json:"channel"`
	// The channel name of the Party B.
	DestinationChannel string `json:"destination_channel,omitempty"`
	// The last dialplan application the Party A executed.
	LastApplication string `json:"last_application,omitempty"`
	// The parameters passed to the last dialplan application the Party A executed.
	LastData string `json:"last_data,omitempty"`
	// The time the CDR was created.
	StartTime time.Time `json:"start_time"`
	// The earliest of either the time when Party A answered, or the start time of this CDR.
	AnswerTime time.Time `json:"answer_time,omitempty"`
	// The time when the CDR was finished. This occurs when the Party A hangs up or when the bridge between Party A and Party B is broken.
	EndTime time.Time `json:"end_time"`
	// The time, in seconds, of EndTime - StartTime.
	Duration int `json:"duration"`
	// The time, in seconds, of AnswerTime - StartTime.
	BillableSeconds int `json:"billable_seconds,omitempty"`
	// The final known disposition of the CDR.
	// 	NO ANSWER - The channel was not answered. This is the default disposition.
	//	FAILED - The channel attempted to dial but the call failed.
	//	BUSY - The channel attempted to dial but the remote party was busy.
	//	ANSWERED - The channel was answered. The hang up cause will no longer impact the disposition of the CDR.
	//	CONGESTION - The channel attempted to dial but the remote party was congested.
	Disposition string `json:"disposition"`
	// A flag that informs a billing system how to treat the CDR.
	//	OMIT - This CDR should be ignored.
	//	BILLING - This CDR contains valid billing data.
	//	DOCUMENTATION - This CDR is for documentation purposes.
	AmaFlags string `json:"ama_flags,omitempty"`
	// A unique identifier for the Party A channel.
	UniqueId string `json:"unique_id"`
	// A user defined field set on the channels. If set on both the Party A and Party B channel, the user-fields of both are concatenated and separated by a ;.
	UserField      string    `json:"user_field,omitempty"`
	DateReceivedAt time.Time `json:"date_received_at"`
	Privilege      string    `json:"privilege,omitempty"`
	Direction      string    `json:"direction"` // inbound, outbound
	Desc           string    `json:"desc"`
	Type           string    `json:"type"`
	UserExtension  string    `json:"user_extension,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	PlaybackUrl    string    `json:"playback_url,omitempty"` // the only cdr has status answered
	symbol         string    `json:"-"`                      // default extension splitter symbol: -, example: SIP/1000-00098fec then split by -
}

type AMIPayloadChanspy struct {
	Join            string `json:"join" binding:"required"`
	SourceExten     string `json:"source_extension" binding:"required"`
	CurrentExten    string `json:"current_extension" binding:"required"`
	ChannelProtocol string `json:"channel_protocol" binding:"required"` // protocols include: SIP, H323, IAX...
	DebugMode       bool   `json:"debug_mode"`                          // allow to trace log
}
