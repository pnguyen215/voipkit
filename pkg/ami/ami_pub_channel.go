package ami

import (
	"log"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewPubSubQueue() *AMIPubSubQueue {
	c := &AMIPubSubQueue{}
	c.message = make(MessageChannel)
	return c
}

func (k *AMIPubSubQueue) TurnOff() {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.Off = true
}

func (k *AMIPubSubQueue) TurnOn() {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.Off = false
}

func (k *AMIPubSubQueue) Destroy() {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	if k.Off {
		log.Println("Destroy pub-sub stopped")
		return
	}

	k.Off = true
	for key, ch := range k.message {
		close(ch)
		delete(k.message, key)
	}
}

func (k *AMIPubSubQueue) SizeMessage() int {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	return len(k.message)
}

func (k *AMIPubSubQueue) Subscribe(key string) PubChannel {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	if k.Off {
		return nil
	}

	key = strings.ToLower(key)

	ch := make(PubChannel)

	if _, ok := k.message[key]; !ok {
		k.message[key] = ch
	}

	return ch
}

func (k *AMIPubSubQueue) Subscribes(keys ...string) PubChannel {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	if k.Off {
		return nil
	}
	ch := make(PubChannel, len((keys)))
	for _, key := range keys {
		key = strings.ToLower(key)
		if _, ok := k.message[key]; !ok {
			k.message[key] = ch
			// ch <- <-k.Message[key]
		}
	}
	// close(ch)
	return ch
}

// Publish broadcasts the provided AMI message to all subscribers interested in the corresponding event type.
// It also broadcasts the message to subscribers interested in all events.
// Returns true if the message is successfully published; otherwise, returns false.
//
// Example:
//
//	pubSubQueue.Publish(amiMessage)
//
// Note: The AMI Pub-Sub mechanism allows subscribers to receive notifications for specific events or all events.
// This method ensures that the message is sent to relevant subscribers based on event type and general subscriptions.
func (k *AMIPubSubQueue) Publish(message *AMIMessage) bool {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	if k.Off {
		return false
	}
	ch, ok := k.message[config.AmiPubSubKeyRef]
	if ok {
		go func(ch PubChannel) {
			ch <- message
		}(ch)
	}
	name := strings.ToLower(message.Field(strings.ToLower(config.AmiEventKey)))
	if name != "" {
		if ch, ok := k.message[name]; ok {
			go func(ch PubChannel) {
				ch <- message
			}(ch)
		}
	}
	return true
}
