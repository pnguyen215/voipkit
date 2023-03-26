package ami

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewChannel() *AMIChannel {
	return &AMIChannel{}
}

func (c *AMIChannel) SetChannelProtocol(protocol string) *AMIChannel {
	if ok := config.AmiChannelProtocols[protocol]; !ok {
		msg := fmt.Sprintf(config.AmiErrorProtocolMessage, strings.Join(utils.Keys(config.AmiChannelProtocols), ","))
		log.Panic(config.AmiErrorInvalidProtocol, "\n", msg)
	}
	c.ChannelProtocol = protocol
	return c
}

func (c *AMIChannel) SetNoDigitExtension(digit int) *AMIChannel {
	c.NoDigitExtension = digit
	return c
}

func (c *AMIChannel) SetExtension(extension string) *AMIChannel {
	c.Extension = extension
	return c
}

func (c *AMIChannel) Valid(regex string, extension string) bool {

	if strings.EqualFold(regex, "") {
		log.Printf(config.AmiErrorFieldRequired, "Regex")
		return false
	}

	if strings.EqualFold(extension, "") {
		log.Printf(config.AmiErrorFieldRequired, "Extension")
		return false
	}

	_regexp, err := regexp.Compile(regex)
	if err != nil {
		log.Printf("regex '%v' has an error compile occurred: %v", regex, err.Error())
		return false
	}

	match := _regexp.MatchString(extension)
	return match
}

func (c *AMIChannel) ValidSIPDefault(extension string) bool {
	c.SetChannelProtocol(config.AmiSIPChannelProtocol)
	return c.Valid(config.AmiDigitExtensionRegexDefault, extension)
}

// JoinHostChannel
// Example: protocol is SIP
// Ip is 127.0.0.1
// Return as form sip@127.0.0.1
func (c *AMIChannel) JoinHostChannel(protocol, ip string) string {
	c.SetChannelProtocol(protocol)
	host, _, _ := utils.IPDecode(ip)
	form := "%v@%v"

	if len(host) > 0 {
		return fmt.Sprintf(form, strings.ToLower(c.ChannelProtocol), host)
	}

	if strings.HasPrefix(ip, config.AmiProtocolHttpKey) {
		ip = strings.Replace(ip, config.AmiProtocolHttpKey, "", -1)
	}

	if strings.HasPrefix(ip, config.AmiProtocolHttpsKey) {
		ip = strings.Replace(ip, config.AmiProtocolHttpsKey, "", -1)
	}

	return fmt.Sprintf(form, strings.ToLower(c.ChannelProtocol), ip)
}

// JoinChannelWith
// Example: protocol is SIP
// Extension is 1010
// Return as form SIP/1010
func (c *AMIChannel) JoinChannelWith(protocol, extension string) string {
	c.SetChannelProtocol(protocol)

	if strings.EqualFold(extension, "") {
		log.Printf(config.AmiErrorFieldRequired, "Extension")
		return extension
	}

	form := "%v/%v"
	return fmt.Sprintf(form, c.ChannelProtocol, strings.TrimSpace(extension))
}

// CoreShowChannels list currently active channels.
func CoreShowChannels(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCoreShowChannels)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventCoreShowChannel}, []string{config.AmiListenerEventCoreShowChannelsComplete})
	return callback.SendSuperLevel()
}

// AbsoluteTimeout set absolute timeout.
// Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
func AbsoluteTimeout(ctx context.Context, s AMISocket, channel string, timeout int) (AMIResultRaw, error) {
	_timeout := strconv.Itoa(timeout)
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionAbsoluteTimeout)
	c.SetV(map[string]string{
		config.AmiFieldChannel: channel,
		config.AmiFieldTimeout: _timeout,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Hangup hangups channel.
func Hangup(ctx context.Context, s AMISocket, channel, cause string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionHangup)
	c.SetV(map[string]string{
		config.AmiFieldChannel: channel,
		config.AmiFieldCause:   cause,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Originate originates a call.
// Generates an outgoing call to a Extension/Context/Priority or Application/Data.
func Originate(ctx context.Context, s AMISocket, originate AMIOriginateData) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionOriginate)
	c.SetVCmd(originate)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ParkedCalls list parked calls.
func ParkedCalls(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionParkedCalls)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventParkedCall}, []string{config.AmiListenerEventParkedCallsComplete})
	return callback.SendSuperLevel()
}

// Park parks a channel.
func Park(ctx context.Context, s AMISocket, channel1, channel2 string, timeout int, parkinglot string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionHangup)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:    channel1,
		config.AmiFieldChannel2:   channel2,
		config.AmiFieldTimeout:    timeout,
		config.AmiFieldParkinglot: parkinglot,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Parkinglots get a list of parking lots.
func Parkinglots(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionParkingLots)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventParkedCall}, []string{config.AmiListenerEventParkinglotsComplete})
	return callback.SendSuperLevel()
}
