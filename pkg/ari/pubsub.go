package ari

import (
	"strings"
	"sync"
)

type pubChan chan *Message // DONE

type messageChanSpy map[string]pubChan // DONE

type pubsub struct { // DONE
	mu  sync.RWMutex
	msg messageChanSpy
	off bool
}

func newPubsub() *pubsub { // DONE
	p := &pubsub{}
	p.msg = make(messageChanSpy)
	return p
}

func (ps *pubsub) disable() { // DONE
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.off = true
}

func (ps *pubsub) lenMessageChanSpy() int { // DONE
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return len(ps.msg)
}

func (ps *pubsub) destroy() { // DONE
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.off {
		return
	}
	ps.off = true
	for key, ch := range ps.msg {
		close(ch)
		delete(ps.msg, key)
	}
}

// subscribe to event by name or by action id as key
func (ps *pubsub) subscribe(key string) pubChan { // DONE
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.off {
		return nil
	}
	key = strings.ToLower(key)

	ch := make(pubChan)
	if _, ok := ps.msg[key]; !ok {
		ps.msg[key] = ch
	}

	return ch
}

func (ps *pubsub) publish(msg *Message) bool { // DONE
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if ps.off {
		return false
	}

	if ch, ok := ps.msg[keyAnyMessage]; ok {
		go func(ch pubChan) {
			ch <- msg
		}(ch)
	}

	if name := strings.ToLower(msg.Field("event")); name != "" {
		if ch, ok := ps.msg[name]; ok {
			go func(ch pubChan) {
				ch <- msg
			}(ch)
		}
	}

	return true
}
