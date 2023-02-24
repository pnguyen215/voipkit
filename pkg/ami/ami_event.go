package ami

import (
	"log"
)

func NewEventListener() *AMIEvent {
	return &AMIEvent{}
}

func (e *AMIEvent) OpenFullEvents(c *AMI) {
	all := c.AllEvents()
	defer c.Close()

	for {
		select {
		case message := <-all:
			log.Printf("ami event received = %s", message.Json())
		case err := <-c.Error():
			c.Close()
			log.Fatalf("ami listener has error occurred = %s", err.Error())
		}
	}
}

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
