package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// FAXSession responds with a detailed description of a single FAX session.
func FAXSession(ctx context.Context, s AMISocket, sessionNumber string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionFAXSession)
	c.SetV(map[string]interface{}{
		config.AmiFieldSessionNumber: sessionNumber,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// FAXSessions list active FAX sessions.
func FAXSessions(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionFAXSessions)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventFAXSessionsEntry}, []string{config.AmiListenerEventFAXSessionsComplete})
	return callback.SendSuperLevel()
}

// FAXStats responds with fax statistics.
func FAXStats(ctx context.Context, s AMISocket) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionFAXStats)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	callback.Send() // preprocessing
	return callback.Send()
}
