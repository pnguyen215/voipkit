package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// QueueStatuses show status all members in queue.
func QueueStatuses(ctx context.Context, s AMISocket, queue string) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueStatus)
	c.SetV(map[string]string{
		config.AmiFieldQueue: queue,
	})
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventQueueMember, config.AmiListenerEventQueueEntry},
		[]string{config.AmiListenerEventQueueStatusComplete})
	return callback.SendSuperLevel()
}

// QueueSummary show queue summary.
func QueueSummary(ctx context.Context, s AMISocket, queue string) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueSummary)
	c.SetV(map[string]string{
		config.AmiFieldQueue: queue,
	})
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventQueueSummary},
		[]string{config.AmiListenerEventQueueSummaryComplete})
	return callback.SendSuperLevel()
}

// QueueMemberRingInUse set the ringinuse value for a queue member.
func QueueMemberRingInUse(ctx context.Context, s AMISocket, _interface, ringInUse, queue string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueMemberRingInUse)
	c.SetV(map[string]string{
		config.AmiFieldInterface: _interface,
		config.AmiFieldRingInUse: ringInUse,
		config.AmiFieldQueue:     queue,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueueStatus show queue status by member.
func QueueStatus(ctx context.Context, s AMISocket, queue, member string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueStatus)
	c.SetV(map[string]string{
		config.AmiFieldQueue:  queue,
		config.AmiFieldMember: member,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueueRule queues Rules.
func QueueRule(ctx context.Context, s AMISocket, rule string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueRule)
	c.SetV(map[string]string{
		config.AmiFieldRule: rule,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueueReset resets queue statistics.
func QueueReset(ctx context.Context, s AMISocket, queue string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueReset)
	c.SetVCmd(AMIPayloadQueue{Queue: queue})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueueRemove removes interface from queue.
func QueueRemove(ctx context.Context, s AMISocket, queue AMIPayloadQueue) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueRemove)
	c.SetVCmd(queue)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueueReload reloads a queue, queues, or any sub-section of a queue or queues.
func QueueReload(ctx context.Context, s AMISocket, queue AMIPayloadQueue) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueReload)
	c.SetVCmd(queue)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueuePenalty sets the penalty for a queue member.
func QueuePenalty(ctx context.Context, s AMISocket, queue AMIPayloadQueue) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueuePenalty)
	c.SetVCmd(queue)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueuePause makes a queue member temporarily unavailable.
func QueuePause(ctx context.Context, s AMISocket, queue AMIPayloadQueue) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueuePause)
	c.SetVCmd(queue)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueueLog adds custom entry in queue_log.
func QueueLog(ctx context.Context, s AMISocket, queue AMIPayloadQueue) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueLog)
	c.SetVCmd(queue)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// QueueAdd adds interface to queue.
func QueueAdd(ctx context.Context, s AMISocket, queue AMIPayloadQueue) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionQueueAdd)
	c.SetVCmd(queue)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
