package ami

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewChannel() *AMIChannel {
	return &AMIChannel{}
}

func (c *AMIChannel) SetChannelProtocol(protocol string) *AMIChannel {
	if ok := config.AmiChannelProtocols[protocol]; !ok {
		msg := fmt.Sprintf(config.AmiErrorProtocolMessage, strings.Join(GetKeys(config.AmiChannelProtocols), ","))
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

func (c *AMIChannel) Verify(regex string, extension string) bool {
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

func (c *AMIChannel) VerifyDefaultSIP(extension string) bool {
	c.SetChannelProtocol(config.AmiSIPChannelProtocol)
	return c.Verify(config.AmiDigitExtensionRegexDefault, extension)
}

func (c *AMIChannel) VerifyWith(channelProtocol string, regex string, digitsExten []interface{}, extension string) bool {
	if len(digitsExten) == 0 {
		return false
	}
	if IsStringEmpty(extension) {
		return false
	}
	if IsStringEmpty(regex) {
		return false
	}
	c.SetChannelProtocol(channelProtocol)
	var has bool = false

	for _, v := range digitsExten {
		_regex := fmt.Sprintf(regex, v)
		valid := c.Verify(_regex, extension)
		if valid {
			has = valid
			break
		}
	}
	return has
}

// VerifyDefaultSIPWith
// Example:
/*
	v := ami.NewChannel().VerifyDefaultSIPWith([]interface{}{4, 5, 6}, "SIP/8103")
	log.Printf("outgoing: %v", v)
*/
func (c *AMIChannel) VerifyDefaultSIPWith(digitsExten []interface{}, extension string) bool {
	c.SetChannelProtocol(config.AmiSIPChannelProtocol)
	return c.VerifyWith(c.ChannelProtocol, config.AmiDigitExtensionRegexWithDigits, digitsExten, extension)
}

// JoinHostChannel
// Example: protocol is SIP
// Ip is 127.0.0.1
// Return as form sip@127.0.0.1
func (c *AMIChannel) JoinHostChannel(protocol, ip string) string {
	c.SetChannelProtocol(protocol)
	host, _, _ := DecodeIp(ip)
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
func Originate(ctx context.Context, s AMISocket, originate AMIPayloadOriginate) (AMIResultRaw, error) {
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

// PlayDTMF plays DTMF signal on a specific channel.
func PlayDTMF(ctx context.Context, s AMISocket, channel, digit string, duration int) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPlayDtmf)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:  channel,
		config.AmiFieldDigit:    digit,
		config.AmiFieldDuration: duration,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Redirect redirects (transfer) a call.
func Redirect(ctx context.Context, s AMISocket, call AMIPayloadCall) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionRedirect)
	c.SetVCmd(call)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// SendText sends text message to channel.
func SendText(ctx context.Context, s AMISocket, channel, message string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSendText)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel: channel,
		config.AmiFieldMessage: message,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// SetVar sets a channel variable. Sets a global or local channel variable.
// Note: If a channel name is not provided then the variable is global.
func SetVar(ctx context.Context, s AMISocket, channel, variable, value string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSetVar)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:  channel,
		config.AmiFieldVariable: variable,
		config.AmiFieldValue:    value,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Status lists channel status.
// Will return the status information of each channel along with the value for the specified channel variables.
func Status(ctx context.Context, s AMISocket, channel, variables string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionStatus)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:   channel,
		config.AmiFieldVariables: variables,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// AOCMessage generates an Advice of Charge message on a channel.
func AOCMessage(ctx context.Context, s AMISocket, aoc AMIPayloadAOC) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionAocMessage)
	c.SetVCmd(aoc)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// GetVar get a channel variable.
func GetVar(ctx context.Context, s AMISocket, channel, variable string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionGetVar)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:  channel,
		config.AmiFieldVariable: variable,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// LocalOptimizeAway optimize away a local channel when possible.
// A local channel created with "/n" will not automatically optimize away.
// Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
func LocalOptimizeAway(ctx context.Context, s AMISocket, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionLocalOptimizeAway)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel: channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MuteAudio mute an audio stream.
func MuteAudio(ctx context.Context, s AMISocket, channel, direction string, state bool) (AMIResultRaw, error) {
	states := map[bool]string{false: "off", true: "on"}
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMuteAudio)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:   channel,
		config.AmiFieldDirection: direction,
		config.AmiFieldState:     states[state],
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
