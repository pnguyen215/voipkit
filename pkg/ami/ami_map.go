package ami

import (
	"context"
	"strconv"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func (c *AMICore) ExtensionStatesMap(ctx context.Context, guard *AMIExtensionGuard) (response []AMIExtensionStatus, err error) {
	peers, err := c.ExtensionStates(ctx)
	if err != nil {
		return response, err
	}

	if len(peers) == 0 {
		return response, nil
	}

	for _, v := range peers {
		if guard.AllowExtensionNumeric {
			_, err := strconv.Atoi(v.Get(config.AmiJsonFieldExten))
			if err != nil {
				continue
			}
		}

		if len(guard.Context) > 0 {
			if !Contains(guard.Context, v.Get(config.AmiJsonFieldContext)) {
				continue
			}
		}

		if len(guard.StatusesText) > 0 {
			if !Contains(guard.StatusesText, v.Get(config.AmiJsonFieldStatusText)) {
				continue
			}
		}

		e := c.convRaw2ExtensionStatus(v)
		response = append(response, *e)
	}

	return response, nil
}

func (c *AMICore) ExtensionStateMap(ctx context.Context, exten, context string) (response AMIExtensionStatus, err error) {
	v, err := c.ExtensionState(ctx, exten, context)
	if err != nil {
		return response, err
	}
	if len(v) == 0 {
		return response, nil
	}
	e := c.convRaw2ExtensionStatus(v)
	response = *e
	return response, nil
}

func (c *AMICore) GetSIPPeersStatusMap(ctx context.Context, g *AMIPeerStatusGuard) (response []AMIPeerStatus, err error) {
	peers, err := c.GetSIPPeersStatus(ctx)
	if err != nil {
		return response, err
	}
	if len(peers) == 0 {
		return response, nil
	}
	for _, v := range peers {
		e := c.convRaw2PeerStatus(v, g)
		response = append(response, *e)
	}
	return response, nil
}

func (c *AMICore) GetSIPPeerStatusMap(ctx context.Context, g *AMIPeerStatusGuard, peer string) (response AMIPeerStatus, err error) {
	peers, err := c.GetSIPPeerStatus(ctx, peer)
	if err != nil {
		return response, err
	}
	if len(peers) == 0 {
		return response, nil
	}
	e := c.convRaw2PeerStatus(peers, g)
	response = *e
	return response, nil
}

func (c *AMICore) convRaw2ExtensionStatus(v AmiReply) *AMIExtensionStatus {
	e := NewAMIExtensionStatus().
		SetActionId(v.Get(config.AmiJsonFieldActionId)).
		SetContext(v.Get(config.AmiJsonFieldContext)).
		SetExtension(v.Get(config.AmiJsonFieldExten)).
		SetHint(v.Get(config.AmiJsonFieldHint)).
		SetMessage(v.Get(config.AmiJsonFieldMessage)).
		SetResponse(v.Get(config.AmiJsonFieldResponse)).
		SetStatus(v.Get(config.AmiJsonFieldStatus)).
		SetStatusText(v.Get(config.AmiJsonFieldStatusText))
	return e
}

func (c *AMICore) convRaw2PeerStatus(v AmiReply, g *AMIPeerStatusGuard) *AMIPeerStatus {
	e := NewAMIPeerStatus().
		SetActionId(v.Get(config.AmiJsonFieldActionId)).
		SetChannelType(v.Get(config.AmiJsonFieldChannelType)).
		SetEvent(v.Get(config.AmiJsonFieldEvent)).
		SetPeer(v.Get(config.AmiJsonFieldPeer)).
		SetPeerStatus(v.Get(config.AmiJsonFieldPeerStatus)).
		SetPrivilege(v.Get(config.AmiJsonFieldPrivilege)).
		SetTimeInMs(v.Get(config.AmiJsonFieldTime)).
		SetPublishedAt(AdjustTimezone(time.Now(), g.Timezone))

	if e.TimeInMs > 0 {
		e.SetPublishedAt(e.PublishedAt.Add(-time.Millisecond * time.Duration(e.TimeInMs)))
	}

	if IsStringEmpty(g.DateTimeLayout) {
		e.SetPrePublishedAt(e.PublishedAt.Format(config.DateTimeFormat20060102150405))
	} else {
		e.SetPrePublishedAt(e.PublishedAt.Format(g.DateTimeLayout))
	}

	return e
}
