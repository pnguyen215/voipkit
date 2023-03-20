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

// SIPShowPeer shows one SIP peer with details on current status.
func SIPShowPeer(ctx context.Context, s AMISocket, peer string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPShowPeer)
	c.SetV(map[string]string{
		config.AmiFieldPeer: peer,
	})
	return c.Send(ctx, s, c)
}

// SIPPeerStatus show the status of one or all of the sip peers.
func SIPPeerStatus(ctx context.Context, s AMISocket, peer string) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPPeerStatus)
	if peer == "" {
		return DoGetResult(ctx, s, c, []string{config.AmiListenerEventPeerStatus},
			[]string{config.AmiListenerEventSIPpeerstatusComplete})
	}
	c.SetV(map[string]string{
		config.AmiFieldPeer: peer,
	})
	return DoGetResult(ctx, s, c, []string{config.AmiListenerEventPeerStatus},
		[]string{config.AmiListenerEventSIPpeerstatusComplete})
}

// SIPShowRegistry shows SIP registrations (text format).
func SIPShowRegistry(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPShowRegistry)
	return DoGetResult(ctx, s, c, []string{config.AmiListenerEventRegistrationEntry},
		[]string{config.AmiListenerEventRegistrationsComplete})
}

// SIPQualifyPeer qualify SIP peers.
func SIPQualifyPeer(ctx context.Context, s AMISocket, peer string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPQualifyPeer)
	c.SetV(map[string]string{
		config.AmiFieldPeer: peer,
	})
	return c.Send(ctx, s, c)
}
