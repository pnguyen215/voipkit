package ami

import (
	"context"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

// SKINNYDevices lists SKINNY devices (text format).
// Lists Skinny devices in text format with details on current status.
// Devicelist will follow as separate events,
// followed by a final event called DevicelistComplete.
func SKINNYDevices(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSKINNYdevices)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventDeviceEntry}, []string{config.AmiListenerEventDeviceListComplete})
	return callback.SendSuperLevel()
}

// SKINNYLines lists SKINNY lines (text format).
// Lists Skinny lines in text format with details on current status.
// Linelist will follow as separate events,
// followed by a final event called LinelistComplete.
func SKINNYLines(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSKINNYlines)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventLineEntry}, []string{config.AmiListenerEventLineListComplete})
	return callback.SendSuperLevel()
}

// SKINNYShowDevice show SKINNY device (text format).
// Show one SKINNY device with details on current status.
func SKINNYShowDevice(ctx context.Context, s AMISocket, device string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSKINNYShowDevice)
	c.SetV(map[string]interface{}{
		config.AmiFieldDevice: device,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// SKINNYShowline shows SKINNY line (text format).
// Show one SKINNY line with details on current status.
func SKINNYShowline(ctx context.Context, s AMISocket, line string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionSKINNYShowLine)
	c.SetV(map[string]interface{}{
		config.AmiFieldLine: line,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
