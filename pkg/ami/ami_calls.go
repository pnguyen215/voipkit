package ami

import (
	"fmt"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewCall() *AMICall {
	c := &AMICall{}
	return c
}

func NewOriginateCall() *AMIOriginateCall {
	c := &AMIOriginateCall{}
	return c
}

func (c *AMICall) SetNoTarget(phone string) *AMICall {
	c.NoTarget = phone
	return c
}

func (c *AMICall) SetChannelProtocol(protocol string) *AMICall {
	a := NewChannel()
	a.SetChannelProtocol(protocol)
	c.ChannelProtocol = a.ChannelProtocol
	return c
}

func (c *AMICall) SetNoExtension(extension int) *AMICall {
	c.NoExtension = extension
	return c
}

func (c *AMIOriginateCall) SetChannel(channel string) *AMIOriginateCall {
	c.Channel = channel
	return c
}

func (c *AMIOriginateCall) SetContext(context string) *AMIOriginateCall {
	c.Context = context
	return c
}

func (c *AMIOriginateCall) SetExtension(extension string) *AMIOriginateCall {
	c.Extension = extension
	return c
}

func (c *AMIOriginateCall) SetPriority(priority int) *AMIOriginateCall {
	c.Priority = priority
	return c
}

func (c *AMIOriginateCall) SetTimeout(timeout int) *AMIOriginateCall {
	if timeout >= config.AmiMinTimeoutInMsForCall && timeout <= config.AmiMaxTimeoutInMsForCall {
		c.Timeout = timeout
	}
	return c
}

func (c *AMIOriginateCall) SetVar(variable string) *AMIOriginateCall {
	c.Variable = variable
	return c
}

func (c *AMIOriginateCall) SetVars(delimiter string, variables ...string) *AMIOriginateCall {
	if strings.EqualFold(delimiter, "") {
		delimiter = ";"
	}
	vars := strings.Join(variables, delimiter)
	c.SetVar(vars)
	return c
}

func (c *AMIOriginateCall) SetVarsMap(delimiter string, variables map[string]interface{}) *AMIOriginateCall {
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

func (c *AMICall) JsonRequest() string {
	return utils.ToJson(c)
}

func (c *AMIOriginateCall) Call(a *AMI, d *AMIDictionary) (*AMIResponse, error) {
	action := NewActionWith(config.AmiActionOriginate)
	var response *AMIResponse
	var _err error

	if len(c.Channel) <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, config.AmiFieldChannel)
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	if len(c.Context) <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, config.AmiFieldContext)
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	if len(c.Extension) <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, config.AmiFieldExtension)
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	action.AddField(config.AmiFieldAsync, "true")
	action.AddField(config.AmiFieldVariable, c.Variable)
	action.AddField(config.AmiFieldChannel, c.Channel)
	action.AddField(config.AmiFieldContext, c.Context)
	action.AddField(config.AmiFieldExtension, c.Extension)

	if c.Priority >= 1 {
		action.AddField(config.AmiFieldPriority, fmt.Sprintf("%v", c.Priority))
	}

	if c.Timeout >= config.AmiMinTimeoutInMsForCall && c.Timeout <= config.AmiMaxTimeoutInMsForCall {
		action.AddField(config.AmiFieldTimeout, fmt.Sprintf("%v", c.Timeout))
	}

	response, _err = NewRevokeAction(config.AmiActionOriginate, c.Timeout).
		RevokeAction(a, d, action, false)
	return response, _err
}

// OriginateExternalCall
// This is outbound call
// Example:
// action: originate
// channel: SIP/1000
// context: outbound-allroutes
// exten: 012345678
// priority: 1
// timeout: 60000
func (c *AMICall) OriginateExternalCall(a *AMI, d *AMIDictionary) (*AMIResponse, error) {
	var response *AMIResponse
	var _err error

	if len(c.NoTarget) <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "No. target")
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	if len(c.ChannelProtocol) <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "Channel protocol")
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	if c.NoExtension <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "No. extension")
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	channel := NewChannel().SetChannelProtocol(c.ChannelProtocol)
	c.SetChannel(channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", c.NoExtension)))
	c.SetContext(config.AmiContextOutbound)
	c.SetExtension(strings.TrimSpace(c.NoTarget))
	c.SetPriority(1)
	return c.Call(a, d)
}

// OriginateInternalCall
// This is internal call
// Example:
// action: originate
// channel: SIP/1000
// context: from-internal
// exten: 1001
// priority: 1
// timeout: 60000
func (c *AMICall) OriginateInternalCall(a *AMI, d *AMIDictionary) (*AMIResponse, error) {
	var response *AMIResponse
	var _err error

	if len(c.NoTarget) <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "No. target")
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	if len(c.ChannelProtocol) <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "Channel protocol")
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	if c.NoExtension <= 0 {
		response.IsSuccess = false
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "No. extension")
		_err = fmt.Errorf(response.ErrorMessage)
		return response, _err
	}

	channel := NewChannel().SetChannelProtocol(c.ChannelProtocol)
	c.SetChannel(channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", c.NoExtension)))
	c.SetContext(config.AmiContextFromInternal)
	c.SetExtension(strings.TrimSpace(c.NoTarget))
	c.SetPriority(1)
	return c.Call(a, d)
}
