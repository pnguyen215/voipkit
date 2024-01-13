package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// PJSIPNotify send NOTIFY to either an endpoint, an arbitrary URI, or inside a SIP dialog.
func PJSIPNotify(ctx context.Context, s AMISocket, endpoint, uri, variable string) (AmiReply, error) {
	params := map[string]string{
		config.AmiFieldVariable: variable,
	}

	if len(endpoint) > 0 {
		params[config.AmiFieldEndpoint] = endpoint
	}

	if len(uri) > 0 {
		params[config.AmiFieldUri] = uri
	}

	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPNotify)
	c.SetV(params)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PJSIPQualify qualify a chan_pjsip endpoint.
func PJSIPQualify(ctx context.Context, s AMISocket, endpoint string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPQualify)
	c.SetV(map[string]interface{}{
		config.AmiFieldEndpoint: endpoint,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PJSIPRegister register an outbound registration.
func PJSIPRegister(ctx context.Context, s AMISocket, registration string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPRegister)
	c.SetVCmd(map[string]interface{}{
		config.AmiFieldRegistration: registration,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PJSIPUnregister unregister an outbound registration.
func PJSIPUnregister(ctx context.Context, s AMISocket, registration string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPUnregister)
	c.SetVCmd(map[string]interface{}{
		config.AmiFieldRegistration: registration,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// PJSIPShowEndpoint detail listing of an endpoint and its objects.
func PJSIPShowEndpoint(ctx context.Context, s AMISocket, endpoint string) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowEndpoint)
	c.SetV(map[string]interface{}{
		config.AmiFieldEndpoint: endpoint,
	})
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventEndpointDetail, config.AmiListenerEventContactStatusDetail,
			config.AmiListenerEventAorDetail, config.AmiListenerEventAuthDetail,
			config.AmiListenerEventTransportDetail, config.AmiListenerEventIdentifyDetail,
		},
		[]string{config.AmiListenerEventEndpointDetailComplete})
	return callback.SendSuperLevel()
}

// PJSIPShowEndpoints list pjsip endpoints.
func PJSIPShowEndpoints(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowEndpoints)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventEndpointList},
		[]string{config.AmiListenerEventEndpointListComplete})
	return callback.SendSuperLevel()
}

// PJSIPShowRegistrationInboundContactStatuses lists ContactStatuses for PJSIP inbound registrations.
func PJSIPShowRegistrationInboundContactStatuses(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowRegistrationInboundContactStatuses)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventContactStatusDetail},
		[]string{config.AmiListenerEventContactStatusDetailComplete})
	return callback.SendSuperLevel()
}

// PJSIPShowRegistrationsInbound lists PJSIP inbound registrations.
func PJSIPShowRegistrationsInbound(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowRegistrationsInbound)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventInboundRegistrationDetail},
		[]string{config.AmiListenerEventInboundRegistrationDetailComplete})
	return callback.SendSuperLevel()
}

// PJSIPShowRegistrationsOutbound lists PJSIP outbound registrations.
func PJSIPShowRegistrationsOutbound(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowRegistrationsOutbound)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventOutboundRegistrationDetail},
		[]string{config.AmiListenerEventOutboundRegistrationDetailComplete})
	return callback.SendSuperLevel()
}

// PJSIPShowResourceLists displays settings for configured resource lists.
func PJSIPShowResourceLists(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowResourceLists)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventResourceListDetail},
		[]string{config.AmiListenerEventResourceListDetailComplete})
	return callback.SendSuperLevel()
}

// PJSIPShowSubscriptionsInbound list of inbound subscriptions.
func PJSIPShowSubscriptionsInbound(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowSubscriptionsInbound)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventInboundSubscriptionDetail},
		[]string{config.AmiListenerEventInboundSubscriptionDetailComplete})
	return callback.SendSuperLevel()
}

// PJSIPShowSubscriptionsOutbound list of outbound subscriptions.
func PJSIPShowSubscriptionsOutbound(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionPJSIPShowSubscriptionsOutbound)
	callback := NewAmiCallbackService(ctx, s, c,
		[]string{config.AmiListenerEventOutboundSubscriptionDetail},
		[]string{config.AmiListenerEventOutboundSubscriptionDetailComplete})
	return callback.SendSuperLevel()
}
