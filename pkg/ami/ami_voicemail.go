package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// VoicemailRefresh tell asterisk to poll mailboxes for a change.
func VoicemailRefresh(ctx context.Context, s AMISocket, context, mailbox string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionVoicemailRefresh)
	c.SetV(map[string]interface{}{
		config.AmiFieldMailbox: mailbox,
		config.AmiFieldContext: context,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// VoicemailUsersList list all voicemail user information.
func VoicemailUsersList(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionVoicemailUsersList)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventVoicemailUserEntry}, []string{config.AmiListenerEventVoicemailUserEntryComplete})
	return callback.SendSuperLevel()
}
