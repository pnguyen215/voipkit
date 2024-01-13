package ami

import (
	"context"
	"fmt"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAmiOriginate() *AMIOriginate {
	o := &AMIOriginate{}
	return o
}

func NewAmiDialCall() *AMIDialCall {
	o := &AMIDialCall{}
	o.SetExtensionExists(true)
	return o
}

func (o *AMIOriginate) SetAsync(value bool) *AMIOriginate {
	states := map[bool]string{true: "true", false: "false"}
	o.Async = states[value]
	return o
}

func (o *AMIOriginate) SetChannel(value string) *AMIOriginate {
	o.Channel = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginate) SetExtension(value string) *AMIOriginate {
	o.Exten = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginate) SetContext(value string) *AMIOriginate {
	o.Context = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginate) SetPriority(value int) *AMIOriginate {
	if value >= 0 {
		o.Priority = value
	}
	return o
}

func (o *AMIOriginate) SetApplication(value string) *AMIOriginate {
	o.Application = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginate) SetData(value string) *AMIOriginate {
	o.Data = value
	return o
}

func (o *AMIOriginate) SetDataWith(value interface{}) *AMIOriginate {
	o.SetData(JsonString(value))
	return o
}

func (o *AMIOriginate) SetTimeout(value int) *AMIOriginate {
	if value >= config.AmiMinTimeoutInMsForCall && value <= config.AmiMaxTimeoutInMsForCall {
		o.Timeout = value
	} else {
		o.Timeout = 30000
	}
	return o
}

func (o *AMIOriginate) SetCallerId(value string) *AMIOriginate {
	o.CallerID = value
	return o
}

func (o *AMIOriginate) SetVar(value ...string) *AMIOriginate {
	o.Variable = append(o.Variable, value...)
	return o
}

func (c *AMIOriginate) SetVars(delimiter string, variables ...string) *AMIOriginate {
	if strings.EqualFold(delimiter, "") {
		delimiter = ";"
	}
	vars := strings.Join(variables, delimiter)
	c.SetVar(vars)
	return c
}

func (c *AMIOriginate) SetVarsMap(delimiter string, variables map[string]interface{}) *AMIOriginate {
	if len(variables) <= 0 {
		return c
	}
	_vars := make([]string, len(variables))

	for k, v := range variables {
		str := fmt.Sprintf("%s=%v", k, v)
		_vars = append(_vars, str)
	}

	c.SetVars(delimiter, _vars...)
	return c
}

func (o *AMIOriginate) SetAccount(value string) *AMIOriginate {
	o.Account = value
	return o
}

func (o *AMIOriginate) SetEarlyMedia(value string) *AMIOriginate {
	o.EarlyMedia = value
	return o
}

func (o *AMIOriginate) SetCodecs(value string) *AMIOriginate {
	o.Codecs = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginate) SetChannelId(value string) *AMIOriginate {
	o.ChannelID = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginate) SetOtherChannelId(value string) *AMIOriginate {
	o.OtherChannelID = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginate) Json() string {
	return JsonString(o)
}

func (o *AMIDialCall) SetTelephone(value string) *AMIDialCall {
	o.Telephone = strings.TrimSpace(value)
	return o
}

func (o *AMIDialCall) SetChannelProtocol(value string) *AMIDialCall {
	o.ChannelProtocol = strings.TrimSpace(value)
	return o
}

func (o *AMIDialCall) SetExtension(value int) *AMIDialCall {
	if value >= 0 {
		o.Extension = value
	}
	return o
}

func (o *AMIDialCall) SetDebugMode(value bool) *AMIDialCall {
	o.DebugMode = value
	return o
}

func (o *AMIDialCall) SetTimeout(value int) *AMIDialCall {
	if value >= config.AmiMinTimeoutInMsForCall && value <= config.AmiMaxTimeoutInMsForCall {
		o.Timeout = value
	}
	return o
}

func (o *AMIDialCall) SetExtensionExists(value bool) *AMIDialCall {
	o.ExtensionExists = value
	return o
}

func (o *AMIDialCall) Json() string {
	return JsonString(o)
}

// DialCall
func DialCall(ctx context.Context, s AMISocket, originate AMIOriginate) (AmiReply, error) {
	return Originate(ctx, s, originate)
}

// DialOut
// This is outbound call
// Example:
// action: originate
// channel: SIP/1000
// context: outbound-allroutes
// exten: 012345678
// priority: 1
// timeout: 60000
func DialOut(ctx context.Context, s AMISocket, d AMIDialCall) (AmiReply, bool, error) {
	channel := NewChannel().
		SetChannelProtocol(d.ChannelProtocol)
	o := NewAmiOriginate().
		SetPriority(1).
		SetAsync(true).
		SetTimeout(d.Timeout).
		SetContext(config.AmiContextOutbound).
		SetExtension(strings.TrimSpace(d.Telephone)).
		SetChannel(channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", d.Extension)))

	if d.ExtensionExists {
		peer, err := SIPPeerStatusShort(ctx, s, fmt.Sprintf("%v", d.Extension))
		if err != nil {
			return nil, false, err
		}
		if peer.Size() == 0 {
			return nil, false, fmt.Errorf("Peer %v not found", d.Extension)
		}
		o.SetChannel(peer.Get(config.AmiJsonFieldPeer))
	}
	if d.DebugMode {
		D().Info("DialOut, an outgoing call with originate request body: %v", o.Json())
		D().Info("DialOut, an outgoing call with original request body (setter): %v", d.Json())
	}
	response, err := DialCall(ctx, s, *o)
	return response, IsSuccess(response), err
}

// DialIn
// This is internal call
// Example:
// action: originate
// channel: SIP/1000
// context: from-internal
// exten: 1001
// priority: 1
// timeout: 60000
func DialIn(ctx context.Context, s AMISocket, d AMIDialCall) (AmiReply, bool, error) {
	channel := NewChannel().
		SetChannelProtocol(d.ChannelProtocol)
	o := NewAmiOriginate().
		SetPriority(1).
		SetAsync(true).
		SetTimeout(d.Timeout).
		SetContext(config.AmiContextFromInternal).
		SetExtension(strings.TrimSpace(d.Telephone)).
		SetChannel(channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", d.Extension)))

	if d.ExtensionExists {
		peer, err := SIPPeerStatusShort(ctx, s, fmt.Sprintf("%v", d.Extension))
		if err != nil {
			return nil, false, err
		}
		if peer.Size() == 0 {
			return nil, false, fmt.Errorf("Peer %v not found", d.Extension)
		}
		o.SetChannel(peer.Get(config.AmiJsonFieldPeer))
	}
	if d.DebugMode {
		D().Info("DialIn, an internal call with originate request body: %v", o.Json())
		D().Info("DialIn, an internal call with original request body (setter): %v", d.Json())
	}
	response, err := DialCall(ctx, s, *o)
	return response, IsSuccess(response), err
}
