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
)

func NewAMISocket() *AMISocket {
	s := &AMISocket{}
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
	return NewAMISocketConn(ctx, conn)
}

// NewSocket provides a new socket client, connecting to a tcp server.
func NewAMISocketConn(ctx context.Context, conn net.Conn) (*AMISocket, error) {
	s := &AMISocket{
		Conn:     conn,
		Incoming: make(chan string, 32),
		Shutdown: make(chan struct{}),
		Errors:   make(chan error),
	}
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
