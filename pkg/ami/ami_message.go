package ami

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/textproto"
	"reflect"
	"strings"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func WithMessage(name string) *AMIMessage {
	a := NewMessage()
	a.AddField(config.AmiActionKey, name)
	return a
}

func NewMessage() *AMIMessage {
	header := make(textproto.MIMEHeader)
	return ofMessage(header)
}

func (m *AMIMessage) SetTimeFormat(value string) *AMIMessage {
	m.TimeFormat = value
	return m
}

func (m *AMIMessage) SetPhonePrefix(value []string) *AMIMessage {
	m.PhonePrefix = value
	return m
}

func (m *AMIMessage) AppendPhonePrefix(values ...string) *AMIMessage {
	m.PhonePrefix = append(m.PhonePrefix, values...)
	return m
}

func (m *AMIMessage) SetRegion(value string) *AMIMessage {
	m.Region = TrimStringSpaces(value)
	return m
}

func (m *AMIMessage) SetTimezone(value string) *AMIMessage {
	m.Timezone = value
	return m
}

func ofMessage(header textproto.MIMEHeader) *AMIMessage {
	m := &AMIMessage{}
	m.header = header
	return m
}

func ofMessageWithDictionary(d *AMIDictionary, header textproto.MIMEHeader) *AMIMessage {
	m := &AMIMessage{}

	if len(header) > 0 {
		p := make(textproto.MIMEHeader)
		for k, v := range header {
			p.Add(d.TranslateField(k), header.Get(k))
			log.Printf("(AMI). header renew with key = %v and value = %v", k, JsonString(v))
		}
		m.header = p
	} else {
		m.header = header
	}
	return m
}

// Authenticate action by message
func Authenticate(username, password string) *AMIMessage {
	a := WithMessage(config.AmiLoginKey)
	a.AddField(config.AmiFieldUsername, username)
	a.AddField(config.AmiFieldSecret, password)
	return a
}

// Field return first value associated with then given key
func (k *AMIMessage) Field(key string) string {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	return k.header.Get(key)
}

func (k *AMIMessage) FieldOrRefer(key, ref string) string {
	value := k.Field(key)

	if value != "" || len(value) > 0 {
		return value
	}

	return k.Field(ref)
}

func (k *AMIMessage) FieldByDictionary(d *AMIDictionary, key string) string {
	return k.Field(d.TranslateKey(key))
}

func (k *AMIMessage) FieldDictionaryOrRefer(d *AMIDictionary, key, ref string) string {
	value := k.FieldByDictionary(d, key)

	if value != "" && len(value) > 0 && !IsStringEmpty(value) {
		return value
	}

	return k.FieldByDictionary(d, ref)
}

// Return first value associated of field by the given key
func (k *AMIMessage) GetFirstValueByField(key string) string {
	return k.Field(key)
}

func (k *AMIMessage) GetFirstValueByFieldDictionary(d *AMIDictionary, key string) string {
	return k.FieldByDictionary(d, key)
}

// Field Values return all values associated with the given key
func (k *AMIMessage) FieldValues(key string) []string {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	return k.header.Values(key)
}

func (k *AMIMessage) FieldValuesOrRefer(key, ref string) []string {
	values := k.FieldValues(key)
	if len(values) > 0 {
		return values
	}
	return k.FieldValues(ref)
}

func (k *AMIMessage) FieldValuesByDictionary(d *AMIDictionary, key string) []string {
	return k.FieldValues(d.TranslateKey(key))
}

func (k *AMIMessage) FieldValuesByDictionaryOrRefer(d *AMIDictionary, key, ref string) []string {
	values := k.FieldValuesByDictionary(d, key)
	if len(values) > 0 {
		return values
	}
	return k.FieldValuesByDictionary(d, ref)
}

// Return all value associated by the given key
func (k *AMIMessage) GetValuesByField(key string) []string {
	return k.FieldValues(key)
}

func (k *AMIMessage) GetValuesByFieldDictionary(d *AMIDictionary, key string) []string {
	return k.FieldValuesByDictionary(d, key)
}

// Added new pair header as form key:value
func (k *AMIMessage) AddField(key, value string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.header.Add(key, value)
}

// SetField
// Reset value for key associated specified
func (k *AMIMessage) SetField(key, value string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.header.Set(key, value)
}

// SetFields
func (k *AMIMessage) SetFields(fields map[string]string) {
	if len(fields) == 0 {
		return
	}

	for _k, v := range fields {
		k.SetField(_k, v)
	}
}

