package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// IAXnetstats show IAX channels network statistics.
func IAXnetstats(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionIAXnetstats)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventPeerEntry}, []string{config.AmiListenerEventPeerlistComplete})
	return callback.SendSuperLevel()
}

// IAXpeerlist show IAX channels network statistics.
func IAXpeerlist(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionIAXpeerlist)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventPeerEntry}, []string{config.AmiListenerEventPeerlistComplete})
	return callback.SendSuperLevel()
}

// IAXpeers list IAX peers.
func IAXpeers(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionIAXpeers)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventPeerEntry}, []string{config.AmiListenerEventPeerlistComplete})
	return callback.SendSuperLevel()
}

// IAXregistry show IAX registrations.
func IAXregistry(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionIAXregistry)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventPeerEntry}, []string{config.AmiListenerEventPeerlistComplete})
	return callback.SendSuperLevel()
}
