package ami

import (
	"fmt"
	"log"
	"regexp"
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
