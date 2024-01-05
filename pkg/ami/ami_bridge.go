package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// Bridge bridges two channels already in the PBX.
func Bridge(ctx context.Context, s AMISocket, channel1, channel2 string, tone string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridge)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel1: channel1,
		config.AmiFieldChannel2: channel2,
		config.AmiFieldTone:     tone,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// BlindTransfer blind transfer channel(s) to the given destination.
func BlindTransfer(ctx context.Context, s AMISocket, channel, context, extension string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBlindTransfer)
	c.SetV(map[string]interface{}{
		config.AmiFieldChannel:   channel,
		config.AmiFieldContext:   context,
		config.AmiFieldExtension: extension,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// BridgeDestroy destroy a bridge.
func BridgeDestroy(ctx context.Context, s AMISocket, bridgeUniqueId string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridgeDestroy)
	c.SetV(map[string]interface{}{
		config.AmiFieldBridgeUniqueId: bridgeUniqueId,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// BridgeInfo get information about a bridge.
func BridgeInfo(ctx context.Context, s AMISocket, bridgeUniqueId string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridgeInfo)
	c.SetV(map[string]interface{}{
		config.AmiFieldBridgeUniqueId: bridgeUniqueId,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// BridgeKick kick a channel from a bridge.
func BridgeKick(ctx context.Context, s AMISocket, bridgeUniqueId, channel string) (AMIResultRaw, error) {
	params := map[string]string{
		config.AmiFieldChannel: channel,
	}
	if len(bridgeUniqueId) > 0 {
		params[config.AmiFieldBridgeUniqueId] = bridgeUniqueId
	}
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridgeKick)
	c.SetV(params)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// BridgeList get a list of bridges in the system.
func BridgeList(ctx context.Context, s AMISocket, bridgeType string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridgeList)
	c.SetV(map[string]interface{}{
		config.AmiFieldBridgeType: bridgeType,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// BridgeTechnologyList list available bridging technologies and their statuses.
func BridgeTechnologyList(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridgeTechnologyList)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventBridgeTechnologyListItem}, []string{config.AmiListenerEventBridgeTechnologyListComplete})
	return callback.SendSuperLevel()
}

// BridgeTechnologySuspend suspend a bridging technology.
func BridgeTechnologySuspend(ctx context.Context, s AMISocket, bridgeTechnology string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridgeTechnologySuspend)
	c.SetV(map[string]interface{}{
		config.AmiFieldBridgeTechnology: bridgeTechnology,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// BridgeTechnologyUnsuspend unsuspend a bridging technology.
func BridgeTechnologyUnsuspend(ctx context.Context, s AMISocket, bridgeTechnology string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionBridgeTechnologyUnsuspend)
	c.SetV(map[string]interface{}{
		config.AmiFieldBridgeTechnology: bridgeTechnology,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
