package ami

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewAMIPayloadOriginate() *AMIPayloadOriginate {
	o := &AMIPayloadOriginate{}
	return o
}

func NewAMIOriginateDirection() *AMIOriginateDirection {
	o := &AMIOriginateDirection{}
	o.SetAllowSysValidator(true)
	return o
}

func (o *AMIPayloadOriginate) SetAsync(value bool) *AMIPayloadOriginate {
	states := map[bool]string{true: "true", false: "false"}
	o.Async = states[value]
	return o
}

func (o *AMIPayloadOriginate) SetChannel(value string) *AMIPayloadOriginate {
	o.Channel = strings.TrimSpace(value)
	return o
}

func (o *AMIPayloadOriginate) SetExtension(value string) *AMIPayloadOriginate {
	o.Exten = strings.TrimSpace(value)
	return o
}

func (o *AMIPayloadOriginate) SetContext(value string) *AMIPayloadOriginate {
	o.Context = strings.TrimSpace(value)
	return o
}

func (o *AMIPayloadOriginate) SetPriority(value int) *AMIPayloadOriginate {
	if value >= 0 {
		o.Priority = value
	}
	return o
}

func (o *AMIPayloadOriginate) SetApplication(value string) *AMIPayloadOriginate {
	o.Application = strings.TrimSpace(value)
	return o
}

func (o *AMIPayloadOriginate) SetData(value string) *AMIPayloadOriginate {
	o.Data = value
	return o
}

func (o *AMIPayloadOriginate) SetDataWith(value interface{}) *AMIPayloadOriginate {
	o.SetData(utils.ToJson(value))
	return o
}

func (o *AMIPayloadOriginate) SetTimeout(value int) *AMIPayloadOriginate {
	if value >= config.AmiMinTimeoutInMsForCall && value <= config.AmiMaxTimeoutInMsForCall {
		o.Timeout = value
	}
	return o
}

func (o *AMIPayloadOriginate) SetCallerId(value string) *AMIPayloadOriginate {
	o.CallerID = value
	return o
}

func (o *AMIPayloadOriginate) SetVar(value ...string) *AMIPayloadOriginate {
	o.Variable = append(o.Variable, value...)
	return o
}

func (c *AMIPayloadOriginate) SetVars(delimiter string, variables ...string) *AMIPayloadOriginate {
	if strings.EqualFold(delimiter, "") {
		delimiter = ";"
	}
	vars := strings.Join(variables, delimiter)
	c.SetVar(vars)
	return c
}

