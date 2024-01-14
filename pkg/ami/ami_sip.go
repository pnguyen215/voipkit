package ami

import (
	"context"
	"fmt"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// SIPPeers lists SIP peers in text format with details on current status.
// Peerlist will follow as separate events, followed by a final event called PeerlistComplete
func SIPPeers(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPPeers)
	callback := NewAmiCallbackService(ctx, s, c, []string{config.AmiListenerEventPeerEntry},
		[]string{config.AmiListenerEventPeerlistComplete})
	return callback.SendSuperLevel()
}

// SIPShowPeer shows one SIP peer with details on current status.
func SIPShowPeer(ctx context.Context, s AMISocket, peer string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPShowPeer)
	c.SetV(map[string]string{
		config.AmiFieldPeer: peer,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// SIPPeerStatus show the status of one or all of the sip peers.
func SIPPeerStatus(ctx context.Context, s AMISocket, peer string) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPPeerStatus)
	if peer == "" {
		callback := NewAmiCallbackService(ctx, s, c, []string{config.AmiListenerEventPeerStatus},
			[]string{config.AmiListenerEventSIPpeerstatusComplete})
		return callback.SendSuperLevel()
	}
	c.SetV(map[string]string{
		config.AmiFieldPeer: peer,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{config.AmiListenerEventPeerStatus},
		[]string{config.AmiListenerEventSIPpeerstatusComplete})
	return callback.SendSuperLevel()
}

// SIPPeerStatusShort
func SIPPeerStatusShort(ctx context.Context, s AMISocket, peer string) (AmiReply, error) {
	peers, err := SIPPeerStatus(ctx, s, peer)
	if err != nil {
		return AmiReply{}, err
	}
	if len(peers) == 0 {
		return AmiReply{}, nil
	}
	return peers[0], nil
}

func SIPPeerStatusExists(ctx context.Context, s AMISocket, peer string) (bool, error) {
	sip, err := SIPPeerStatusShort(ctx, s, peer)
	if err != nil {
		return false, err
	}
	if sip.Size() == 0 {
		return false, fmt.Errorf("Peer %v not found", peer)
	}
	return true, nil
}

// SIPShowRegistry shows SIP registrations (text format).
func SIPShowRegistry(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPShowRegistry)
	callback := NewAmiCallbackService(ctx, s, c, []string{config.AmiListenerEventRegistrationEntry},
		[]string{config.AmiListenerEventRegistrationsComplete})
	return callback.SendSuperLevel()
}

// SIPQualifyPeer qualify SIP peers.
func SIPQualifyPeer(ctx context.Context, s AMISocket, peer string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSIPQualifyPeer)
	c.SetV(map[string]string{
		config.AmiFieldPeer: peer,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
