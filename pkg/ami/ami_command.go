package ami

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
	"golang.org/x/exp/slices"
)

func NewCommand() *AMICommand {
	a := &AMICommand{}
	return a
}

func (a *AMICommand) SetAction(action string) *AMICommand {
	a.Action = action
	return a
}

func (a *AMICommand) SetId(id string) *AMICommand {
	a.ID = id
	return a
}

func (a *AMICommand) SetV(v ...interface{}) *AMICommand {
	a.V = v
	return a
}

func (a *AMICommand) SetVCmd(v interface{}) *AMICommand {
	a.SetV(v)
	return a
}

func (a *AMICommand) AppendV(v ...interface{}) *AMICommand {
	a.V = append(a.V, v...)
	return a
}

func (a *AMICommand) TransformCommand(c *AMICommand) ([]byte, error) {
	if len(c.Action) <= 0 {
		return nil, fmt.Errorf("Invalid 'Action'")
	}
	if len(c.ID) <= 0 {
		c.ID, _ = GenUUID()
		log.Printf("TransformCommand for Action = %v has been generated ID = %v", c.Action, c.ID)
	}
	return Marshal(c)
}

// Send
func (a *AMICommand) Send(ctx context.Context, socket AMISocket, c *AMICommand) (AmiReply, error) {
	b, err := a.TransformCommand(c)
	if err != nil {
		return nil, err
	}
	if err := socket.Send(string(b)); err != nil {
		return nil, err
	}
	return a.Read(ctx, socket)
}

// SendLevel
func (a *AMICommand) SendLevel(ctx context.Context, socket AMISocket, c *AMICommand) (AmiReplies, error) {
	b, err := a.TransformCommand(c)
	if err != nil {
		return nil, err
	}
	if err := socket.Send(string(b)); err != nil {
		return nil, err
	}
	return a.ReadLevel(ctx, socket)
}

// DoGetResult
// Get result while fetch response command has been sent to asterisk server
// Arguments:
// 1. AMISocket - to create new instance connection socket
// 2. AMICommand - to build command cli will be sent to server
// 3. acceptedEvents - select event will captured as response
// 4. ignoreEvents - the event will been stopped fetching command
func (a *AMICommand) DoGetResult(ctx context.Context, s AMISocket, c *AMICommand, acceptedEvents []string, ignoreEvents []string) ([]AmiReply, error) {
	bytes, err := c.TransformCommand(c)

	if err != nil {
		return nil, err
	}

	if err := s.Send(string(bytes)); err != nil {
		return nil, err
	}

	response := make([]AmiReply, 0)

	for {
		raw, err := c.Read(ctx, s)
		if err != nil {
			return nil, err
		}
		_event := raw.Get(strings.ToLower(config.AmiEventKey))
		_response := raw.Get(strings.ToLower(config.AmiResponseKey))

		if len(acceptedEvents) == 0 {
			if s.DebugMode {
				log.Printf(config.AmiErrorMissingSocketEvent, _event, _response)
			}
			break
		}

		if len(ignoreEvents) > 0 {
			if slices.Contains(ignoreEvents, _event) || (_response != "" && !strings.EqualFold(_response, config.AmiStatusSuccessKey)) {
				if s.DebugMode {
					log.Printf(config.AmiErrorBreakSocketIgnoredEvent, _event)
				}
				break
			}
		}

		if slices.Contains(acceptedEvents, _event) {
			response = append(response, raw)
		}
	}

	return response, nil
}

// Read
func (a *AMICommand) Read(ctx context.Context, socket AMISocket) (AmiReply, error) {
	var buffer bytes.Buffer
	var concurrency int64 = 0
	_start := time.Now().UnixMilli()
	for {
		input, err := socket.Received(ctx)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(input)
		_end := time.Now().UnixMilli() - _start
		concurrency += _end

		if socket.MaxConcurrencyMillis > 0 {
			if concurrency >= socket.MaxConcurrencyMillis {
				if socket.DebugMode {
					D().Warn("Read(). max over concurrency: %v (ms) and the concurrency allowed: %v (ms)",
						concurrency, socket.MaxConcurrencyMillis)
				}
				break
			}
		}

		if strings.HasSuffix(buffer.String(), config.AmiSignalLetters) {
			break
		}
	}
	return ParseReply(socket, buffer.String())
}

// ReadLevel
func (a *AMICommand) ReadLevel(ctx context.Context, socket AMISocket) (AmiReplies, error) {
	var buffer bytes.Buffer
	var concurrency int64 = 0
	_start := time.Now().UnixMilli()
	for {
		input, err := socket.Received(ctx)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(input)
		_end := time.Now().UnixMilli() - _start
		concurrency += _end

		if socket.MaxConcurrencyMillis > 0 {
			if concurrency >= socket.MaxConcurrencyMillis {
				if socket.DebugMode {
					D().Warn("ReadLevel(). max over concurrency: %v (ms) and the concurrency allowed: %v (ms)",
						concurrency, socket.MaxConcurrencyMillis)
				}
				break
			}
		}

		if strings.HasSuffix(buffer.String(), config.AmiSignalLetters) {
			break
		}
	}
	return ParseReplies(socket, buffer.String())
}
