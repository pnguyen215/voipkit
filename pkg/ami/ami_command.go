package ami

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
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

func (a *AMICommand) Send(ctx context.Context, socket AMISocket, c *AMICommand) (AMIResultRaw, error) {
	b, err := a.TransformCommand(c)
	if err != nil {
		return nil, err
	}
	if err := socket.Send(string(b)); err != nil {
		return nil, err
	}
	return a.Read(ctx, socket)
}

func (a *AMICommand) Read(ctx context.Context, socket AMISocket) (AMIResultRaw, error) {
	var buffer bytes.Buffer
	for {
		input, err := socket.Received(ctx)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(input)
		if strings.HasSuffix(buffer.String(), config.AmiSignalLetters) {
			break
		}
	}
	return ParseResult(socket, buffer.String())
}
