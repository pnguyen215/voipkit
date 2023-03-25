package ami

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

type AMICallbackService interface {
	Send() (AMIResultRaw, error)
	SendLevel() (AMIResultRawLevel, error)
	SendSuperLevel() ([]AMIResultRaw, error)
}

func NewAMICallbackService(ctx context.Context, socket AMISocket, command *AMICommand,
	acceptedEvents []string, ignoreEvents []string) AMICallbackService {
	a := NewAMICallbackHandler()
	a.SetContext(ctx)
	a.SetSocket(socket)
	a.SetCommand(command)
	a.SetAcceptedEvents(acceptedEvents)
	a.SetIgnoredEvents(ignoreEvents)
	return a
}

func NewAMICallbackServiceWith(handler *AMICallbackHandler) AMICallbackService {
	return handler
}

func NewAMICallbackHandler() *AMICallbackHandler {
	a := &AMICallbackHandler{}
	return a
}

func (a *AMICallbackHandler) SetContext(value context.Context) *AMICallbackHandler {
	a.Ctx = value
	return a
}

func (a *AMICallbackHandler) SetSocket(value AMISocket) *AMICallbackHandler {
	a.Socket = value
	return a
}

func (a *AMICallbackHandler) SetCommand(value *AMICommand) *AMICallbackHandler {
	a.Command = value
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
	return utils.ToJson(a)
}

func (h *AMICallbackHandler) Send() (AMIResultRaw, error) {
	if !h.Socket.Retry {
		return h.Command.Send(h.Ctx, h.Socket, h.Command)
	}

	if h.Socket.MaxRetries == 1 || h.Socket.MaxRetries <= 0 {
		return h.Command.Send(h.Ctx, h.Socket, h.Command)
	}

	var response AMIResultRaw
	var err error
	var total int64 = 0

	for i := 1; i <= h.Socket.MaxRetries; i++ {
		_start := time.Now().UnixMilli()
		response, err = h.Command.Send(h.Ctx, h.Socket, h.Command)
		_end := time.Now().UnixMilli() - _start
		total += _end
		if _end == 0 || strings.EqualFold(response.GetVal(config.AmiJsonFieldStatus), config.AmiFullyBootedKey) {
			continue
		}

		if len(response) > 0 && err == nil {
			if h.Socket.AllowTrace {
				log.Printf("Send(). callback return for the %v time(s) and waste time = %v (milliseconds)", i, _end)
			}
			break
		}
	}

	if h.Socket.AllowTrace {
		log.Printf("Send(). callback total waste time = %v (milliseconds)", total)
	}

	return response, err
}

func (h *AMICallbackHandler) SendLevel() (AMIResultRawLevel, error) {
	if !h.Socket.Retry {
		return h.Command.SendLevel(h.Ctx, h.Socket, h.Command)
	}

	if h.Socket.MaxRetries == 1 || h.Socket.MaxRetries <= 0 {
		return h.Command.SendLevel(h.Ctx, h.Socket, h.Command)
	}

	var response AMIResultRawLevel
	var err error
	var total int64 = 0

	for i := 1; i <= h.Socket.MaxRetries; i++ {
		_start := time.Now().UnixMilli()
		response, err = h.Command.SendLevel(h.Ctx, h.Socket, h.Command)
		_end := time.Now().UnixMilli() - _start
		total += _end
		if _end == 0 || strings.EqualFold(response.GetVal(config.AmiJsonFieldStatus), config.AmiFullyBootedKey) {
			continue
		}

		if len(response) > 0 && err == nil {
			if h.Socket.AllowTrace {
				log.Printf("SendLevel(). callback return for the %v time(s) and waste time = %v (milliseconds)", i, _end)
			}
			break
		}
	}

	if h.Socket.AllowTrace {
		log.Printf("SendLevel(). callback total waste time = %v (milliseconds)", total)
	}

	return response, err
}

func (h *AMICallbackHandler) SendSuperLevel() ([]AMIResultRaw, error) {
	if !h.Socket.Retry {
		return h.Command.DoGetResult(h.Ctx, h.Socket, h.Command, h.AcceptedEvents, h.IgnoreEvents)
	}

	if h.Socket.MaxRetries == 1 || h.Socket.MaxRetries <= 0 {
		return h.Command.DoGetResult(h.Ctx, h.Socket, h.Command, h.AcceptedEvents, h.IgnoreEvents)
	}

	var response []AMIResultRaw
	var err error
	var total int64 = 0

	for i := 1; i <= h.Socket.MaxRetries; i++ {
		_start := time.Now().UnixMilli()
		response, err = h.Command.DoGetResult(h.Ctx, h.Socket, h.Command, h.AcceptedEvents, h.IgnoreEvents)
		_end := time.Now().UnixMilli() - _start
		total += _end
		if _end == 0 {
			continue
		}

		if len(response) > 0 && err == nil {
			if h.Socket.AllowTrace {
				log.Printf("SendSuperLevel(). callback return for the %v time(s) and waste time = %v (milliseconds)", i, _end)
			}
			break
		}
	}

	if h.Socket.AllowTrace {
		log.Printf("SendSuperLevel(). callback total waste time = %v (milliseconds)", total)
	}

	return response, err
}
