package ami

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"strings"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func (c *AMI) EmitError(err error) {
	go func(err error) {
		c.Mutex.Lock()
		defer c.Mutex.Unlock()
		if err == nil || c.Err == nil {
			return
		}
		// c.mu.Lock()
		// defer c.mu.Unlock()
		c.Err <- err
	}(err)
}

// Error returns channel for error and signals that client should be probably restarted
func (c *AMI) Error() <-chan error {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.Err
}

// Close client and destroy all subscriptions to events and action responses
func (c *AMI) Close() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Cancel()
	c.Subs.Destroy()
	c.Conn.Close()
	close(c.Err)
	c.Err = nil
}

// Action send AMI message to Asterisk server and returns send-only
// response channel on nil
func (c *AMI) Action(message *AMIMessage) bool {
	if message.GetActionId() == "" {
		message.AddActionId()
	}
	err := c.write(message.Bytes())
	if err != nil {
		c.EmitError(err)
		return false
	}
	return true
}

// AllEvents subscribes to any AMI message received from Asterisk server
// returns send-only channel or nil
func (c *AMI) AllEvents() <-chan *AMIMessage {
	return c.Subs.Subscribe(config.AmiPubSubKeyRef)
}

// OnEvent subscribes by event name (case insensitive) and
// returns send-only channel or nil
func (c *AMI) OnEvent(name string) <-chan *AMIMessage {
	return c.Subs.Subscribe(name)
}

// OnEvents subscribes by events name (case insensitive) and
// return send-only channel or nil
func (c *AMI) OnEvents(keys ...string) <-chan *AMIMessage {
	return c.Subs.Subscribes(keys...)
}

func (c *AMI) publish(message *AMIMessage) {
	if message != nil {
		c.Subs.Publish(message)
	}
}

func (c *AMI) apply(_ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(_ctx, timeout)
	defer cancel()
	reader := func() (chan string, chan error) {
		prompt := make(chan string)
		fail := make(chan error)
		go func() {
			defer close(prompt)
			defer close(fail)
			line, err := c.Reader.ReadLine()
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
		return ErrorAsteriskNetwork.ErrorWrap(err.Error())
	case promptLine := <-prompt:
		if !strings.HasPrefix(promptLine, config.AmiCallManagerKey) {
			return ErrorAsteriskInvalidPrompt.ErrorWrap(promptLine)
		}
	}
	return nil
}

func (c *AMI) read() (*AMIMessage, error) {
	headers, err := c.Reader.ReadMIMEHeader()
	if err != nil {
		if err.Error() == ErrorEOF || err.Error() == ErrorIO {
			return nil, ErrorAsteriskNetwork
		}
		return nil, err
	}
	message := ofMessage(headers)
	c.Raw = message
	return message, nil
}

func (c *AMI) write(bytes []byte) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	_, err := c.Writer.Write(bytes)
	if err != nil {
		return err
	}
	err = c.Writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (c *AMI) release(ctx context.Context) {
	c.Subs = NewPubSubQueue()
	c.Err = make(chan error)
	go func() {
		defer c.Subs.Disabled()
		for {
			select {
			case <-ctx.Done():
				c.EmitError(ctx.Err())
				return
			default:
				message, err := c.read()
				if ctx.Err() != nil {
					c.EmitError(ctx.Err())
					return
				}
				if err != nil {
					if err == ErrorAsteriskNetwork {
						c.EmitError(err)
						return
					}
				}
				c.Raw = message
				c.publish(message)
			}
		}
	}()
}

func (c *AMI) authenticate(_ctx context.Context, request AmiClient) error {
	ctx, cancel := context.WithTimeout(_ctx, request.timeout)
	defer cancel()
	message := Authenticate(request.username, request.password)
	reader := func() (chan *AMIMessage, chan error) {
		chMessage := make(chan *AMIMessage)
		fail := make(chan error)
		go func() {
			defer close(chMessage)
			defer close(fail)
			_message, err := c.read()
			if err != nil {
				fail <- err
				return
			}
			chMessage <- _message
		}()
		return chMessage, fail
	}
	message_reader, fail := reader()
	if err := c.write(message.Bytes()); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ErrorAsteriskConnTimeout.ErrorWrap(ErrorAuthenticatedUnsuccessfully)
	case err := <-fail:
		return ErrorAsteriskNetwork.ErrorWrap(err.Error())
	case msg := <-message_reader:
		if !msg.IsSuccess() {
			return ErrorAsteriskAuthenticated
		} else {
			c.Raw = msg
		}
	}
	return nil
}

// create initializes a new AMI client with the provided network connection.
// It returns the initialized AMI client along with a cancellable context for managing its lifecycle.
//
// Parameters:
//   - conn: The network connection used by the AMI client.
//
// Returns:
//   - An initialized AMI client (*AMI).
//   - A cancellable context for managing the AMI client's lifecycle (context.Context).
//
// Example:
//
//	// Creating an AMI client with a network connection
//	conn, err := net.Dial("tcp", "localhost:5038")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	client, ctx := create(conn)
//	// Use the client and context for AMI operations.
//	// Make sure to close the connection and cancel the context when done.
//	defer client.Close()
//	defer client.Cancel()
func create(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	c := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
	}
	if conn != nil {
		addr := conn.RemoteAddr().String()
		_socket, err := WithSocket(ctx, addr)
		if err == nil {
			c.Socket = _socket
			D().Info("Ami cloning (addr: %v) socket connection succeeded", addr)
		}
	}
	return c, ctx
}

func serve(conn net.Conn, request AmiClient) (*AMI, error) {
	ins, ctx := create(conn)
	err := ins.apply(ctx, request.timeout)
	if err != nil {
		return nil, err
	}
	err = ins.authenticate(ctx, request)
	if err != nil {
		return nil, err
	}
	ins.release(ctx)
	return ins, nil
}