func (c *AMIPayloadOriginate) SetVarsMap(delimiter string, variables map[string]interface{}) *AMIPayloadOriginate {
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

func (o *AMIPayloadOriginate) SetAccount(value string) *AMIPayloadOriginate {
	o.Account = value
	return o
}

func (o *AMIPayloadOriginate) SetEarlyMedia(value string) *AMIPayloadOriginate {
	o.EarlyMedia = value
	return o
}

func (o *AMIPayloadOriginate) SetCodecs(value string) *AMIPayloadOriginate {
	o.Codecs = strings.TrimSpace(value)
	return o
}

func (o *AMIPayloadOriginate) SetChannelId(value string) *AMIPayloadOriginate {
	o.ChannelID = strings.TrimSpace(value)
	return o
}

func (o *AMIPayloadOriginate) SetOtherChannelId(value string) *AMIPayloadOriginate {
	o.OtherChannelID = strings.TrimSpace(value)
	return o
}

func (o *AMIPayloadOriginate) Json() string {
	return utils.ToJson(o)
}

func (o *AMIOriginateDirection) SetTelephone(value string) *AMIOriginateDirection {
	o.Telephone = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginateDirection) SetChannelProtocol(value string) *AMIOriginateDirection {
	o.ChannelProtocol = strings.TrimSpace(value)
	return o
}

func (o *AMIOriginateDirection) SetExtension(value int) *AMIOriginateDirection {
	if value >= 0 {
		o.Extension = value
	}
	return o
}

func (o *AMIOriginateDirection) SetAllowDebug(value bool) *AMIOriginateDirection {
	o.AllowDebug = value
	return o
}

func (o *AMIOriginateDirection) SetTimeout(value int) *AMIOriginateDirection {
	if value >= config.AmiMinTimeoutInMsForCall && value <= config.AmiMaxTimeoutInMsForCall {
		o.Timeout = value
	}
	return o
}

func (o *AMIOriginateDirection) SetAllowSysValidator(value bool) *AMIOriginateDirection {
	o.AllowSysValidator = value
	return o
}

func (o *AMIOriginateDirection) Json() string {
	return utils.ToJson(o)
}

// MakeCall
func MakeCall(ctx context.Context, s AMISocket, originate AMIPayloadOriginate) (AMIResultRaw, error) {
	return Originate(ctx, s, originate)
}

// MakeOutboundCall
// This is outbound call
// Example:
// action: originate
// channel: SIP/1000
// context: outbound-allroutes
// exten: 012345678
// priority: 1
// timeout: 60000
func MakeOutboundCall(ctx context.Context, s AMISocket, d AMIOriginateDirection) (AMIResultRaw, bool, error) {
	channel := NewChannel().SetChannelProtocol(d.ChannelProtocol)
	o := NewAMIPayloadOriginate().
		SetAsync(true).
		SetContext(config.AmiContextOutbound).
		SetExtension(strings.TrimSpace(d.Telephone)).
		SetPriority(1).
		SetChannel(channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", d.Extension)))

	if d.Timeout >= config.AmiMinTimeoutInMsForCall && d.Timeout <= config.AmiMaxTimeoutInMsForCall {
		o.SetTimeout(d.Timeout)
	} else {
		o.SetTimeout(30000) // as default
	}

	if d.AllowSysValidator {
		peer, err := SIPPeerStatusShort(ctx, s, fmt.Sprintf("%v", d.Extension))
		if err != nil {
			return nil, false, err
		}
		if peer.LenValue() == 0 {
			return nil, false, fmt.Errorf("Peer %v not found", d.Extension)
		}
		o.SetChannel(peer.GetVal(config.AmiJsonFieldPeer))
	}

	if d.AllowDebug {
		log.Printf("MakeOutboundCall, an outgoing call with originate request body = %v", o.Json())
		log.Printf("MakeOutboundCall, an outgoing call with original request body (setter) = %v", d.Json())
	}

	response, err := MakeCall(ctx, s, *o)
	return response, IsSuccess(response), err
}

// MakeInternalCall
// This is internal call
// Example:
// action: originate
// channel: SIP/1000
// context: from-internal
// exten: 1001
// priority: 1
// timeout: 60000
func MakeInternalCall(ctx context.Context, s AMISocket, d AMIOriginateDirection) (AMIResultRaw, bool, error) {
	channel := NewChannel().SetChannelProtocol(d.ChannelProtocol)
	o := NewAMIPayloadOriginate().
		SetAsync(true).
		SetContext(config.AmiContextFromInternal).
		SetExtension(strings.TrimSpace(d.Telephone)).
		SetPriority(1).
		SetChannel(channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", d.Extension)))

	if d.Timeout >= config.AmiMinTimeoutInMsForCall && d.Timeout <= config.AmiMaxTimeoutInMsForCall {
		o.SetTimeout(d.Timeout)
	} else {
		o.SetTimeout(30000) // as default
	}

	if d.AllowSysValidator {
		peer, err := SIPPeerStatusShort(ctx, s, fmt.Sprintf("%v", d.Extension))
		if err != nil {
			return nil, false, err
		}
		if peer.LenValue() == 0 {
			return nil, false, fmt.Errorf("Peer %v not found", d.Extension)
		}
		o.SetChannel(peer.GetVal(config.AmiJsonFieldPeer))
	}

	if d.AllowDebug {
		log.Printf("MakeInternalCall, an internal call with originate request body = %v", o.Json())
		log.Printf("MakeInternalCall, an internal call with original request body (setter) = %v", d.Json())
	}

	response, err := MakeCall(ctx, s, *o)
	return response, IsSuccess(response), err
}
