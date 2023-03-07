package ami

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

// NewAmi connect to Asterisk Server using net connection and try to login
func NewAmi(host string, port int, username, password string) (*AMI, error) {
	conn, err := OpenDial(host, port)
	if err != nil {
		return nil, err
	}
	return NewAmiDial(conn, username, password)
}

// NewAmiDial connect to Asterisk Server using net connection and try to login
func NewAmiDial(conn net.Conn, username, password string) (*AMI, error) {
	return NewAmiWith(conn, username, password, config.NetworkTimeoutAfterSeconds)
}

// NewAmi connect to Asterisk Server using net connection and try to login
func NewAmiWithTimeout(host string, port int, username, password string, timeout time.Duration) (*AMI, error) {
	conn, err := OpenDial(host, port)
	if err != nil {
		return nil, err
	}
	return NewAmiWith(conn, username, password, timeout)
}

// NewAmiWith connect to Asterisk Server using net connection and try to login
// using username, password and timeout after seconds. Create new Asterisk Manager Interface (AMI) and return client or error
func NewAmiWith(conn net.Conn, username, password string, timeout time.Duration) (*AMI, error) {
	client, ctx := OpenContext(conn)

	err := client.ReadPrompt(ctx, timeout)

	if err != nil {
		return nil, err
	}

	err = client.Login(ctx, timeout, username, password)

	if err != nil {
		return nil, err
	}

	client.SetReader(ctx)
	return client, nil
}

func (c *AMI) ReadPrompt(parentCtx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
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
		return config.ErrorAsteriskConnTimeout
	case err := <-fail:
		return config.ErrorAsteriskNetwork.AMIError(err.Error())
	case promptLine := <-prompt:
		if !strings.HasPrefix(promptLine, config.AmiCallManagerKey) {
			return config.ErrorAsteriskInvalidPrompt.AMIError(promptLine)
		}
	}

	return nil
}

func (c *AMI) Read() (*AMIMessage, error) {
	headers, err := c.Reader.ReadMIMEHeader()

	if err != nil {
		if err.Error() == config.ErrorEOF || err.Error() == config.ErrorIO {
			return nil, config.ErrorAsteriskNetwork
		}
		return nil, err
	}

	message := ofMessage(headers)
	return message, nil
}

func (c *AMI) Write(bytes []byte) error {
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

func (c *AMI) Login(parentContext context.Context, timeout time.Duration, username, password string) error {
	ctx, cancel := context.WithTimeout(parentContext, timeout)
	defer cancel()

	message := LoginWith(username, password)

	reader := func() (chan *AMIMessage, chan error) {
		chMessage := make(chan *AMIMessage)
		fail := make(chan error)

		go func() {
			defer close(chMessage)
			defer close(fail)

			_message, err := c.Read()

			if err != nil {
				fail <- err
				return
			}

			chMessage <- _message
		}()

		return chMessage, fail
	}

	chMessage, fail := reader()
	if err := c.Write(message.Bytes()); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return config.ErrorAsteriskConnTimeout.AMIError(config.ErrorLoginFailed)
	case err := <-fail:
		return config.ErrorAsteriskNetwork.AMIError(err.Error())
	case msg := <-chMessage:
		if !msg.IsSuccess() {
			return config.ErrorAsteriskLogin
		}
	}

	return nil
}

func (c *AMI) Publish(message *AMIMessage) {
	if message != nil {
		c.Subs.Publish(message)
	}
}

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

func (c *AMI) SetReader(ctx context.Context) {
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
				message, err := c.Read()

				if ctx.Err() != nil {
					c.EmitError(ctx.Err())
					return
				}

				if err != nil {
					if err == config.ErrorAsteriskNetwork {
						c.EmitError(err)
						return
					}
				}

				c.Publish(message)
			}
		}
	}()
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

// Action sends AMI message to Asterisk server and returns send-only
// response channel on nil
func (c *AMI) Action(message *AMIMessage) bool {
	if message.GetActionId() == "" {
		message.AddActionId()
	}

	err := c.Write(message.Bytes())
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

// RunCmd run cli core on asterisk server
func (c *AMI) RunCmd(cmd string, timeout int) (*AMIResponse, error) {
	return NewCliWith(cmd, timeout).RunCmd(c)
}

// RunCmdDictionaries
func (c *AMI) RunCmdDictionaries(cmd string, timeout int, dictionaries map[string]string) (*AMIResponse, error) {
	return NewCliWith(cmd, timeout).RunCmdDictionaries(c, dictionaries)
}
