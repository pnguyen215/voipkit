package ami

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAmiChanspy() *AMIChanspy {
	s := &AMIChanspy{}
	return s
}

func (s *AMIChanspy) SetJoin(value string) *AMIChanspy {
	ok := config.AmiChanspy[value]
	if !ok {
		msg := fmt.Sprintf(config.AmiErrorChanspyMessage, strings.Join(GetKeys(config.AmiChanspy), ","))
		log.Panic(config.AmiErrorInvalidChanspy, "\n", msg)
	}
	s.Join = TrimStringSpaces(value)
	return s
}

func (s *AMIChanspy) SetExtensionConnected(value string) *AMIChanspy {
	s.ExtensionConnected = TrimStringSpaces(value)
	return s
}

func (s *AMIChanspy) SetExtensionJoined(value string) *AMIChanspy {
	s.ExtensionJoined = TrimStringSpaces(value)
	return s
}

func (s *AMIChanspy) SetChannelProtocol(value string) *AMIChanspy {
	channel := NewChannel().SetChannelProtocol(value)
	s.ChannelProtocol = channel.ChannelProtocol
	return s
}

func (s *AMIChanspy) SetDebugMode(value bool) *AMIChanspy {
	s.DebugMode = value
	return s
}

func (s *AMIChanspy) CommandChanspy(c_extension string) string {
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

func Chanspy(ctx context.Context, s AMISocket, ch AMIChanspy) (AmiReplies, error) {
	ok := config.AmiChanspy[ch.Join]
	if !ok {
		msg := fmt.Sprintf(config.AmiErrorChanspyMessage, strings.Join(GetKeys(config.AmiChanspy), ","))
		D().Warn(msg)
		return nil, fmt.Errorf(config.AmiErrorInvalidChanspy)
	}
	if IsStringEmpty(ch.ExtensionConnected) {
		return AmiReplies{}, fmt.Errorf("Extension connected is required")
	}
	if IsStringEmpty(ch.ExtensionJoined) {
		return AmiReplies{}, fmt.Errorf("Extension joined is required")
	}
	extension_connected_verify, err := SIPPeerStatusExists(ctx, s, ch.ExtensionConnected)
	if err != nil {
		return AmiReplies{}, err
	}
	if !extension_connected_verify {
		return AmiReplies{}, fmt.Errorf("Extension connected '%v' not found", ch.ExtensionConnected)
	}
	extension_joined_verify, err := SIPPeerStatusExists(ctx, s, ch.ExtensionJoined)
	if err != nil {
		return AmiReplies{}, err
	}
	if !extension_joined_verify {
		return AmiReplies{}, fmt.Errorf("Extension joined '%v' not found", ch.ExtensionJoined)
	}
	channel := NewChannel().SetChannelProtocol(ch.ChannelProtocol)
	extensionConnected := channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", ch.ExtensionConnected))
	extensionJoined := channel.JoinChannelWith(channel.ChannelProtocol, fmt.Sprintf("%v", ch.ExtensionJoined))
	channelId := ch.CommandChanspy(extensionConnected)
	cmd := fmt.Sprintf("channel originate %s application ChanSpy %s", extensionJoined, channelId)
	if ch.DebugMode {
		D().Debug("Chanspy command: %v", cmd)
		D().Debug("Chanspy channel_id: %v", channelId)
		D().Debug("Chanspy extension.connected: %v", extensionConnected)
		D().Debug("Chanspy extension.joined: %v", extensionJoined)
	}
	return Command(ctx, s, cmd)
}
