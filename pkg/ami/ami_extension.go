package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// ExtensionStateList list the current known extension states.
func ExtensionStateList(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionExtensionStateList)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventExtensionStatus}, []string{config.AmiListenerEventExtensionStateListComplete})
	return callback.SendSuperLevel()
}

// ExtensionState checks extension status.
func ExtensionState(ctx context.Context, s AMISocket, exten, context string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionExtensionState)
	c.SetV(map[string]string{
		config.AmiFieldExtension: exten,
		config.AmiFieldContext:   context,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
