package ami

import (
	"strconv"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

var (
	AmiExtensionStatesString []AMIExtensionStatesConf = []AMIExtensionStatesConf{
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionNotInUse).SetText("Idle"),
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionInUse).SetText("InUse"),
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionBusy).SetText("Busy"),
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionUnavailable).SetText("Unavailable"),
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionRinging).SetText("Ringing"),
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionInUse | config.AmiExtensionRinging).SetText("InUse&Ringing"),
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionOnHold).SetText("Hold"),
		*NewAMIExtensionStatesConf().SetExtensionState(config.AmiExtensionInUse | config.AmiExtensionOnHold).SetText("InUse&Hold"),
	}
)

func NewAMIExtensionStatesConf() *AMIExtensionStatesConf {
	e := &AMIExtensionStatesConf{}
	return e
}

func NewAMIConf() *AMIConf {
	e := &AMIConf{}
	return e
}

func NewAMIExtensionStatus() *AMIExtensionStatus {
	e := &AMIExtensionStatus{}
	return e
}

func NewAMIExtensionGuard() *AMIExtensionGuard {
	e := &AMIExtensionGuard{}
	return e
}

func NewAMIPeerStatus() *AMIPeerStatus {
	e := &AMIPeerStatus{}
	return e
}

func NewAMIPeerStatusGuard() *AMIPeerStatusGuard {
	e := &AMIPeerStatusGuard{}
	e.SetDateTimeLayout(config.DateTimeFormat20060102150405)
	e.SetTimezone(config.DefaultTimezoneVietnam)
	return e
}

func (e *AMIExtensionStatesConf) SetExtensionState(value int) *AMIExtensionStatesConf {
	e.ExtensionState = value
	return e
}

func (e *AMIExtensionStatesConf) SetText(value string) *AMIExtensionStatesConf {
	e.Text = value
	return e
}

func (e *AMIConf) ConvDeviceStateToExtensionState(deviceState int) int {
	switch deviceState {
	case config.AmiDeviceStateOnHold:
		return config.AmiExtensionOnHold
	case config.AmiDeviceStateBusy:
		return config.AmiExtensionBusy
	case config.AmiDeviceStateUnknown:
		return config.AmiExtensionNotInUse
	case config.AmiDeviceStateUnavailable, config.AmiDeviceStateInvalid:
		return config.AmiExtensionUnavailable
	case config.AmiDeviceStateRingInUse:
		return (config.AmiExtensionInUse | config.AmiExtensionRinging)
	case config.AmiDeviceStateRinging:
		return config.AmiExtensionRinging
	case config.AmiDeviceStateInUse:
		return config.AmiExtensionInUse
	case config.AmiDeviceStateNotInUse:
		return config.AmiExtensionNotInUse
	case config.AmiDeviceStateTotal:
		break
	}
	return config.AmiExtensionNotInUse
}

func (e *AMIConf) ConvChannelStateToDeviceState(channelState int) int {
	switch channelState {
	case config.AmiChannelStateDown:
		return config.AmiDeviceStateNotInUse
	case config.AmiChannelStateReserved,
		config.AmiChannelStateOffHook,
		config.AmiChannelStateDialing,
		config.AmiChannelStateRing,
		config.AmiChannelStateUp,
		config.AmiChannelStateDialingOffHook:
		return config.AmiDeviceStateInUse
	case config.AmiChannelStateRinging,
		config.AmiChannelStatePreRing:
		return config.AmiDeviceStateRinging
	case config.AmiChannelStateBusy:
		return config.AmiDeviceStateBusy
	}
	return config.AmiDeviceStateNotInUse
}

func (e *AMIExtensionStatus) SetActionId(value string) *AMIExtensionStatus {
	e.ActionId = value
	return e
}

func (e *AMIExtensionStatus) SetResponse(value string) *AMIExtensionStatus {
	e.Response = TrimStringSpaces(value)
	return e
}