// Added collection pair header as form key:value
func (k *AMIMessage) AddFields(fields map[string]string) {
	if len(fields) <= 0 {
		return
	}
	for key, value := range fields {
		k.AddField(key, value)
	}
}

// Added new Action-Id and insert into message
func (k *AMIMessage) AddActionIdWith(id string) {
	k.AddField(config.AmiActionIdKey, id)
}

// Added Action-id generated random
func (k *AMIMessage) AddActionId() {
	b := make([]byte, 12)

	_, err := rand.Read(b)

	if err == nil {
		k.AddActionIdWith(fmt.Sprintf("%x", b))
	}
}

// Added Date Received at generated
func (k *AMIMessage) AddFieldDateReceivedAt() {
	if !IsStringEmpty(k.TimeFormat) {
		if IsStringEmpty(k.Timezone) {
			k.AddField(config.AmiFieldDateReceivedAt, time.Now().Format(k.TimeFormat))
		} else {
			k.AddField(config.AmiFieldDateReceivedAt, AdjustTimezone(time.Now(), k.Timezone).Format(k.TimeFormat))
		}
	} else {
		if IsStringEmpty(k.Timezone) {
			k.AddField(config.AmiFieldDateReceivedAt, time.Now().String())
		} else {
			k.AddField(config.AmiFieldDateReceivedAt, AdjustTimezone(time.Now(), k.Timezone).String())
		}
	}
}

// Return AMI message Action Id as string
func (k *AMIMessage) GetActionId() string {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	return k.header.Get(strings.ToLower(config.AmiActionIdKey))
}

// Return AMI message Date Received At as string
func (k *AMIMessage) GetDateReceivedAt() string {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	return k.header.Get(strings.ToLower(config.AmiFieldDateReceivedAt))
}

// Remove fields associated with the given key
func (k *AMIMessage) RemoveField(key string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.header.Del(key)
}

func (k *AMIMessage) RemoveFieldDictionary(d *AMIDictionary, key string) {
	k.RemoveField(d.TranslateKey(key))
}

// Remove fields associated with the collection given keys
func (k *AMIMessage) RemoveFields(keys ...string) {
	for _, key := range keys {
		k.RemoveField(key)
	}
}

func (k *AMIMessage) RemoveFieldsDictionary(d *AMIDictionary, keys ...string) {
	for _, key := range keys {
		k.RemoveFieldDictionary(d, key)
	}
}

