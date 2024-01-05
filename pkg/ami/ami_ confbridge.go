package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// ConfbridgeList lists all users in a particular ConfBridge conference.
func ConfbridgeList(ctx context.Context, s AMISocket, conference string) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeList).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
	})
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventConfbridgeList}, []string{config.AmiListenerEventConfbridgeListComplete})
	return callback.SendSuperLevel()
}

// ConfbridgeListRooms lists data about all active conferences.
func ConfbridgeListRooms(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeListRooms)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventConfbridgeListRooms}, []string{config.AmiListenerEventConfbridgeListRoomsComplete})
	return callback.SendSuperLevel()
}

// ConfbridgeMute mutes a specified user in a specified conference.
func ConfbridgeMute(ctx context.Context, s AMISocket, conference string, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeMute).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
		config.AmiFieldChannel:    channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ConfbridgeUnmute unmute a specified user in a specified conference.
func ConfbridgeUnmute(ctx context.Context, s AMISocket, conference string, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeUnmute).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
		config.AmiFieldChannel:    channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ConfbridgeKick removes a specified user from a specified conference.
func ConfbridgeKick(ctx context.Context, s AMISocket, conference string, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeKick).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
		config.AmiFieldChannel:    channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ConfbridgeLock locks a specified conference.
func ConfbridgeLock(ctx context.Context, s AMISocket, conference string, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeLock).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
		config.AmiFieldChannel:    channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ConfbridgeUnlock unlocks a specified conference.
func ConfbridgeUnlock(ctx context.Context, s AMISocket, conference string, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeUnlock).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
		config.AmiFieldChannel:    channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ConfbridgeSetSingleVideoSrc sets a conference user as the single video source distributed to all other video-capable participants.
func ConfbridgeSetSingleVideoSrc(ctx context.Context, s AMISocket, conference string, channel string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeSetSingleVideoSrc).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
		config.AmiFieldChannel:    channel,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ConfbridgeStartRecord starts a recording in the context of given conference and creates a file with the name specified by recordFile
func ConfbridgeStartRecord(ctx context.Context, s AMISocket, conference string, recordFile string) (AMIResultRaw, error) {
	params := map[string]string{
		config.AmiFieldConference: conference,
	}
	if len(recordFile) > 0 {
		params[config.AmiFieldRecordFile] = recordFile
	}
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeStartRecord).SetVCmd(params)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ConfbridgeStopRecord stops a recording pertaining to the given conference
func ConfbridgeStopRecord(ctx context.Context, s AMISocket, conference string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionConfbridgeStopRecord).SetVCmd(map[string]interface{}{
		config.AmiFieldConference: conference,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
