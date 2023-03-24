package ami

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewAMISocket() *AMISocket {
	s := &AMISocket{}
	d := NewDictionary()
	d.SetAllowForceTranslate(true)
	s.SetDictionary(d)
	s.SetUsedDictionary(true)
	return s
}

func NewAMIResultRaw() *AMIResultRaw {
	s := &AMIResultRaw{}
	return s
}

// NewSocket provides a new socket client, connecting to a tcp server.
func NewAMISocketWith(ctx context.Context, address string) (*AMISocket, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, config.AmiNetworkTcpKey, address)
	if err != nil {
		return nil, err
	}
	return NewAMISocketConn(ctx, conn, true)
}

// NewSocket provides a new socket client, connecting to a tcp server.
// If the reuseConn = true, then using current connection.
// Otherwise, clone the connection from current connection
func NewAMISocketConn(ctx context.Context, conn net.Conn, reuseConn bool) (*AMISocket, error) {
	s := &AMISocket{
		Incoming: make(chan string, 32),
		Shutdown: make(chan struct{}),
		Errors:   make(chan error),
	}
	if reuseConn {
		s.Conn = conn
	} else {
		// checking conn available
		if conn != nil {
			var dialer net.Dialer
			_conn, err := dialer.DialContext(ctx, config.AmiNetworkTcpKey, conn.RemoteAddr().String())
			if err == nil {
				s.Conn = _conn
			}
		}
	}
	d := NewDictionary()
	d.SetAllowForceTranslate(true)
	s.SetDictionary(d)
	s.SetUsedDictionary(true)
	go s.Run(ctx, conn)
	return s, nil
}

func (s *AMISocket) SetConn(conn net.Conn) *AMISocket {
	s.Conn = conn
	return s
}

func (s *AMISocket) SetErrors(_err chan error) *AMISocket {
	s.Errors = _err
	return s
}

func (s *AMISocket) SetShutdown(_shutdown chan struct{}) *AMISocket {
	s.Shutdown = _shutdown
	return s
}

func (s *AMISocket) SetIncoming(incoming chan string) *AMISocket {
	s.Incoming = incoming
	return s
}

func (s *AMISocket) SetUUID(uuid string) *AMISocket {
	s.UUID = uuid
	return s
}

func (s *AMISocket) SetDictionary(dictionary *AMIDictionary) *AMISocket {
	s.Dictionary = dictionary
	return s
}

func (s *AMISocket) SetUsedDictionary(value bool) *AMISocket {
	s.IsUsedDictionary = value
	return s
}

func (s *AMISocket) Connected() bool {
	return s.Conn != nil
}

func (s *AMISocket) Close(ctx context.Context) error {
	close(s.Shutdown)
	if s.Connected() {
		return s.Conn.Close()
	}
	return nil
}

// Send
// Send the message to socket using fprintf format
func (s *AMISocket) Send(message string) error {
	_, err := fmt.Fprintf(s.Conn, message)
	return err
}

// Received
func (s *AMISocket) Received(ctx context.Context) (string, error) {
	var buffer bytes.Buffer
	for {
		select {
		case msg, ok := <-s.Incoming:
			if !ok {
				return buffer.String(), io.EOF
			}
			buffer.WriteString(msg)
			if strings.HasSuffix(buffer.String(), config.AmiSignalLetter) {
				return buffer.String(), nil
			}
		case err := <-s.Errors:
			return buffer.String(), err
		case <-s.Shutdown:
			return buffer.String(), io.EOF
		case <-ctx.Done():
			return buffer.String(), io.EOF
		}
	}
}

func (s *AMISocket) Run(ctx context.Context, conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			s.Errors <- err
			return
		}
		s.Incoming <- msg
	}
}

func (s AMIResultRaw) GetVal(key string) string {
	if s == nil {
		return ""
	}

	v := s[key]
	if len(v) == 0 {
		return ""
	}
	return v
}

func (s AMIResultRaw) GetValOrPref(key, pref string) string {
	_v := s.GetVal(key)

	if len(_v) == 0 {
		return s.GetVal(pref)
	}
	return _v
}

func (s AMIResultRawLevel) GetVal(key string) string {
	if s == nil {
		return ""
	}

	v := s[key]
	if len(v) == 0 {
		return ""
	}
	if len(v) == 1 {
		return v[0]
	}
	return utils.ToJson(v)
}

func (s AMIResultRawLevel) GetValOrPref(key, pref string) string {
	_v := s.GetVal(key)

	if len(_v) == 0 {
		return s.GetVal(pref)
	}
	return _v
}
