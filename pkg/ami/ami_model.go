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
	Mutex  sync.RWMutex
	Conn   net.Conn
	Cancel context.CancelFunc
	Reader *textproto.Reader
	Writer *bufio.Writer
	Subs   *AMIPubSubQueue
	Err    chan error
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

type AMICommand struct {
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
