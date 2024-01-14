package ami

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAmiSocket() *AMISocket {
	s := &AMISocket{
		incoming:  make(chan string, 32),
		shutdown:  make(chan struct{}),
		errors:    make(chan error),
		DebugMode: false,
	}
	d := NewDictionary()
	d.SetEnabledForceTranslate(true)

	s.SetDictionary(d)
	s.SetUsedDictionary(true)
	s.SetRetry(true)
	s.SetMaxRetries(3)
	s.SetMaxConcurrencyMillis(-1)
	return s
}

func (s *AMISocket) SetRetry(value bool) *AMISocket {
	s.Retry = value
	return s
}

func (s *AMISocket) SetMaxRetries(value int) *AMISocket {
	s.MaxRetries = value
	return s
}

func (s *AMISocket) SetMaxConcurrencyMillis(value int64) *AMISocket {
	s.MaxConcurrencyMillis = value
	return s
}

func (s *AMISocket) Json() string {
	return JsonString(s)
}

// NewSocket provides a new socket client, connecting to a tcp server.
func WithSocket(ctx context.Context, address string) (*AMISocket, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, config.AmiNetworkTcpKey, address)
	if err != nil {
		return nil, err
	}
	return WithAmiSocketOver(ctx, conn, true)
}

// WithAmiSocketOver provides a new socket client, connecting to a tcp server.
// If the reuseConn = true, then using current connection.
// Otherwise, clone the connection from current connection
func WithAmiSocketOver(ctx context.Context, conn net.Conn, reuseConn bool) (*AMISocket, error) {
	s := NewAmiSocket()
	if reuseConn {
		s.conn = conn
	} else {
		if conn != nil {
			var dialer net.Dialer
			_conn, err := dialer.DialContext(ctx, config.AmiNetworkTcpKey, conn.RemoteAddr().String())
			if err == nil {
				s.conn = _conn
			}
		}
	}
	go s.Run(ctx, conn)
	return s, nil
}

func (s *AMISocket) SetConn(conn net.Conn) *AMISocket {
	s.conn = conn
	return s
}

func (s *AMISocket) SetErrors(_err chan error) *AMISocket {
	s.errors = _err
	return s
}

func (s *AMISocket) SetShutdown(_shutdown chan struct{}) *AMISocket {
	s.shutdown = _shutdown
	return s
}

func (s *AMISocket) SetIncoming(incoming chan string) *AMISocket {
	s.incoming = incoming
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

func (s *AMISocket) SetDebugMode(value bool) *AMISocket {
	s.DebugMode = value
	return s
}

func (s *AMISocket) ResetUUID() *AMISocket {
	s.SetUUID(GenUUIDShorten())
	return s
}

func (s *AMISocket) Connected() bool {
	return s.conn != nil
}

func (s *AMISocket) Close(ctx context.Context) error {
	close(s.shutdown)
	if s.Connected() {
		return s.conn.Close()
	}
	return nil
}

// Send
// Send the message to socket using fprintf format
func (s *AMISocket) Send(message string) error {
	v, err := fmt.Fprintf(s.conn, message)
	if s.DebugMode {
		D().Info("Ami command, the number of %v byte(s) written and message: %v", v, message)
	}
	return err
}

// Received
func (s *AMISocket) Received(ctx context.Context) (string, error) {
	var buffer bytes.Buffer
	for {
		select {
		case msg, ok := <-s.incoming:
			if !ok {
				return buffer.String(), io.EOF
			}
			buffer.WriteString(msg)
			if strings.HasSuffix(buffer.String(), config.AmiSignalLetter) {
				return buffer.String(), nil
			}
		case err := <-s.errors:
			return buffer.String(), err
		case <-s.shutdown:
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
			s.errors <- err
			return
		}
		s.incoming <- msg
	}
}