// Return new bytes buffer
func (k *AMIMessage) toBytesBuffer() bytes.Buffer {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	var buffer bytes.Buffer

	for key := range k.header {
		for _, value := range k.header.Values(key) {
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
func (k *AMIMessage) String() string {
	_buffer := k.toBytesBuffer()
	return _buffer.String()
}

// Return headers Asterisk Manager Interface (AMI) as byte array
func (k *AMIMessage) Bytes() []byte {
	_buffer := k.toBytesBuffer()
	return _buffer.Bytes()
}

// Return true if the AMI message is event type
func (k *AMIMessage) IsEvent() bool {
	return k.Field(config.AmiEventKey) != ""
}

// Return true if the AMI message is response type
func (k *AMIMessage) IsResponse() bool {
	return k.Field(config.AmiResponseKey) != ""
}

// Return true if the AMI message is response type with value success
func (k *AMIMessage) IsSuccess() bool {
	return strings.EqualFold(k.Field(config.AmiResponseKey), config.AmiStatusSuccessKey)
}

func (k *AMIMessage) PreVars() []string {
	vars := append(k.FieldValues("variable"), k.FieldValues("chanvariable")...)
	vars = append(vars, k.FieldValues("ParkeeChanVariable")...)
	vars = append(vars, k.FieldValues("OrigTransfererChanVariable")...)
	vars = append(vars, k.FieldValues("SecondTransfererChanVariable")...)
	vars = append(vars, k.FieldValues("TransfereeChanVariable")...)
	vars = append(vars, k.FieldValues("TransferTargetChanVariable")...)
	vars = append(vars, k.FieldValues("SpyeeChanVariable")...)
	vars = append(vars, k.FieldValues("SpyerChanVariable")...)

	return vars
}

// Var search in AMI message fields Variable and ChanVariable for a value
// of the type key=value or just key. If found, returns value as string
// and true. Variable name is case sensitive.
func (k *AMIMessage) Var(key string) (string, bool) {
	return k.VarWith(key, k.PreVars())
}

func (k *AMIMessage) VarWith(key string, vars []string) (string, bool) {
	if len(vars) == 0 || key == "" {
		return key, false
	}

	for _, value := range vars {
		e, v := VarsSplit(value)
		if e == key || strings.EqualFold(e, key) {
			return v, true
		}
	}

	return "", false
}

// Return AMI message as interface{}
func (k *AMIMessage) ProduceMessage() map[string]interface{} {
	return k.ProduceMessageWith(false)
}

// Return AMI message as interface{}
func (k *AMIMessage) ProduceMessageTranslator(d *AMIDictionary) map[string]interface{} {
	return k.ProduceMessageWithDictionaries(false, d)
}

// Return AMI message as interface{}
func (k *AMIMessage) ProduceMessageWith(lowercaseField bool) map[string]interface{} {
	translator := NewDictionary()
	return k.ProduceMessageWithDictionaries(lowercaseField, translator)
}

// Return AMI message as interface{}
func (k *AMIMessage) ProduceMessageWithDictionaries(lowercaseField bool, translator *AMIDictionary) map[string]interface{} {
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	data := make(map[string]interface{})

	for key, value := range k.header {
		var field string = key

		if lowercaseField {
			field = strings.ToLower(translator.TranslateField(key))
		} else {
			field = translator.TranslateField(key)
		}

		if len(value) == 1 {
			data[field] = value[0]
		} else {
			data[field] = VarsMap(value)
		}
	}

	return data
}

// Return AMI message as interface{}
func (k *AMIMessage) ProduceMessagePure() map[string]interface{} {
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	data := make(map[string]interface{})

	for key, value := range k.header {
		var field string = key

		if len(value) == 1 {
			data[field] = value[0]
		} else {
			data[field] = VarsMap(value)
		}
	}

	return data
}

// Return AMI message as Json string
func (k *AMIMessage) Json() string {
	return JsonString(k.ProduceMessage())
}

// Return AMI message as Json string
func (k *AMIMessage) JsonTranslator(d *AMIDictionary) string {
	return JsonString(k.ProduceMessageTranslator(d))
}

// Return AMI message as Json pure string
func (k *AMIMessage) JsonPure() string {
	return JsonString(k.ProduceMessagePure())
}

func (k *AMIMessage) apply(e *AMIEvent) *AMIMessage {
	k.SetTimeFormat(e.TimeFormat).
		SetPhonePrefix(e.PhonePrefix).
		SetRegion(e.Region).
		SetTimezone(e.Timezone).
		AddFieldDateReceivedAt()
	return k
}

// Create AMI message from json string
func FromJson(jsonString string) (*AMIMessage, error) {
	var builder interface{}

	message := ofMessage(textproto.MIMEHeader{})

	err := json.Unmarshal([]byte(jsonString), &builder)

	if err != nil {
		log.Printf(err.Error())
		return message, err
	}

	for k, v := range builder.(map[string]interface{}) {
		refType := reflect.ValueOf(v)
		switch refType.Kind() {
		case reflect.Map:
			for v_name, v_val := range v.(map[string]interface{}) {
				message.AddField(k, fmt.Sprintf("%s=%v", v_name, v_val))
			}
		default:
			message.AddField(k, fmt.Sprintf("%v", v))
		}
	}

	return message, nil
}

// MessageSend send an out of call message to an endpoint.
func MessageSend(ctx context.Context, s AMISocket, message AMIPayloadMessage) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMessageSend)
	c.SetVCmd(message)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

func NewAMIPayloadMessage() *AMIPayloadMessage {
	a := &AMIPayloadMessage{}
	return a
}

func (a *AMIPayloadMessage) SetTo(value string) *AMIPayloadMessage {
	a.To = value
	return a
}

func (a *AMIPayloadMessage) SetFrom(value string) *AMIPayloadMessage {
	a.From = value
	return a
}

func (a *AMIPayloadMessage) SetBody(value string) *AMIPayloadMessage {
	a.Body = value
	return a
}

func (a *AMIPayloadMessage) SetBase64Body(value interface{}) *AMIPayloadMessage {
	a.Base64Body = Base64Encode(value)
	return a
}

func (a *AMIPayloadMessage) SetVariable(value string) *AMIPayloadMessage {
	a.Variable = value
	return a
}
