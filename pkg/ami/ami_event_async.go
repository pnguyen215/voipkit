package ami

func (e *AMIEvent) OpenFullEventsAsyncFunc(c *AMI) {
	go func() {
		e.OpenFullEvents(c)
	}()
}

func (e *AMIEvent) OpenFullEventsTranslatorAsyncFunc(c *AMI, d *AMIDictionary) {
	go func() {
		e.OpenFullEventsTranslator(c, d)
	}()
}

func (e *AMIEvent) OpenFullEventsCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenFullEventsCallbackTranslator(c, d, callback)
	}()
}

func (e *AMIEvent) OpenEventAsyncFunc(c *AMI, name string) {
	go func() {
		e.OpenEvent(c, name)
	}()
}

func (e *AMIEvent) OpenEventTranslatorAsyncFunc(c *AMI, d *AMIDictionary, name string) {
	go func() {
		e.OpenEventTranslator(c, d, name)
	}()
}

func (e *AMIEvent) OpenEventCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, name string, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenEventCallbackTranslator(c, d, name, callback)
	}()
}

func (e *AMIEvent) OpenEventsAsyncFunc(c *AMI, keys ...string) {
	go func() {
		e.OpenEvents(c, keys...)
	}()
}

func (e *AMIEvent) OpenEventsTranslatorAsyncFunc(c *AMI, d *AMIDictionary, keys ...string) {
	go func() {
		e.OpenEventsTranslator(c, d, keys...)
	}()
}

func (e *AMIEvent) OpenEventsCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error), keys ...string) {
	go func() {
		e.OpenEventsCallbackTranslator(c, d, callback, keys...)
	}()
}

func (e *AMIEvent) OpenCdrEventAsyncFunc(c *AMI) {
	go func() {
		e.OpenCdrEvent(c)
	}()
}

func (e *AMIEvent) OpenCdrEventTranslatorAsyncFunc(c *AMI, d *AMIDictionary) {
	go func() {
		e.OpenCdrEventTranslator(c, d)
	}()
}

func (e *AMIEvent) OpenCdrEventCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenCdrEventCallbackTranslator(c, d, callback)
	}()
}

func (e *AMIEvent) OpenBridgeEnterEventAsyncFunc(c *AMI) {
	go func() {
		e.OpenBridgeEnterEvent(c)
	}()
}

func (e *AMIEvent) OpenBridgeEnterEventTranslatorAsyncFunc(c *AMI, d *AMIDictionary) {
	go func() {
		e.OpenBridgeEnterEventTranslator(c, d)
	}()
}

func (e *AMIEvent) OpenBridgeEnterEventCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenBridgeEnterEventCallbackTranslator(c, d, callback)
	}()
}

func (e *AMIEvent) OpenConnectedEventAsyncFunc(c *AMI) {
	go func() {
		e.OpenConnectedEvent(c)
	}()
}

func (e *AMIEvent) OpenConnectedEventTranslatorAsyncFunc(c *AMI, d *AMIDictionary) {
	go func() {
		e.OpenConnectedEventTranslator(c, d)
	}()
}

func (e *AMIEvent) OpenConnectedEventCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenConnectedEventCallbackTranslator(c, d, callback)
	}()
}

func (e *AMIEvent) OpenDeviceHangupRequestEventAsyncFunc(c *AMI) {
	go func() {
		e.OpenDeviceHangupRequestEvent(c)
	}()
}

func (e *AMIEvent) OpenDeviceHangupRequestEventTranslatorAsyncFunc(c *AMI, d *AMIDictionary) {
	go func() {
		e.OpenDeviceHangupRequestEventTranslator(c, d)
	}()
}

func (e *AMIEvent) OpenDeviceHangupRequestEventCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenDeviceHangupRequestEventCallbackTranslator(c, d, callback)
	}()
}

func (e *AMIEvent) OpenHangupFinishEventAsyncFunc(c *AMI) {
	go func() {
		e.OpenHangupFinishEvent(c)
	}()
}

func (e *AMIEvent) OpenHangupFinishEventTranslatorAsyncFunc(c *AMI, d *AMIDictionary) {
	go func() {
		e.OpenHangupFinishEventTranslator(c, d)
	}()
}

func (e *AMIEvent) OpenHangupFinishEventCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenHangupFinishEventCallbackTranslator(c, d, callback)
	}()
}

func (e *AMIEvent) OpenHangupRequestEventAsyncFunc(c *AMI) {
	go func() {
		e.OpenHangupRequestEvent(c)
	}()
}

func (e *AMIEvent) OpenHangupRequestEventTranslatorAsyncFunc(c *AMI, d *AMIDictionary) {
	go func() {
		e.OpenHangupRequestEventTranslator(c, d)
	}()
}

func (e *AMIEvent) OpenHangupRequestEventCallbackTranslatorAsyncFunc(c *AMI, d *AMIDictionary, callback func(*AMIMessage, string, error)) {
	go func() {
		e.OpenHangupRequestEventCallbackTranslator(c, d, callback)
	}()
}