func (e *AMIExtensionStatus) SetMessage(value string) *AMIExtensionStatus {
	e.Message = TrimStringSpaces(value)
	return e
}

func (e *AMIExtensionStatus) SetContext(value string) *AMIExtensionStatus {
	e.Context = TrimStringSpaces(value)
	return e
}

func (e *AMIExtensionStatus) SetExtension(value string) *AMIExtensionStatus {
	e.Extension = TrimStringSpaces(value)
	return e
}

func (e *AMIExtensionStatus) SetHint(value string) *AMIExtensionStatus {
	e.Hint = TrimStringSpaces(value)
	return e
}

func (e *AMIExtensionStatus) SetStatus(value string) *AMIExtensionStatus {
	status, _ := strconv.Atoi(value)
	e.Status = status
	return e
}

func (e *AMIExtensionStatus) SetStatusInt(value int) *AMIExtensionStatus {
	e.Status = value
	return e
}

func (e *AMIExtensionStatus) SetStatusText(value string) *AMIExtensionStatus {
	e.StatusText = TrimStringSpaces(value)
	return e
}

func (e *AMIExtensionGuard) SetEnabledExtensionNumeric(value bool) *AMIExtensionGuard {
	e.EnabledExtensionNumeric = value
	return e
}

func (e *AMIExtensionGuard) SetContexts(value []string) *AMIExtensionGuard {
	e.Context = value
	return e
}

func (e *AMIExtensionGuard) SetContext(value string) *AMIExtensionGuard {
	e.SetContexts([]string{value})
	return e
}

func (e *AMIExtensionGuard) AppendContext(value ...string) *AMIExtensionGuard {
	e.Context = append(e.Context, value...)
	return e
}

func (e *AMIExtensionGuard) SetStatusesText(values []string) *AMIExtensionGuard {
	e.StatusesText = values
	return e
}

func (e *AMIExtensionGuard) AppendStatusText(values ...string) *AMIExtensionGuard {
	e.StatusesText = append(e.StatusesText, values...)
	return e
}

func (e *AMIPeerStatus) SetActionId(value string) *AMIPeerStatus {
	e.ActionId = value
	return e
}

func (e *AMIPeerStatus) SetChannelType(value string) *AMIPeerStatus {
	e.ChannelType = TrimStringSpaces(value)
	return e
}

func (e *AMIPeerStatus) SetEvent(value string) *AMIPeerStatus {
	e.Event = TrimStringSpaces(value)
	return e
}

func (e *AMIPeerStatus) SetPeer(value string) *AMIPeerStatus {
	e.Peer = TrimStringSpaces(value)
	return e
}

func (e *AMIPeerStatus) SetPeerStatus(value string) *AMIPeerStatus {
	e.PeerStatus = TrimStringSpaces(value)
	return e
}

func (e *AMIPeerStatus) SetPrivilege(value string) *AMIPeerStatus {
	e.Privilege = TrimStringSpaces(value)
	return e
}

func (e *AMIPeerStatus) SetTimeInMs(value string) *AMIPeerStatus {
	v, err := strconv.Atoi(value)
	if err == nil {
		e.SetTimeInMsInt(v)
	}
	return e
}

func (e *AMIPeerStatus) SetTimeInMsInt(value int) *AMIPeerStatus {
	if value >= 0 {
		e.TimeInMs = value
	}
	return e
}

func (e *AMIPeerStatus) SetPrePublishedAt(value string) *AMIPeerStatus {
	e.PrePublishedAt = TrimStringSpaces(value)
	return e
}

func (e *AMIPeerStatus) SetPublishedAt(value time.Time) *AMIPeerStatus {
	e.PublishedAt = value
	return e
}

func (e *AMIPeerStatusGuard) SetDateTimeLayout(value string) *AMIPeerStatusGuard {
	e.DateTimeLayout = TrimStringSpaces(value)
	return e
}

func (e *AMIPeerStatusGuard) SetTimezone(value string) *AMIPeerStatusGuard {
	e.Timezone = TrimStringSpaces(value)
	return e
}
