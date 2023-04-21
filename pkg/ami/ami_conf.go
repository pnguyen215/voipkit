package ami

import "github.com/pnguyen215/gobase-voip-core/pkg/ami/config"

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
