package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// PresenceState check presence state.
func PresenceState(ctx context.Context, s AMISocket, provider string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPresenceState)
	c.SetV(map[string]interface{}{
		config.AmiFieldProvider: provider,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PresenceStateList list the current known presence states.
func PresenceStateList(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPresenceStateList)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventAgents}, []string{config.AmiListenerEventPresenceStateListComplete})
	return callback.SendSuperLevel()
}
