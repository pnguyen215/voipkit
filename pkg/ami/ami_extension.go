package ami

import (
	"context"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

// ExtensionStateList list the current known extension states.
func ExtensionStateList(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionExtensionStateList)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventExtensionStatus}, []string{config.AmiListenerEventExtensionStateListComplete})
	return callback.SendSuperLevel()
}

// ExtensionState checks extension status.
func ExtensionState(ctx context.Context, s AMISocket, exten, context string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionExtensionState)
	c.SetV(map[string]string{
		config.AmiFieldExtension: exten,
		config.AmiFieldContext:   context,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
