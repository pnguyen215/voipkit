package ami

import (
	"log"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewPubSubQueue() *AMIPubSubQueue {
	c := &AMIPubSubQueue{}
	c.Message = make(MessageChannel)
	return c
}

func (k *AMIPubSubQueue) Disabled() {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()
	k.Off = true
}

func (k *AMIPubSubQueue) Destroy() {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()

	if k.Off {
		log.Println("Destroy pub-sub stopped")
		return
	}

	k.Off = true
	for key, ch := range k.Message {
		close(ch)
		delete(k.Message, key)
	}
}

func (k *AMIPubSubQueue) SizeMessage() int {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()
	return len(k.Message)
}

func (k *AMIPubSubQueue) Subscribe(key string) PubChannel {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()

	if k.Off {
		return nil
	}

	key = strings.ToLower(key)

	ch := make(PubChannel)

	if _, ok := k.Message[key]; !ok {
		k.Message[key] = ch
	}

	return ch
}

func (k *AMIPubSubQueue) Subscribes(keys ...string) PubChannel {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()

	if k.Off {
		return nil
	}

	ch := make(PubChannel, len((keys)))

	for _, key := range keys {
		key = strings.ToLower(key)
		if _, ok := k.Message[key]; !ok {
			k.Message[key] = ch
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
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	if k.Off {
		return false
	}
	ch, ok := k.Message[config.AmiPubSubKeyRef]
	if ok {
		go func(ch PubChannel) {
			ch <- message
		}(ch)
	}
	name := strings.ToLower(message.Field(strings.ToLower(config.AmiEventKey)))
	if name != "" {
		if ch, ok := k.Message[name]; ok {
			go func(ch PubChannel) {
				ch <- message
			}(ch)
		}
	}
	return true
}
