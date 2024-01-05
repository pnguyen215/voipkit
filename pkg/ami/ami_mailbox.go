package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// MailboxCount checks Mailbox Message Count.
func MailboxCount(ctx context.Context, s AMISocket, mailbox string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMailboxCount)
	c.SetV(map[string]interface{}{
		config.AmiFieldMailbox: mailbox,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MailboxStatus checks Mailbox Message Count.
func MailboxStatus(ctx context.Context, s AMISocket, mailbox string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMailboxStatus)
	c.SetV(map[string]interface{}{
		config.AmiFieldMailbox: mailbox,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MWIDelete delete selected mailboxes.
func MWIDelete(ctx context.Context, s AMISocket, mailbox string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMWIDelete)
	c.SetV(map[string]interface{}{
		config.AmiFieldMailbox: mailbox,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MWIGet get selected mailboxes with message counts.
func MWIGet(ctx context.Context, s AMISocket, mailbox string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMWIGet)
	c.SetV(map[string]interface{}{
		config.AmiFieldMailbox: mailbox,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MWIUpdate update the mailbox message counts.
func MWIUpdate(ctx context.Context, s AMISocket, mailbox, oldMessages, newMessages string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMWIUpdate)
	c.SetV(map[string]interface{}{
		config.AmiFieldMailbox:     mailbox,
		config.AmiFieldOldMessages: oldMessages,
		config.AmiFieldNewMessages: newMessages,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
