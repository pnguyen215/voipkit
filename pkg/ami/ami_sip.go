package ami

import (
	"context"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

// SIPPeers lists SIP peers in text format with details on current status.
// Peerlist will follow as separate events, followed by a final event called PeerlistComplete
func SIPPeers(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPPeers)
	return DoGetResult(ctx, s, c, []string{config.AmiListenerEventPeerEntry},
		[]string{config.AmiListenerEventPeerlistComplete})
}
