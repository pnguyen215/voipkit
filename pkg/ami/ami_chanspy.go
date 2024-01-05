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

func (s *AMIPayloadChanspy) SetAllowDebug(value bool) *AMIPayloadChanspy {
	s.AllowDebug = value
	return s
}

func (s *AMIPayloadChanspy) CommandChanspy(channelExten string) string {
	if IsStringEmpty(s.Join) {
		return ""
	}
	if IsStringEmpty(channelExten) {
		return ""
	}
	if strings.EqualFold(s.Join, config.AmiChanspySpy) {
		return channelExten
	}
	if strings.EqualFold(s.Join, config.AmiChanspyWhisper) {
		return fmt.Sprintf("%s,w", channelExten)
	}
	if strings.EqualFold(s.Join, config.AmiChanspyBarge) {
		return fmt.Sprintf("%s,B", channelExten)
	}
	return channelExten
}

// Chanspy
func Chanspy(ctx context.Context, s AMISocket, ch AMIPayloadChanspy) (AMIResultRawLevel, error) {
	ok := config.AmiChanspy[ch.Join]
	if !ok {
		msg := fmt.Sprintf(config.AmiErrorChanspyMessage, strings.Join(GetKeys(config.AmiChanspy), ","))
		log.Panic(config.AmiErrorInvalidChanspy, "\n", msg)
	}
	if IsStringEmpty(ch.SourceExten) {
		return AMIResultRawLevel{}, fmt.Errorf("Source exten is required")
	}
	if IsStringEmpty(ch.CurrentExten) {
		return AMIResultRawLevel{}, fmt.Errorf("Current exten is required")
	}
	sourceValid, err := HasSIPPeerStatus(ctx, s, ch.SourceExten)
	if err != nil {
		return AMIResultRawLevel{}, err
	}
	if !sourceValid {
		return AMIResultRawLevel{}, fmt.Errorf("Source exten '%v' not found", ch.SourceExten)
	}
	currentValid, err := HasSIPPeerStatus(ctx, s, ch.CurrentExten)
	if err != nil {
		return AMIResultRawLevel{}, err
	}
	if !currentValid {
		return AMIResultRawLevel{}, fmt.Errorf("Current exten '%v' not found", ch.CurrentExten)
	}
	channel := NewChannel().SetChannelProtocol(ch.ChannelProtocol)
	sourceExt := channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", ch.SourceExten))
	currentExt := channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", ch.CurrentExten))
	channelId := ch.CommandChanspy(sourceExt)
	cmd := fmt.Sprintf("channel originate %s application ChanSpy %s", currentExt, channelId)
	if ch.AllowDebug {
		log.Printf("Chanspy command: %v \n", cmd)
		log.Printf("Chanspy channel_id: %v \n", channelId)
		log.Printf("Chanspy source.ext: %v \n", sourceExt)
		log.Printf("Chanspy current.ext: %v \n", currentExt)
	}
	return Command(ctx, s, cmd)
}
