package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAMIPayloadExtension() *AMIPayloadExtension {
	e := &AMIPayloadExtension{}
	return e
}

func (e *AMIPayloadExtension) SetContext(value string) *AMIPayloadExtension {
	e.Context = value
	return e
}

func (e *AMIPayloadExtension) SetExtension(value string) *AMIPayloadExtension {
	e.Extension = value
	return e
}

func (e *AMIPayloadExtension) SetPriority(value string) *AMIPayloadExtension {
	e.Priority = value
	return e
}

func (e *AMIPayloadExtension) SetApplication(value string) *AMIPayloadExtension {
	e.Application = value
	return e
}

func (e *AMIPayloadExtension) SetApplicationData(value string) *AMIPayloadExtension {
	e.ApplicationData = value
	return e
}

func (e *AMIPayloadExtension) SetReplace(value string) *AMIPayloadExtension {
	e.Replace = value
	return e
}

func (e *AMIPayloadExtension) SetApplicationDataWith(v interface{}) *AMIPayloadExtension {
	e.SetApplicationData(JsonString(v))
	return e
}

// AddDialplanExtension add an extension to the dialplan.
func AddDialplanExtension(ctx context.Context, s AMISocket, extension AMIPayloadExtension) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDialplanExtensionAdd)
	c.SetVCmd(extension)
	_s := s
	_s.SetRetry(false)
	callback := NewAMICallbackService(ctx, _s, c, []string{}, []string{})
	return callback.Send()
}

// RemoveDialplanExtension remove an extension from the dialplan.
func RemoveDialplanExtension(ctx context.Context, s AMISocket, extension AMIPayloadExtension) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDialplanExtensionRemove)
	c.SetVCmd(extension)
	_s := s
	_s.SetRetry(false)
	callback := NewAMICallbackService(ctx, _s, c, []string{}, []string{})
	return callback.Send()
}
