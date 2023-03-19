package ami

import (
	"context"
	"log"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func NewCore() *AMICore {
	c := &AMICore{}
	return c
}

func (c *AMICore) SetSocket(socket *AMISocket) *AMICore {
	c.Socket = socket
	return c
}

func (c *AMICore) SetUUID(id string) *AMICore {
	c.UUID = id
	return c
}

func (c *AMICore) SetEvent(event chan AMIResultRaw) *AMICore {
	c.Event = event
	return c
}

func (c *AMICore) SetStop(stop chan struct{}) *AMICore {
	c.Stop = stop
	return c
}

func (c *AMICore) SetDictionary(dictionary *AMIDictionary) *AMICore {
	c.Dictionary = dictionary
	return c
}

// NewAmiCore
// Creating new instance asterisk server connection
// Firstly, create new instance AMISocket
// Secondly, create new request body to login
func NewAmiCore(ctx context.Context, socket AMISocket, auth *AMIAuth) (*AMICore, error) {
	uuid, err := GenUUID()

	if err != nil {
		return nil, err
	}

	socket.SetUUID(uuid)
	err = Login(ctx, socket, auth)

	if err != nil {
		return nil, err
	}

	core := NewCore()
	core.SetSocket(&socket)
	core.SetUUID(uuid)
	core.SetEvent(make(chan AMIResultRaw))
	core.SetStop(make(chan struct{}))
	core.SetDictionary(socket.Dictionary)

	core.Wg.Add(1)
	go core.Run(ctx)
	return core, nil
}

// Run
// Go-func to consume event from asterisk server response
func (c *AMICore) Run(ctx context.Context) {
	defer c.Wg.Done()
	for {
		select {
		case <-c.Stop:
			return
		case <-ctx.Done():
			return
		default:
			event, err := Events(ctx, *c.Socket)
			if err != nil {
				log.Printf(config.AmiErrorConsumeEvent, err)
				return
			}
			c.Event <- event
		}
	}
}

// Events
// Consume all events will be returned an channel received from asterisk server log.
func (c *AMICore) Events() <-chan AMIResultRaw {
	return c.Event
}

// GetSIPPeers
// GetSIPPeers fetch the list of SIP peers present on asterisk.
func (c *AMICore) GetSIPPeers(ctx context.Context) ([]AMIResultRaw, error) {
	return SIPPeers(ctx, *c.Socket)
}

// Logoff
// Logoff closes the current session with AMI.
func (c *AMICore) Logoff(ctx context.Context) error {
	close(c.Stop)
	c.Wg.Wait()
	return Logoff(ctx, *c.Socket)
}
