package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// PRIDebugFileSet set the file used for PRI debug message output.
func PRIDebugFileSet(ctx context.Context, s AMISocket, filename string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPRIDebugFileSet)
	c.SetV(map[string]interface{}{
		config.AmiFieldFile: filename,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PRIDebugFileUnset disables file output for PRI debug messages.
func PRIDebugFileUnset(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPRIDebugFileUnset)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PRIDebugSet set PRI debug levels for a span.
func PRIDebugSet(ctx context.Context, s AMISocket, span, level string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPRIDebugSet)
	c.SetV(map[string]interface{}{
		config.AmiFieldSpan:  span,
		config.AmiFieldLevel: level,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PRIShowSpans show status of PRI spans.
func PRIShowSpans(ctx context.Context, s AMISocket, span string) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPRIShowSpans)
	c.SetV(map[string]interface{}{
		config.AmiFieldSpan: span,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{config.AmiActionPRIShowSpans}, []string{config.AmiListenerEventPRIShowSpansComplete})
	return callback.SendSuperLevel()
}
