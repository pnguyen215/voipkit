package ami

import (
	"log"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewEventListener() *AMIEvent {
	e := &AMIEvent{}
	e.SnapChargingEvent()
	e.SetRegion("VN")
	e.SetWrap(make([]error, 0))
	e.AppendWrap(ErrorAsteriskNetwork)
	return e
}

func (m *AMIEvent) SetTimeFormat(value string) *AMIEvent {
	m.TimeFormat = value
	return m
}

func (m *AMIEvent) SetPhonePrefix(value []string) *AMIEvent {
	m.PhonePrefix = value
	return m
}

func (m *AMIEvent) AppendPhonePrefix(values ...string) *AMIEvent {
	m.PhonePrefix = append(m.PhonePrefix, values...)
	return m
}

func (m *AMIEvent) SetRegion(value string) *AMIEvent {
	m.Region = TrimStringSpaces(value)
	return m
}

func (m *AMIEvent) SetTimezone(value string) *AMIEvent {
	m.Timezone = value
	return m
}

func (m *AMIEvent) SetAttempt(value amiRetry) *AMIEvent {
	m.Attempt = value
	return m
}

func (m *AMIEvent) SetPost(value *amiPost) *AMIEvent {
	m.Post = value
	return m
}

func (m *AMIEvent) IsPost() bool {
	return m.Post != nil
}

func (m *AMIEvent) SetWrap(values []error) *AMIEvent {
	m.wrap = values
	return m
}

func (m *AMIEvent) AppendWrap(values ...error) *AMIEvent {
	m.wrap = append(m.wrap, values...)
	return m
}

func (m *AMIEvent) IsWrap() bool {
	return len(m.wrap) > 0 && m.wrap != nil
}

func (m *AMIEvent) IsRetry() bool {
	return m.Attempt.Retry
}

func (m *AMIEvent) Json() string {
	return JsonString(m)
}

func NewAmiRetry() *amiRetry {
	a := &amiRetry{}
	return a
}

func (a *amiRetry) SetRetry(value bool) *amiRetry {
	a.Retry = value
	return a
}

func (a *amiRetry) SetDebugMode(value bool) *amiRetry {
	a.DebugMode = value
	return a
}

func (a *amiRetry) Json() string {
	return JsonString(a)
}

func GetAmiRetrySample() *amiRetry {
	a := NewAmiRetry().
		SetDebugMode(false).
		SetRetry(true)
	return a
}

func NewAmiPost() *amiPost {
	return &amiPost{}
}

func (a *amiPost) SetErr(value error) *amiPost {
	a.err = value
	return a
}

func (a *amiPost) Err() error {
	return a.err
}

func (a *amiPost) WrapErr() string {
	return a.err.Error()
}

func (a *amiPost) IsError() bool {
	return a.err != nil
}

func (m *AMIEvent) Reconnect(ins *AMI, err error) {
	if err == nil {
		return
	}
	if !m.IsWrap() {
		return
	}
	if !m.IsRetry() {
		return
	}
	for _, err := range m.wrap {
		switch e := err.(type) {
		case *AmiError:
			{
				// reconnecting for network error
				if ErrorAsteriskNetwork.S == e.S {
					if m.Attempt.DebugMode {
						D().Error("Ami wrap error occurred: %v", e.Error())
					}
					addr := ins.Conn().RemoteAddr().String()
					current_socket, err := WithSocket(ins.Context(), addr)
					if err != nil {
						if m.Attempt.DebugMode {
							D().Error("Ami wrap error while reconnecting socket connection: %v", err.Error())
						}
						break
					}
					ins.setSocket(current_socket)
					conn, err := WithCore(ins.Context(), ins.Socket(), ins.Auth())
					if err != nil {
						if m.Attempt.DebugMode {
							D().Error("Ami wrap error while creating server reconnection: %v", err.Error())
						}
						break
					}
					// updating stateless
					ins.SetCore(conn)
					ins.release(ins.Context())
					if m.Attempt.DebugMode {
						D().Info("Ami socket reconnected successfully")
					}
				}
			}
		default:
			D().Error("Ami unknown error occurred: %v", err)
			m.SetPost(NewAmiPost().SetErr(err))
			break
		}
	}
}

func (e *AMIEvent) OpenFullEvents(c *AMI) {
	all := c.AllEvents()
	defer c.Close()
	for {
		select {
		case message := <-all:
			message.apply(e)
			log.Printf("ami event: '%s' received: %s", message.Field(strings.ToLower(config.AmiEventKey)), message.Json())
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener has error occurred: %s", err.Error())
		}
	}
}

// Listen All Events with translator dictionary
func (e *AMIEvent) OpenFullEventsTranslator(c *AMI, d *AMIDictionary) {
	all := c.AllEvents()
	defer c.Close()
	for {
		select {
		case message := <-all:
			message.apply(e)
			log.Printf("ami event: '%s' received: %s", message.Field(strings.ToLower(config.AmiEventKey)), message.JsonTranslator(d))
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener has error occurred: %s", err.Error())
		}
	}
}

// Listen All Events Callback with translator dictionary
func (e *AMIEvent) OpenFullEventsCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	all := c.AllEvents()
	defer c.Close()
	for {
		select {
		case message := <-all:
			message.apply(e)
			callback(message, message.JsonTranslator(d), nil)
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener has error occurred: %s", err.Error())
			callback(nil, err.Error(), err)
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
			message.apply(e)
			log.Printf("ami event: '%s' received: %s", name, message.Json())
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener event: '%s' has error occurred: %s", name, err.Error())
		}
	}
}

