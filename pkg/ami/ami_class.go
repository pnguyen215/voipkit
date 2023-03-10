package ami

import (
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

var chargingEvents *map[string]string = &map[string]string{}

// SnapChargingEvent
func (e *AMIEvent) SnapChargingEvent() *map[string]string {

	if len(*chargingEvents) > 0 {
		return chargingEvents
	}

	_merged := make(map[string]string)

	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassCommands))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassSecurities))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassCalls))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassSystems))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassUsers))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassDialPlans))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassAgents))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassAgis))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassAocs))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassCallDetailRecords))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassChannelEventLoggings))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClasses))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassReports))
	_merged = utils.MergeMaps(_merged, e.SwapCharging(config.AmiClassDualToneMultiFrequencies))

	chargingEvents = &_merged
	return chargingEvents
}

// SnapChargingEventWith
func (e *AMIEvent) SnapChargingEventWith(ls map[string][]string) *map[string]string {
	chargingEvents := utils.MergeMaps(*chargingEvents, e.SwapCharging(ls))
	return &chargingEvents
}

// SwapCharging
func (e *AMIEvent) SwapCharging(ls map[string][]string) map[string]string {

	if len(ls) <= 0 {
		return map[string]string{}
	}

	data := make(map[string]string)

	for key, value := range ls {

		if len(value) > 0 {
			for _, v := range value {
				data[v] = key
			}
		}
	}

	return data
}

// FindChargingValue
func (e *AMIEvent) FindChargingValue(ls map[string][]string, event string) (v string, ok bool) {
	data := e.SwapCharging(ls)
	v, ok = data[event]
	return v, ok
}

// FindChargingValueWith
func (e *AMIEvent) FindChargingValueWith(event string) (v string, ok bool) {
	data := *e.SnapChargingEvent()
	v, ok = data[event]
	return v, ok
}

// ResetChargingEvent
func (e *AMIEvent) ResetChargingEvent() {
	*chargingEvents = nil
	chargingEvents = e.SnapChargingEvent()
}

// LengthChargingEvent
func (e *AMIEvent) LengthChargingEvent() int {
	return len(*chargingEvents)
}
