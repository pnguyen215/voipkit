package ami

import (
	"context"
	"strings"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

type AmiCallbackService interface {
	Send() (AmiReply, error)
	SendLevel() (AmiReplies, error)
	SendSuperLevel() ([]AmiReply, error)
}

func NewAmiCallbackService(ctx context.Context, socket AMISocket, command *AMICommand,
	acceptedEvents []string, ignoreEvents []string) AmiCallbackService {
	a := NewAmiCallbackHandler()
	a.SetContext(ctx)
	a.SetSocket(socket)
	a.SetCommand(command)
	a.SetAcceptedEvents(acceptedEvents)
	a.SetIgnoredEvents(ignoreEvents)
	return a
}

func NewAmiCallbackHandlerService(handler *AMICallbackHandler) AmiCallbackService {
	return handler
}

func NewAmiCallbackHandler() *AMICallbackHandler {
	a := &AMICallbackHandler{}
	return a
}

func (a *AMICallbackHandler) SetContext(value context.Context) *AMICallbackHandler {
	a.ctx = value
	return a
}

func (a *AMICallbackHandler) SetSocket(value AMISocket) *AMICallbackHandler {
	a.socket = value
	return a
}

func (a *AMICallbackHandler) SetCommand(value *AMICommand) *AMICallbackHandler {
	a.command = value
	return a
}

func (a *AMICallbackHandler) SetAcceptedEvents(values []string) *AMICallbackHandler {
	a.AcceptedEvents = values
	return a
}

func (a *AMICallbackHandler) AppendAcceptedEvents(values ...string) *AMICallbackHandler {
	a.AcceptedEvents = append(a.AcceptedEvents, values...)
	return a
}

func (a *AMICallbackHandler) SetIgnoredEvents(values []string) *AMICallbackHandler {
	a.IgnoreEvents = values
	return a
}

func (a *AMICallbackHandler) AppendIgnoredEvents(values ...string) *AMICallbackHandler {
	a.IgnoreEvents = append(a.IgnoreEvents, values...)
	return a
}

func (a *AMICallbackHandler) Json() string {
	return JsonString(a)
}

func (h *AMICallbackHandler) Send() (AmiReply, error) {
	if !h.socket.Retry {
		return h.command.Send(h.ctx, h.socket, h.command)
	}

	if h.socket.MaxRetries == 1 || h.socket.MaxRetries <= 0 {
		return h.command.Send(h.ctx, h.socket, h.command)
	}

	var response AmiReply
	var err error
	var total time.Duration = 0

	for i := 1; i <= h.socket.MaxRetries; i++ {
		_start := time.Now()
		response, err = h.command.Send(h.ctx, h.socket, h.command)
		_end := time.Since(_start)
		total += _end
		if _end == 0 || strings.EqualFold(response.Get(config.AmiJsonFieldStatus), config.AmiFullyBootedKey) {
			continue
		}
		if len(response) > 0 && err == nil {
			if h.socket.DebugMode {
				D().Info("Send(). callback return for the %v time(s) and waste time: %v", i, _end)
			}
			break
		}
	}
	if h.socket.DebugMode {
		D().Info("Send(). callback total waste time: %v", total)
	}
	return response, err
}

func (h *AMICallbackHandler) SendLevel() (AmiReplies, error) {
	if !h.socket.Retry {
		return h.command.SendLevel(h.ctx, h.socket, h.command)
	}

	if h.socket.MaxRetries == 1 || h.socket.MaxRetries <= 0 {
		return h.command.SendLevel(h.ctx, h.socket, h.command)
	}

	var response AmiReplies
	var err error
	var total time.Duration = 0

	for i := 1; i <= h.socket.MaxRetries; i++ {
		_start := time.Now()
		response, err = h.command.SendLevel(h.ctx, h.socket, h.command)
		_end := time.Since(_start)
		total += _end
		if _end == 0 || strings.EqualFold(response.Get(config.AmiJsonFieldStatus), config.AmiFullyBootedKey) {
			continue
		}
		if len(response) > 0 && err == nil {
			if h.socket.DebugMode {
				D().Info("SendLevel(). callback return for the %v time(s) and waste time: %v", i, _end)
			}
			break
		}
	}
	if h.socket.DebugMode {
		D().Info("SendLevel(). callback total waste time: %v", total)
	}
	return response, err
}

func (h *AMICallbackHandler) SendSuperLevel() ([]AmiReply, error) {
	if !h.socket.Retry {
		return h.command.DoGetResult(h.ctx, h.socket, h.command, h.AcceptedEvents, h.IgnoreEvents)
	}

	if h.socket.MaxRetries == 1 || h.socket.MaxRetries <= 0 {
		return h.command.DoGetResult(h.ctx, h.socket, h.command, h.AcceptedEvents, h.IgnoreEvents)
	}

	var response []AmiReply
	var err error
	var total time.Duration = 0

	for i := 1; i <= h.socket.MaxRetries; i++ {
		_start := time.Now()
		response, err = h.command.DoGetResult(h.ctx, h.socket, h.command, h.AcceptedEvents, h.IgnoreEvents)
		_end := time.Since(_start)
		total += _end
		if _end == 0 {
			continue
		}
		if len(response) > 0 && err == nil {
			if h.socket.DebugMode {
				D().Info("SendSuperLevel(). callback return for the %v time(s) and waste time: %v", i, _end)
			}
			break
		}
	}
	if h.socket.DebugMode {
		D().Info("SendSuperLevel(). callback total waste time: %v", total)
	}
	return response, err
}