// Listen Event by key name
func (e *AMIEvent) OpenEventTranslator(c *AMI, d *AMIDictionary, name string) {
	event := c.OnEvent(name)
	defer c.Close()
	for {
		select {
		case message := <-event:
			message.apply(e)
			log.Printf("ami event: '%s' received: %s", name, message.JsonTranslator(d))
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener event: '%s' has error occurred: %s", name, err.Error())
		}
	}
}

// Listen Event Callback key name
func (e *AMIEvent) OpenEventCallbackTranslator(c *AMI, d *AMIDictionary, name string, callback func(*AMIMessage, string, error)) {
	event := c.OnEvent(name)
	defer c.Close()
	for {
		select {
		case message := <-event:
			message.apply(e)
			callback(message, message.JsonTranslator(d), nil)
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener event: '%s' has error occurred: %s", name, err.Error())
			callback(nil, err.Error(), err)
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
			message.apply(e)
			log.Printf("ami event(s): '%s' received: %s", keys, message.Json())
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener event(s): '%s' has error occurred: %s", keys, err.Error())
		}
	}
}

// Listen Events by collect of keys string
func (e *AMIEvent) OpenEventsTranslator(c *AMI, d *AMIDictionary, keys ...string) {
	event := c.OnEvents(keys...)
	defer c.Close()
	for {
		select {
		case message := <-event:
			message.apply(e)
			log.Printf("ami event(s): '%s' received: %s", keys, message.JsonTranslator(d))
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener event(s): '%s' has error occurred: %s", keys, err.Error())
		}
	}
}

// Listen Events Callback by collect of keys string
func (e *AMIEvent) OpenEventsCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error), keys ...string) {
	event := c.OnEvents(keys...)
	defer c.Close()
	for {
		select {
		case message := <-event:
			message.apply(e)
			callback(message, message.JsonTranslator(d), nil)
		case err := <-c.Error():
			c.Close()
			log.Printf("ami listener event(s): '%s' has error occurred: %s", keys, err.Error())
			callback(nil, err.Error(), err)
		}
	}
}

// Listen CDR Event - Call Detail Record, include: link playback (example: .wav, .mp3)
func (e *AMIEvent) OpenCdrEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventCdr)
}

func (e *AMIEvent) OpenCdrEventTranslator(c *AMI, d *AMIDictionary) {
	e.OpenEventTranslator(c, d, config.AmiListenerEventCdr)
}

func (e *AMIEvent) OpenCdrEventCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	e.OpenEventCallbackTranslator(c, d, config.AmiListenerEventCdr, callback)
}

// Listen Bridge Enter Event to mark point connected state
func (e *AMIEvent) OpenBridgeEnterEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventBridgeEnter)
}

func (e *AMIEvent) OpenBridgeEnterEventTranslator(c *AMI, d *AMIDictionary) {
	e.OpenEventTranslator(c, d, config.AmiListenerEventBridgeEnter)
}

func (e *AMIEvent) OpenBridgeEnterEventCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	e.OpenEventCallbackTranslator(c, d, config.AmiListenerEventBridgeEnter, callback)
}

func (e *AMIEvent) OpenConnectedEvent(c *AMI) {
	e.OpenBridgeEnterEvent(c)
}

func (e *AMIEvent) OpenConnectedEventTranslator(c *AMI, d *AMIDictionary) {
	e.OpenBridgeEnterEventTranslator(c, d)
}

func (e *AMIEvent) OpenConnectedEventCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	e.OpenBridgeEnterEventCallbackTranslator(c, d, callback)
}

func (e *AMIEvent) OpenDeviceHangupRequestEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventSoftHangupRequest)
}

func (e *AMIEvent) OpenDeviceHangupRequestEventTranslator(c *AMI, d *AMIDictionary) {
	e.OpenEventTranslator(c, d, config.AmiListenerEventSoftHangupRequest)
}

func (e *AMIEvent) OpenDeviceHangupRequestEventCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	e.OpenEventCallbackTranslator(c, d, config.AmiListenerEventSoftHangupRequest, callback)
}

func (e *AMIEvent) OpenHangupFinishEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventHangup)
}

func (e *AMIEvent) OpenHangupFinishEventTranslator(c *AMI, d *AMIDictionary) {
	e.OpenEventTranslator(c, d, config.AmiListenerEventHangup)
}

func (e *AMIEvent) OpenHangupFinishEventCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	e.OpenEventCallbackTranslator(c, d, config.AmiListenerEventHangup, callback)
}

func (e *AMIEvent) OpenHangupRequestEvent(c *AMI) {
	e.OpenEvent(c, config.AmiListenerEventHangupRequest)
}

func (e *AMIEvent) OpenHangupRequestEventTranslator(c *AMI, d *AMIDictionary) {
	e.OpenEventTranslator(c, d, config.AmiListenerEventHangupRequest)
}

func (e *AMIEvent) OpenHangupRequestEventCallbackTranslator(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	e.OpenEventCallbackTranslator(c, d, config.AmiListenerEventHangupRequest, callback)
}
