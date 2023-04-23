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

func (r *AMICdr) IsCdrNoAnswer() bool {
	_, ok := config.AmiCdrDispositionText[r.Disposition]
	if !ok {
		return false
	}
	return r.Disposition == config.AmiCdrDispositionNoAnswer
}

func (r *AMICdr) IsCdrFailed() bool {
	_, ok := config.AmiCdrDispositionText[r.Disposition]
	if !ok {
		return false
	}
	return r.Disposition == config.AmiCdrDispositionFailed
}

func (r *AMICdr) IsCdrBusy() bool {
	_, ok := config.AmiCdrDispositionText[r.Disposition]
	if !ok {
		return ok
	}
	return r.Disposition == config.AmiCdrDispositionBusy
}

func (r *AMICdr) IsCdrAnswered() bool {
	_, ok := config.AmiCdrDispositionText[r.Disposition]
	if !ok {
		return ok
	}
	return r.Disposition == config.AmiCdrDispositionAnswered
}

func (r *AMICdr) IsCdrCongestion() bool {
	_, ok := config.AmiCdrDispositionText[r.Disposition]
	if !ok {
		return ok
	}
	return r.Disposition == config.AmiCdrDispositionCongestion
}

func (r *AMICdr) IsCdrFlagOmit() bool {
	return r.AmaFlags == config.AmiAmaFlagOmit
}

func (r *AMICdr) IsCdrFlagBilling() bool {
	return r.AmaFlags == config.AmiAmaFlagBilling
}

func (r *AMICdr) IsCdrFlagDocumentation() bool {
	return r.AmaFlags == config.AmiAmaFlagDocumentation
}

func ParseCdr(e *AMIMessage, d *AMIDictionary) *AMICdr {
	if d == nil {
		d = NewDictionary()
	}
	if !d.AllowForceTranslate {
		d.SetAllowForceTranslate(true)
	}
	r := NewAMICdr().
		SetAccountCode(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldAccountCode, "AccountCode")).
		SetAmaFlag(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldAmaFlags, "AmaFlags")).
		SetAnswerTimeWith(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldAnswerTime, "AnswerTime")).
		SetBillableSecondWith(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldBillableSeconds, "BillableSeconds")).
		SetCallerId(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldCallerId, "CallerID")).
		SetChannel(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldChannel, "Channel")).
		SetDateReceivedAtWith(e.DateTimeLayout, e.FieldDictionaryOrRefer(d, config.AmiJsonFieldDateReceivedAt, "DateReceivedAt")).
		SetDestination(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldDestination, "Destination")).
		SetDestinationChannel(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldDestinationChannel, "DestinationChannel")).
		SetDestinationContext(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldDestinationContext, "DestinationContext")).
		SetDisposition(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldDisposition, "Disposition")).
		SetDurationWith(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldDuration, "Duration")).
		SetEndTimeWith(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldEndTime, "EndTime")).
		SetLastApplication(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldLastApplication, "LastApplication")).
		SetLastData(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldLastData, "LastData")).
		SetPrivilege(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldPrivilege, "Privilege")).
		SetSource(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldSource, "Source")).
		SetStartTimeWith(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldStartTime, "StartTime")).
		SetUniqueId(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldUniqueId, "UniqueID")).
		SetUserField(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldUserField, "UserField"))
	return r
}
