package ami

import (
	"context"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

// DAHDIDialOffhook dials over DAHDI channel while offhook.
// Generate DTMF control frames to the bridged peer.
func DAHDIDialOffhook(ctx context.Context, s AMISocket, channel, number string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDAHDIDialOffhook)
	c.SetV(map[string]interface{}{
		config.AmiFieldDAHDIChannel: channel,
		config.AmiFieldNumber:       number,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DAHDIDNDoff toggles DAHDI channel Do Not Disturb status OFF.
func DAHDIDNDoff(ctx context.Context, s AMISocket, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDAHDIDNDoff)
	c.SetV(map[string]interface{}{
		config.AmiFieldDAHDIChannel: channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DAHDIDNDon toggles DAHDI channel Do Not Disturb status ON.
func DAHDIDNDon(ctx context.Context, s AMISocket, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDAHDIDNDon)
	c.SetV(map[string]interface{}{
		config.AmiFieldDAHDIChannel: channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DAHDIHangup hangups DAHDI Channel.
func DAHDIHangup(ctx context.Context, s AMISocket, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDAHDIHangup)
	c.SetV(map[string]interface{}{
		config.AmiFieldDAHDIChannel: channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DAHDIRestart fully Restart DAHDI channels (terminates calls).
func DAHDIRestart(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDAHDIRestart)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DAHDIShowChannels show status of DAHDI channels.
func DAHDIShowChannels(ctx context.Context, s AMISocket, channel string) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDAHDIShowChannels)
	c.SetV(map[string]interface{}{
		config.AmiFieldDAHDIChannel: channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{config.AmiListenerEventDAHDIShowChannels}, []string{config.AmiListenerEventDAHDIShowChannelsComplete})
	return callback.SendSuperLevel()
}

// DAHDITransfer transfers DAHDI Channel.
func DAHDITransfer(ctx context.Context, s AMISocket, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDAHDITransfer)
	c.SetV(map[string]interface{}{
		config.AmiFieldDAHDIChannel: channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
