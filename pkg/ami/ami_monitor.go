package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAMIPayloadMonitor() *AMIPayloadMonitor {
	m := &AMIPayloadMonitor{}
	return m
}

func (m *AMIPayloadMonitor) SetChannel(value string) *AMIPayloadMonitor {
	m.Channel = value
	return m
}

func (m *AMIPayloadMonitor) SetDirection(value string) *AMIPayloadMonitor {
	m.Direction = value
	return m
}

func (m *AMIPayloadMonitor) SetState(value string) *AMIPayloadMonitor {
	m.State = value
	return m
}

func (m *AMIPayloadMonitor) SetFile(value string) *AMIPayloadMonitor {
	m.File = value
	return m
}

func (m *AMIPayloadMonitor) SetFormat(value string) *AMIPayloadMonitor {
	m.Format = value
	return m
}

func (m *AMIPayloadMonitor) SetMix(value bool) *AMIPayloadMonitor {
	m.Mix = value
	return m
}

func (m *AMIPayloadMonitor) SetMixMonitorId(value string) *AMIPayloadMonitor {
	m.MixMonitorId = value
	return m
}

// Monitor monitors a channel.
// This action may be used to record the audio on a specified channel.
func Monitor(ctx context.Context, s AMISocket, payload AMIPayloadMonitor) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMonitor)
	c.SetVCmd(payload)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Monitor monitors a channel.
// This action may be used to record the audio on a specified channel.
func MonitorWith(ctx context.Context, s AMISocket, channel, file, format string, mix bool) (AmiReply, error) {
	p := NewAMIPayloadMonitor().SetChannel(channel).SetFile(file).SetFormat(format).SetMix(mix)
	return Monitor(ctx, s, *p)
}

// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func ChangeMonitor(ctx context.Context, s AMISocket, payload AMIPayloadMonitor) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionChangeMonitor)
	c.SetVCmd(payload)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func ChangeMonitorWith(ctx context.Context, s AMISocket, channel, file string) (AmiReply, error) {
	p := NewAMIPayloadMonitor().SetChannel(channel).SetFile(file)
	return ChangeMonitor(ctx, s, *p)
}

// MixMonitor record a call and mix the audio during the recording.
func MixMonitor(ctx context.Context, s AMISocket, payload AMIPayloadMonitor) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMixMonitor)
	c.SetVCmd(payload)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MixMonitor record a call and mix the audio during the recording.
func MixMonitorWith(ctx context.Context, s AMISocket, channel, file, options, command string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMixMonitor)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel: channel,
		config.AmiFieldFile:    file,
		config.AmiFieldOptions: options,
		config.AmiFieldCommand: command,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MixMonitorMute Mute / unMute a Mixmonitor recording.
// This action may be used to mute a MixMonitor recording.
func MixMonitorMute(ctx context.Context, s AMISocket, channel, direction string, state bool) (AmiReply, error) {
	states := map[bool]string{false: "0", true: "1"}
	p := NewAMIPayloadMonitor().SetChannel(channel).SetDirection(direction).SetState(states[state])
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMixMonitorMute)
	c.SetVCmd(p)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PauseMonitor pauses monitoring of a channel.
// This action may be used to temporarily stop the recording of a channel.
func PauseMonitor(ctx context.Context, s AMISocket, channel string) (AmiReply, error) {
	p := NewAMIPayloadMonitor().SetChannel(channel)
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPauseMonitor)
	c.SetVCmd(p)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// UnpauseMonitor unpause monitoring of a channel.
// This action may be used to re-enable recording of a channel after calling PauseMonitor.
func UnpauseMonitor(ctx context.Context, s AMISocket, channel string) (AmiReply, error) {
	p := NewAMIPayloadMonitor().SetChannel(channel)
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionUnpauseMonitor)
	c.SetVCmd(p)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// StopMonitor stops monitoring a channel.
// This action may be used to end a previously started 'Monitor' action.
func StopMonitor(ctx context.Context, s AMISocket, channel string) (AmiReply, error) {
	p := NewAMIPayloadMonitor().SetChannel(channel)
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionStopMonitor)
	c.SetVCmd(p)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// StopMixMonitor stop recording a call through MixMonitor, and free the recording's file handle.
func StopMixMonitor(ctx context.Context, s AMISocket, channel, mixMonitorId string) (AmiReply, error) {
	p := NewAMIPayloadMonitor().SetChannel(channel).SetMixMonitorId(mixMonitorId)
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionStopMixMonitor)
	c.SetVCmd(p)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
