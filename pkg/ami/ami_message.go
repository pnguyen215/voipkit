package ami

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/textproto"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/model"
)

type amiMessage model.AMIMessage

func NewActionWith(name string) *amiMessage {
	a := NewMessage()
	a.AddField(config.ASTERISK_ACTION_KEY, name)
	return a
}

func NewMessage() *amiMessage {
	header := make(textproto.MIMEHeader)
	return ofMessage(header)
}

func ofMessage(header textproto.MIMEHeader) *amiMessage {
	m := &amiMessage{}
	m.Header = header
	return m
}

// Field return first value associated with then given key
func (k *amiMessage) Field(key string) string {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	return k.Header.Get(key)
}

// Return first value associated of field by the given key
func (k *amiMessage) GetFirstValueByField(key string) string {
	return k.Field(key)
}

// Field Values return all values associated with the given key
func (k *amiMessage) FieldValues(key string) []string {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	return k.Header.Values(key)
}

// Return all value associated by the given key
func (k *amiMessage) GetValuesByField(key string) []string {
	return k.FieldValues(key)
}

// Added new pair header as form key:value
func (k *amiMessage) AddField(key, value string) {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()
	k.Header.Add(key, value)
}

// Added collection pair header as form key:value
func (k *amiMessage) AddFields(fields map[string]string) {
	for key, value := range fields {
		k.AddField(key, value)
	}
}

// Added new Action-Id and insert into message
func (k *amiMessage) AddActionIdWith(id string) {
	k.AddField(config.ASTERISK_ACTION_ID_KEY, id)
}

// Added Action-id generated random
func (k *amiMessage) AddActionId() {
	b := make([]byte, 12)

	_, err := rand.Read(b)

	if err == nil {
		k.AddActionIdWith(fmt.Sprintf("%x", b))
	}
}

// Return AMI message Action Id as string
func (k *amiMessage) GetActionId() string {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	return k.Header.Get(strings.ToLower(config.ASTERISK_ACTION_ID_KEY))
}

// Remove fields associated with the given key
func (k *amiMessage) RemoveField(key string) {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()
	k.Header.Del(key)
}

// Remove fields associated with the collection given keys
func (k *amiMessage) RemoveFields(keys ...string) {
	for _, key := range keys {
		k.RemoveField(key)
	}
}

// Return new bytes buffer
func (k *amiMessage) toBytesBuffer() bytes.Buffer {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	var buffer bytes.Buffer

	for key := range k.Header {
		for _, value := range k.Header.Values(key) {
			buffer.WriteString(key)
			buffer.WriteString(": ")
			buffer.WriteString(value)
			buffer.WriteString("\r\n")
		}
	}
	buffer.WriteString("\r\n")

	return buffer
}

// Return headers Asterisk Manager Interface (AMI) as string
func (k *amiMessage) String() string {
	_buffer := k.toBytesBuffer()
	return _buffer.String()
}

// Return headers Asterisk Manager Interface (AMI) as byte array
func (k *amiMessage) Bytes() []byte {
	_buffer := k.toBytesBuffer()
	return _buffer.Bytes()
}

// Return true if the AMI message is event type
func (k *amiMessage) IsEvent() bool {
	return k.Field(config.ASTERISK_EVENT_KEY) != ""
}

// Return true if the AMI message is response type
func (k *amiMessage) IsResponse() bool {
	return k.Field(config.ASTERISK_RESPONSE_KEY) != ""
}

// Return true if the AMI message is response type with value success
func (k *amiMessage) IsSuccess() bool {
	return strings.EqualFold(k.Field(config.ASTERISK_RESPONSE_KEY), config.ASTERISK_STATUS_SUCCESS_KEY)
}
