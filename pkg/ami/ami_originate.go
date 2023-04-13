package ami

import (
	"fmt"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewAMIPayloadOriginate() *AMIPayloadOriginate {
	o := &AMIPayloadOriginate{}
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
	if value >= 0 {
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
