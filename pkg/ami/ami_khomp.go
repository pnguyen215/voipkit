package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAMIPayloadKhompSMS() *AMIPayloadKhompSMS {
	a := &AMIPayloadKhompSMS{}
	return a
}

func (a *AMIPayloadKhompSMS) SetDevice(value string) *AMIPayloadKhompSMS {
	a.Device = value
	return a
}

func (a *AMIPayloadKhompSMS) SetDestination(value string) *AMIPayloadKhompSMS {
	a.Destination = value
	return a
}

func (a *AMIPayloadKhompSMS) SetConfirmation(value bool) *AMIPayloadKhompSMS {
	a.Confirmation = value
	return a
}

func (a *AMIPayloadKhompSMS) SetMessage(value string) *AMIPayloadKhompSMS {
	a.Message = value
	return a
}

// KSendSMS sends a SMS using KHOMP device.
func KSendSMS(ctx context.Context, s AMISocket, payload AMIPayloadKhompSMS) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionKSendSMS)
	c.SetVCmd(payload)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
