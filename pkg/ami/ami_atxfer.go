package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// Atxfer attended transfer.
func Atxfer(ctx context.Context, s AMISocket, channel, extension, context string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionAtxfer)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:   channel,
		config.AmiFieldExtension: extension,
		config.AmiFieldContext:   context,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// CancelAtxfer cancel an attended transfer.
func CancelAtxfer(ctx context.Context, s AMISocket, channel string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCancelAtxfer)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel: channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
