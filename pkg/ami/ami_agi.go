package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// AGI add an AGI command to execute by Async AGI.
func AGI(ctx context.Context, s AMISocket, channel, agiCommand, agiCommandID string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionAgi)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:   channel,
		config.AmiFieldCommand:   agiCommand,
		config.AmiFieldCommandID: agiCommandID,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ControlPlayback control the playback of a file being played to a channel.
func ControlPlayback(ctx context.Context, s AMISocket, channel string, control config.AGIControl) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionControlPlayback)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel: channel,
		config.AmiFieldControl: control,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
