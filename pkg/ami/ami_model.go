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
	Conn     net.Conn      `json:"-"`
	Incoming chan string   `json:"-"`
	Shutdown chan struct{} `json:"-"`
	Errors   chan error    `json:"-"`
}

type AMISocketRaw map[string][]string

type AMICommand struct {
	Action string `ami:"Action" json:"action"`
	ID     string `ami:"ActionID,omitempty" json:"action_id"`
	V      []interface{}
}
