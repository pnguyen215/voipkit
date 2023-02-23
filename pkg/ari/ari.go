package ari

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"strings"
	"sync"
	"time"
)

// NetworkTimeoutAfterSeconds timeout for net connections and requests
var NetworkTimeoutAfterSeconds = time.Second * 3 // default is 3 seconds

// AriClient structure
type AriClient struct {
	mu     sync.RWMutex
	reader *textproto.Reader
	writer *bufio.Writer
	conn   net.Conn
	cancel context.CancelFunc
	subs   *pubsub
	err    chan error
}

func NewClient(conn net.Conn, username, password string) (*AriClient, error) { // DONE
	return NewClientWith(conn, username, password, NetworkTimeoutAfterSeconds)
}

// NewClientWith connect to Asterisk Server using net connection and try to login
// using username, password and timeout after seconds. Create new Asterisk Manager Interface (AMI) and return client or error
func NewClientWith(conn net.Conn, username, password string, timeout time.Duration) (*AriClient, error) { // DONE
	client, ctx := newClient(conn)

	err := client.readPrompt(ctx, timeout)

	if err != nil {
		return nil, err
	}

	err = client.login(ctx, timeout, username, password)

	if err != nil {
		return nil, err
	}

	client.initReader(ctx)
	return client, nil
}

// AllEvents subscribes to any AMI message received from Asterisk server
// returns send-only channel or nil
func (c *AriClient) AllEvents() <-chan *Message { // DONE
	return c.subs.subscribe(keyAnyMessage)
}

// OnEvent subscribes by event by name (case insensitive) and
// returns send-only channel or nil
func (c *AriClient) OnEvent(name string) <-chan *Message { // DONE
	return c.subs.subscribe(name)
}

// Action sends AMI message to Asterisk server and returns send-only
// response channel on nil
func (c *AriClient) Action(message *Message) bool { // DONE
	if message.ActionID() == "" {
		message.AddActionID()
	}

	if err := c.write(message.Bytes()); err != nil {
		c.emitError(err)
		return false
	}
	return true
}

// Close client and destroy all subscriptions to events and action responses
func (c *AriClient) Close() { // DONE
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cancel()
	c.subs.destroy()
	c.conn.Close()
	close(c.err)
	c.err = nil
}

// Error returns channel for error and signals that client should be probably restarted
func (c *AriClient) Error() <-chan error { // DONE
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.err
}

func newClient(conn net.Conn) (*AriClient, context.Context) { // DONE
	ctx, cancel := context.WithCancel(context.Background())
	client := &AriClient{
		reader: textproto.NewReader(bufio.NewReader(conn)),
		writer: bufio.NewWriter(conn),
		conn:   conn,
		cancel: cancel,
	}
	return client, ctx
}

func (c *AriClient) readPrompt(parentCtx context.Context, timeout time.Duration) error { // DONE
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	reader := func() (chan string, chan error) {
		prompt := make(chan string)
		fail := make(chan error)
		go func() {
			defer close(prompt)
			defer close(fail)
			line, err := c.reader.ReadLine()
			if err != nil {
				fail <- err
				return
			}
			prompt <- line
		}()
		return prompt, fail
	}

	prompt, fail := reader()
	select {
	case <-ctx.Done():
		return ErrorAsteriskConnTimeout
	case err := <-fail:
		return ErrorAsteriskNetwork.AsteriskErrorWith(err.Error())
	case promptLine := <-prompt:
		if !strings.HasPrefix(promptLine, "Asterisk Call Manager") {
			return ErrorAsteriskInvalidPrompt.AsteriskErrorWith(promptLine)
		}
	}
	return nil
}

func (c *AriClient) login(parentCtx context.Context, timeout time.Duration, username, password string) error { // DONE
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	msg := loginMessage(username, password)

	reader := func() (chan *Message, chan error) {
		chMsg := make(chan *Message)
		fail := make(chan error)
		go func() {
			defer close(chMsg)
			defer close(fail)
			msg, err := c.read()
			if err != nil {
				fail <- err
				return
			}
			chMsg <- msg
		}()
		return chMsg, fail
	}

	chMsg, fail := reader()
	if err := c.write(msg.Bytes()); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ErrorAsteriskConnTimeout.AsteriskErrorWith("failed login")
	case err := <-fail:
		return ErrorAsteriskNetwork.AsteriskErrorWith(err.Error())
	case msg := <-chMsg:
		if !msg.IsSuccess() {
			return ErrorAsteriskLogin
		}
	}

	return nil
}

func (c *AriClient) initReader(ctx context.Context) { // DONE
	c.subs = newPubsub()
	c.err = make(chan error)

	go func() {
		defer c.subs.disable()
		for {
			select {
			case <-ctx.Done():
				c.emitError(ctx.Err())
				return
			default:
				message, err := c.read()
				if ctx.Err() != nil {
					c.emitError(ctx.Err())
					return // already closed
				}
				if err != nil {
					if err == ErrorAsteriskNetwork {
						c.emitError(err)
						return
					}
				}
				c.publish(message)
			}
		}
	}()
}

func (c *AriClient) publish(msg *Message) { // DONE
	if msg != nil {
		c.subs.publish(msg)
	}
}

func (c *AriClient) emitError(err error) { // DONE
	go func(err error) {
		c.mu.Lock()
		defer c.mu.Unlock()
		if err == nil || c.err == nil {
			return
		}
		// c.mu.Lock()
		// defer c.mu.Unlock()
		c.err <- err
	}(err)
}

func (c *AriClient) write(bytes []byte) error { // DONE
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, err := c.writer.Write(bytes); err != nil {
		return err
	}
	if err := c.writer.Flush(); err != nil {
		return err
	}
	return nil
}

func (c *AriClient) read() (*Message, error) { // DONE
	headers, err := c.reader.ReadMIMEHeader()
	if err != nil {
		if err.Error() == ErrorEOF || err.Error() == ErrorIO {
			return nil, ErrorAsteriskNetwork
		}
		return nil, err
	}
	msg := newMessage(headers)
	return msg, nil
}
