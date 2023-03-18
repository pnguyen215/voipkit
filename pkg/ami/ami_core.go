package ami

import (
	"context"
	"log"
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

func NewAmiCore(ctx context.Context, socket AMISocket, auth *AMIAuth) (*AMICore, error) {
	uuid, err := GenUUID()

	if err != nil {
		return nil, err
	}

	auth.ID = uuid
	err = Login(ctx, socket, auth)

	if err != nil {
		return nil, err
	}

	core := NewCore()
	core.SetSocket(&socket)
	core.SetUUID(uuid)
	core.SetEvent(make(chan AMIResultRaw))
	core.SetStop(make(chan struct{}))

	core.Wg.Add(1)
	go core.Run(ctx)
	return core, nil
}

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
				log.Printf("Ami events failed: %v\n", err)
				return
			}
			c.Event <- event
		}
	}
}

// Events returns an channel with events received from AMI.
func (c *AMICore) Events() <-chan AMIResultRaw {
	return c.Event
}

// SIPPeers fetch the list of SIP peers present on asterisk.
func (c *AMICore) SIPPeers(ctx context.Context) ([]AMIResultRaw, error) {
	return SIPPeers(ctx, *c.Socket, c.UUID)
}

// Logoff closes the current session with AMI.
func (c *AMICore) Logoff(ctx context.Context) error {
	close(c.Stop)
	c.Wg.Wait()
	return Logoff(ctx, *c.Socket, c.UUID)
}
