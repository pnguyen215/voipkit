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
			_, err := strconv.Atoi(v.GetVal(config.AmiJsonFieldExten))
			if err != nil {
				continue
			}
		}

		if len(guard.Context) > 0 {
			if !Contains(guard.Context, v.GetVal(config.AmiJsonFieldContext)) {
				continue
			}
		}

		if len(guard.StatusesText) > 0 {
			if !Contains(guard.StatusesText, v.GetVal(config.AmiJsonFieldStatusText)) {
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

func (c *AMICore) convRaw2ExtensionStatus(v AMIResultRaw) *AMIExtensionStatus {
	e := NewAMIExtensionStatus().
		SetActionId(v.GetVal(config.AmiJsonFieldActionId)).
		SetContext(v.GetVal(config.AmiJsonFieldContext)).
		SetExtension(v.GetVal(config.AmiJsonFieldExten)).
		SetHint(v.GetVal(config.AmiJsonFieldHint)).
		SetMessage(v.GetVal(config.AmiJsonFieldMessage)).
		SetResponse(v.GetVal(config.AmiJsonFieldResponse)).
		SetStatus(v.GetVal(config.AmiJsonFieldStatus)).
		SetStatusText(v.GetVal(config.AmiJsonFieldStatusText))
	return e
}

func (c *AMICore) convRaw2PeerStatus(v AMIResultRaw, g *AMIPeerStatusGuard) *AMIPeerStatus {
	e := NewAMIPeerStatus().
		SetActionId(v.GetVal(config.AmiJsonFieldActionId)).
		SetChannelType(v.GetVal(config.AmiJsonFieldChannelType)).
		SetEvent(v.GetVal(config.AmiJsonFieldEvent)).
		SetPeer(v.GetVal(config.AmiJsonFieldPeer)).
		SetPeerStatus(v.GetVal(config.AmiJsonFieldPeerStatus)).
		SetPrivilege(v.GetVal(config.AmiJsonFieldPrivilege)).
		SetTimeInMs(v.GetVal(config.AmiJsonFieldTime)).
		SetPublishedAt(AdjustTimezone(time.Now(), g.Timezone))

	if e.TimeInMs > 0 {
		e.SetPublishedAt(e.PublishedAt.Add(-time.Millisecond * time.Duration(e.TimeInMs)))
	}

	if IsStringEmpty(g.DateTimeLayout) {
		e.SetPrePublishedAt(e.PublishedAt.Format(config.DateTimeFormatYYYYMMDDHHMMSS))
	} else {
		e.SetPrePublishedAt(e.PublishedAt.Format(g.DateTimeLayout))
	}

	return e
}
