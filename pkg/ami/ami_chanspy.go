package ami

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAMIPayloadChanspy() *AMIPayloadChanspy {
	s := &AMIPayloadChanspy{}
	return s
}

func (s *AMIPayloadChanspy) SetJoin(value string) *AMIPayloadChanspy {
	ok := config.AmiChanspy[value]
	if !ok {
		msg := fmt.Sprintf(config.AmiErrorChanspyMessage, strings.Join(GetKeys(config.AmiChanspy), ","))
		log.Panic(config.AmiErrorInvalidChanspy, "\n", msg)
	}
	s.Join = TrimStringSpaces(value)
	return s
}

func (s *AMIPayloadChanspy) SetSourceExten(value string) *AMIPayloadChanspy {
	s.SourceExten = TrimStringSpaces(value)
	return s
}

func (s *AMIPayloadChanspy) SetCurrentExten(value string) *AMIPayloadChanspy {
	s.CurrentExten = TrimStringSpaces(value)
	return s
}

func (s *AMIPayloadChanspy) SetChannelProtocol(value string) *AMIPayloadChanspy {
	channel := NewChannel().SetChannelProtocol(value)
	s.ChannelProtocol = channel.ChannelProtocol
	return s
}

func (s *AMIPayloadChanspy) SetDebugMode(value bool) *AMIPayloadChanspy {
	s.DebugMode = value
	return s
}

func (s *AMIPayloadChanspy) CommandChanspy(c_extension string) string {
	if IsStringEmpty(s.Join) {
		return ""
	}
	if IsStringEmpty(c_extension) {
		return ""
	}
	if strings.EqualFold(s.Join, config.AmiChanspySpy) {
		return c_extension
	}
	if strings.EqualFold(s.Join, config.AmiChanspyWhisper) {
		return fmt.Sprintf("%s,w", c_extension)
	}
	if strings.EqualFold(s.Join, config.AmiChanspyBarge) {
		return fmt.Sprintf("%s,B", c_extension)
	}
	return c_extension
}

// Chanspy
func Chanspy(ctx context.Context, s AMISocket, ch AMIPayloadChanspy) (AMIResultRawLevel, error) {
	ok := config.AmiChanspy[ch.Join]
	if !ok {
		msg := fmt.Sprintf(config.AmiErrorChanspyMessage, strings.Join(GetKeys(config.AmiChanspy), ","))
		log.Panic(config.AmiErrorInvalidChanspy, "\n", msg)
	}
	if IsStringEmpty(ch.SourceExten) {
		return AMIResultRawLevel{}, fmt.Errorf("Source extension is required")
	}
	if IsStringEmpty(ch.CurrentExten) {
		return AMIResultRawLevel{}, fmt.Errorf("Current extension is required")
	}
	source_extension_verify, err := HasSIPPeerStatus(ctx, s, ch.SourceExten)
	if err != nil {
		return AMIResultRawLevel{}, err
	}
	if !source_extension_verify {
		return AMIResultRawLevel{}, fmt.Errorf("Source extension '%v' not found", ch.SourceExten)
	}
	current_extension_verify, err := HasSIPPeerStatus(ctx, s, ch.CurrentExten)
	if err != nil {
		return AMIResultRawLevel{}, err
	}
	if !current_extension_verify {
		return AMIResultRawLevel{}, fmt.Errorf("Current extension '%v' not found", ch.CurrentExten)
	}
	channel := NewChannel().SetChannelProtocol(ch.ChannelProtocol)
	sourceExt := channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", ch.SourceExten))
	currentExt := channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", ch.CurrentExten))
	channelId := ch.CommandChanspy(sourceExt)
	cmd := fmt.Sprintf("channel originate %s application ChanSpy %s", currentExt, channelId)
	if ch.DebugMode {
		log.Printf("Chanspy command: %v \n", cmd)
		log.Printf("Chanspy channel_id: %v \n", channelId)
		log.Printf("Chanspy source.ext: %v \n", sourceExt)
		log.Printf("Chanspy current.ext: %v \n", currentExt)
	}
	return Command(ctx, s, cmd)
}
