package ami

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAMICdr() *AMICdr {
	r := &AMICdr{}
	r.SetEvent(config.AmiListenerEventCdr)
	r.SetSymbol("-")
	return r
}

func (r *AMICdr) SetEvent(value string) *AMICdr {
	r.Event = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetAccountCode(value string) *AMICdr {
	r.AccountCode = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetSource(value string) *AMICdr {
	r.Source = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetDestination(value string) *AMICdr {
	r.Destination = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetDestinationContext(value string) *AMICdr {
	r.DestinationContext = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetCallerId(value string) *AMICdr {
	r.CallerId = value
	return r
}

func (r *AMICdr) SetChannel(value string) *AMICdr {
	r.Channel = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetDestinationChannel(value string) *AMICdr {
	r.DestinationChannel = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetLastApplication(value string) *AMICdr {
	r.LastApplication = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetLastData(value string) *AMICdr {
	r.LastData = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetStartTime(value time.Time) *AMICdr {
	r.StartTime = value
	return r
}

func (r *AMICdr) SetStartTimeWith(value string) *AMICdr {
	t, err := time.Parse(config.DateTimeFormat20060102150405, value)
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
	t, err := time.Parse(config.DateTimeFormat20060102150405, value)
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
	t, err := time.Parse(config.DateTimeFormat20060102150405, value)
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
	r.Disposition = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetAmaFlag(value string) *AMICdr {
	r.AmaFlags = TrimStringSpaces(value)
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

func (r *AMICdr) SetDirection(value string) *AMICdr {
	r.Direction = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetDesc(value string) *AMICdr {
	r.Desc = value
	return r
}

func (r *AMICdr) SetType(value string) *AMICdr {
	r.Type = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetExtension(value string) *AMICdr {
	r.Extension = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetNumber(value string) *AMICdr {
	r.Number = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetSymbol(value string) *AMICdr {
	r.symbol = TrimStringSpaces(value)
	return r
}

func (r *AMICdr) SetMediaLink(value string) *AMICdr {
	r.MediaLink = value
	return r
}

func (r *AMICdr) Json() string {
	return JsonString(r)
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

func (r *AMICdr) IsCdrInbound() bool {
	return strings.EqualFold(r.Direction, config.AmiInboundDirection)
}

func (r *AMICdr) IsCdrOutbound() bool {
	return strings.EqualFold(r.Direction, config.AmiOutboundDirection)
}

func (r *AMICdr) IsCdrInboundDial() bool {
	return strings.EqualFold(r.Type, config.AmiTypeInboundDialDirection)
}

func (r *AMICdr) IsCdrInboundQueue() bool {
	return strings.EqualFold(r.Type, config.AmiTypeInboundQueueDirection)
}

func (r *AMICdr) IsCdrOutboundNormal() bool {
	return strings.EqualFold(r.Type, config.AmiTypeOutboundNormalDirection)
}

func (r *AMICdr) IsCdrOutboundChanSpy() bool {
	return strings.EqualFold(r.Type, config.AmiLastApplicationChanSpy)
}

func ParseCdr(e *AMIMessage, d *AMIDictionary) *AMICdr {
	if d == nil {
		d = NewDictionary()
	}
	if !d.EnabledForceTranslate {
		d.SetEnabledForceTranslate(true)
	}
	r := NewAMICdr().
		SetAccountCode(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldAccountCode, "AccountCode")).
		SetAmaFlag(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldAmaFlags, "AmaFlags")).
		SetAnswerTimeWith(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldAnswerTime, "AnswerTime")).
		SetBillableSecondWith(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldBillableSeconds, "BillableSeconds")).
		SetCallerId(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldCallerId, "CallerID")).
		SetChannel(e.FieldDictionaryOrRefer(d, config.AmiJsonFieldChannel, "Channel")).
		SetDateReceivedAtWith(e.TimeFormat, e.FieldDictionaryOrRefer(d, config.AmiJsonFieldDateReceivedAt, "DateReceivedAt")).
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

	// detect outbound, inbound
	// if the field destination is phone number, so mark this cdr belong to outbound, otherwise mark as inbound
	form := "CDR.call_from_'%v'_to_'%v'"
	phone := RemoveStringPrefix(r.Destination, e.PhonePrefix...)
	if VerifyPhoneNo(phone, e.Region) {
		flow := fmt.Sprintf(form, r.Channel, phone)
		r.SetDesc(flow)
		r.SetDirection(config.AmiOutboundDirection)
		r.SetType(config.AmiTypeOutboundNormalDirection)
		r.SetExtension(strings.Split(r.Channel, r.symbol)[0])
		r.SetNumber(phone)
	} else {
		var inCase bool = false
		// from outbound chan-spy
		if strings.EqualFold(r.LastApplication, config.AmiLastApplicationChanSpy) {
			inCase = true
			flow := fmt.Sprintf(form, r.Channel, r.LastData)
			r.SetDesc(flow)
			r.SetType(config.AmiTypeChanSpyDirection)
			r.SetDirection(config.AmiOutboundDirection)
			r.SetExtension(strings.Split(r.Channel, r.symbol)[0])
		}
		// from inbound dial
		if strings.EqualFold(r.LastApplication, config.AmiLastApplicationDial) {
			inCase = true
			flow := fmt.Sprintf(form, r.Source, r.DestinationChannel)
			r.SetDesc(flow)
			r.SetDirection(config.AmiInboundDirection)
			r.SetType(config.AmiTypeInboundDialDirection)
			r.SetExtension(strings.Split(r.DestinationChannel, r.symbol)[0])
			r.SetNumber(r.Source)
		}
		// from inbound queue
		if strings.EqualFold(r.LastApplication, config.AmiLastApplicationQueue) {
			inCase = true
			flow := fmt.Sprintf(form, r.Source, r.Channel)
			r.SetDesc(flow)
			r.SetDirection(config.AmiInboundDirection)
			r.SetType(config.AmiTypeInboundQueueDirection)
			r.SetExtension(strings.Split(r.Channel, r.symbol)[0])
			r.SetNumber(r.Source)
		}
		if !inCase {
			D().Error("ParseCdr, CDR got an error exception case:: %v", JsonString(r))
		}
	}
	return r
}
