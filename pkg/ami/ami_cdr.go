package ami

import (
	"strconv"
	"time"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewAMICdr() *AMICdr {
	r := &AMICdr{}
	r.SetEvent(config.AmiListenerEventCdr)
	return r
}

func (r *AMICdr) SetEvent(value string) *AMICdr {
	r.Event = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetAccountCode(value string) *AMICdr {
	r.AccountCode = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetSource(value string) *AMICdr {
	r.Source = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetDestination(value string) *AMICdr {
	r.Destination = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetDestinationContext(value string) *AMICdr {
	r.DestinationContext = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetCallerId(value string) *AMICdr {
	r.CallerId = value
	return r
}

func (r *AMICdr) SetChannel(value string) *AMICdr {
	r.Channel = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetDestinationChannel(value string) *AMICdr {
	r.DestinationChannel = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetLastApplication(value string) *AMICdr {
	r.LastApplication = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetLastData(value string) *AMICdr {
	r.LastData = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetStartTime(value time.Time) *AMICdr {
	r.StartTime = value
	return r
}

func (r *AMICdr) SetStartTimeWith(value string) *AMICdr {
	t, err := time.Parse(config.DateTimeFormatYYYYMMDDHHMMSS, value)
	if err == nil {
		r.SetStartTime(t)
	}
	return r
}

func (r *AMICdr) SetStartTimeWithLayout(layout, value string) *AMICdr {
	t, err := time.Parse(layout, value)
	if err == nil {
		r.SetStartTime(t)
	}
	return r
}

func (r *AMICdr) SetAnswerTime(value time.Time) *AMICdr {
	r.AnswerTime = value
	return r
}

func (r *AMICdr) SetAnswerTimeWith(value string) *AMICdr {
	t, err := time.Parse(config.DateTimeFormatYYYYMMDDHHMMSS, value)
	if err == nil {
		r.SetAnswerTime(t)
	}
	return r
}

func (r *AMICdr) SetAnswerTimeWithLayout(layout, value string) *AMICdr {
	t, err := time.Parse(layout, value)
	if err == nil {
		r.SetAnswerTime(t)
	}
	return r
}

func (r *AMICdr) SetEndTime(value time.Time) *AMICdr {
	r.EndTime = value
	return r
}

func (r *AMICdr) SetEndTimeWith(value string) *AMICdr {
	t, err := time.Parse(config.DateTimeFormatYYYYMMDDHHMMSS, value)
	if err == nil {
		r.SetEndTime(t)
	}
	return r
}

func (r *AMICdr) SetEndTimeWithLayout(layout, value string) *AMICdr {
	t, err := time.Parse(layout, value)
	if err == nil {
		r.SetEndTime(t)
	}
	return r
}

func (r *AMICdr) SetDuration(value int) *AMICdr {
	if value >= 0 {
		r.Duration = value
	}
	return r
}

func (r *AMICdr) SetDurationWith(value string) *AMICdr {
	v, err := strconv.Atoi(value)
	if err == nil {
		r.SetDuration(v)
	}
	return r
}

func (r *AMICdr) SetBillableSecond(value int) *AMICdr {
	if value >= 0 {
		r.BillableSeconds = value
	}
	return r
}

func (r *AMICdr) SetBillableSecondWith(value string) *AMICdr {
	v, err := strconv.Atoi(value)
	if err == nil {
		r.SetBillableSecond(v)
	}
	return r
}

func (r *AMICdr) SetDisposition(value string) *AMICdr {
	r.Disposition = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetAmaFlag(value string) *AMICdr {
	r.AmaFlags = utils.TrimAllSpace(value)
	return r
}

func (r *AMICdr) SetUniqueId(value string) *AMICdr {
	r.UniqueId = value
	return r
}

func (r *AMICdr) SetUserField(value string) *AMICdr {
	r.UserField = value
	return r
}

func (r *AMICdr) SetDateReceivedAt(value time.Time) *AMICdr {
	r.DateReceivedAt = value
	return r
}

func (r *AMICdr) SetDateReceivedAtWith(layout, value string) *AMICdr {
	t, err := time.Parse(layout, value)
	if err == nil {
		r.SetDateReceivedAt(t)
	}
	return r
}

func (r *AMICdr) SetPrivilege(value string) *AMICdr {
	r.Privilege = value
	return r
}

func (r *AMICdr) Json() string {
	return utils.ToJson(r)
}

func ParseCdr(e *AMIMessage, d *AMIDictionary) *AMICdr {
	r := NewAMICdr().
		SetAccountCode(e.FieldByDictionary(d, config.AmiJsonFieldAccountCode)).
		SetAmaFlag(e.FieldByDictionary(d, config.AmiJsonFieldAmaFlags)).
		SetAnswerTimeWith(e.FieldByDictionary(d, config.AmiJsonFieldAnswerTime)).
		SetBillableSecondWith(e.FieldByDictionary(d, config.AmiJsonFieldBillableSeconds)).
		SetCallerId(e.FieldByDictionary(d, config.AmiJsonFieldCallerId)).
		SetChannel(e.FieldByDictionary(d, config.AmiJsonFieldChannel)).
		SetDateReceivedAtWith(e.DateTimeLayout, e.FieldByDictionary(d, config.AmiJsonFieldDateReceivedAt)).
		SetDestination(e.FieldByDictionary(d, config.AmiJsonFieldDestination)).
		SetDestinationChannel(e.FieldByDictionary(d, config.AmiJsonFieldDestinationChannel)).
		SetDestinationContext(e.FieldByDictionary(d, config.AmiJsonFieldDestinationContext)).
		SetDisposition(e.FieldByDictionary(d, config.AmiJsonFieldDisposition)).
		SetDurationWith(e.FieldByDictionary(d, config.AmiJsonFieldDuration)).
		SetEndTimeWith(e.FieldByDictionary(d, config.AmiJsonFieldEndTime)).
		SetLastApplication(e.FieldByDictionary(d, config.AmiJsonFieldLastApplication)).
		SetLastData(e.FieldByDictionary(d, config.AmiJsonFieldLastData)).
		SetPrivilege(e.FieldByDictionary(d, config.AmiJsonFieldPrivilege)).
		SetSource(e.FieldByDictionary(d, config.AmiJsonFieldSource)).
		SetStartTimeWith(e.FieldByDictionary(d, config.AmiJsonFieldStartTime)).
		SetUniqueId(e.FieldByDictionary(d, config.AmiJsonFieldUniqueId)).
		SetUserField(e.FieldByDictionary(d, config.AmiJsonFieldUserField))
	return r
}
