package model

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"sync"
)

type pubChannel chan *AMIMessage
type messageChannel map[string]pubChannel

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
	Message messageChannel
	Mutex   sync.RWMutex
	Off     bool
}

type AMIMessage struct {
	Header textproto.MIMEHeader
	Mutex  sync.RWMutex
}
