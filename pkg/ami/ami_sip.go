package ami

import (
	"context"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

// SIPPeers lists SIP peers in text format with details on current status.
// Peerlist will follow as separate events, followed by a final event called PeerlistComplete
func SIPPeers(ctx context.Context, socket AMISocket, uuid string) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(uuid).SetAction(config.AmiActionSIPPeers)
	return DoGetResult(ctx, socket, c, []string{config.AmiListenerEventPeerEntry},
		[]string{config.AmiListenerEventPeerlistComplete})
}
