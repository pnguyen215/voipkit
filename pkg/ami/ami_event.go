package ami

import (
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func NewEventListener() *AMIEvent {
	return &AMIEvent{}
}

// Listen All Events
func (e *AMIEvent) OpenFullEvents(c *AMI) {
	all := c.AllEvents()
	defer c.Close()

	for {
		select {
		case message := <-all:
			log.Printf("ami event: '%s' received = %s", message.Field(strings.ToLower(config.AmiEventKey)), message.Json())
		case err := <-c.Error():
			c.Close()
			log.Fatalf("ami listener has error occurred = %s", err.Error())
		}
	}
}

// Listen Event by key name
func (e *AMIEvent) OpenEvent(c *AMI, name string) {
	event := c.OnEvent(name)
	defer c.Close()

	for {
		select {
		case message := <-event:
			log.Printf("ami event: '%s' received = %s", name, message.Json())
		case err := <-c.Error():
			c.Close()
			log.Fatalf("ami listener event: '%s' has error occurred = %s", name, err.Error())
		}
	}
}

// Lister Events by collect of keys string
func (e *AMIEvent) OpenEvents(c *AMI, keys ...string) {
	event := c.OnEvents(keys...)
	defer c.Close()

	for {
		select {
		case message := <-event:
			log.Printf("ami event(s): '%s' received = %s", keys, message.Json())
		case err := <-c.Error():
			c.Close()
			log.Fatalf("ami listener event(s): '%s' has error occurred = %s", keys, err.Error())
		}
	}
}


// Listen CDR Event - Call Detail Record, include: link playback (example: .wav, .mp3)
func (e *AMIEvent) OpenCdrEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventCdr)
}

// Listen Bridge Enter Event to mark point connected state
func (e *AMIEvent) OpenBridgeEnterEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventBridgeEnter)
}

func (e *AMIEvent) OpenConnectedEvent(c *AMI) {
	e.OpenBridgeEnterEvent(c)
}

func (e *AMIEvent) OpenDeviceHangupRequestEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventSoftHangupRequest)
}

func (e *AMIEvent) OpenHangupFinishEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventHangup)
}

func (e *AMIEvent) OpenHangupRequestEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventHangupRequest)
}
