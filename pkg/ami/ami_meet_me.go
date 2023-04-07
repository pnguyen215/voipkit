package ami

import (
	"context"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

// MeetMeList lists all users in a particular MeetMe conference.
// Will follow as separate events, followed by a final event called MeetmeListComplete.
func MeetMeList(ctx context.Context, s AMISocket, conference string) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMeetMeList)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventMeetMeEntry}, []string{config.AmiListenerEventMeetMeListComplete})
	return callback.SendSuperLevel()
}

// MeetMeMute mute a Meetme user.
func MeetMeMute(ctx context.Context, s AMISocket, meetme, userNumber string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMeetMeMute)
	c.SetV(map[string]interface{}{
		config.AmiFieldMeetMe:     meetme,
		config.AmiFieldUserNumber: userNumber,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MeetMeUnMute unmute a Meetme user.
func MeetMeUnMute(ctx context.Context, s AMISocket, meetme, userNumber string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMeetMeUnmute)
	c.SetV(map[string]interface{}{
		config.AmiFieldMeetMe:     meetme,
		config.AmiFieldUserNumber: userNumber,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// MeetMeListRooms list active conferences.
func MeetMeListRooms(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionMeetMeListRooms)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventMeetMeEntry}, []string{config.AmiListenerEventMeetMeListRoomsComplete})
	return callback.SendSuperLevel()
}
