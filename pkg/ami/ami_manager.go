package ami

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewAuth() *AMIAuth {
	a := &AMIAuth{}
	return a
}

func (a *AMIAuth) SetUsername(username string) *AMIAuth {
	a.Username = username
	return a
}

func (a *AMIAuth) SetSecret(password string) *AMIAuth {
	a.Secret = password
	return a
}

func (a *AMIAuth) SetPassword(password string) *AMIAuth {
	a.SetSecret(password)
	return a
}

func (a *AMIAuth) SetEvent(event string) *AMIAuth {
	a.Events = event
	return a
}

func (a *AMIAuth) SetEvents(events ...string) *AMIAuth {
	_e := strings.Join(events, ",")
	a.SetEvent(_e)
	return a
}

// Login
// Login provides the login manager.
func Login(ctx context.Context, socket AMISocket, auth *AMIAuth) error {
	if len(auth.Username) <= 0 {
		return fmt.Errorf(config.AmiErrorUsernameRequired)
	}

	if len(auth.Secret) <= 0 {
		return fmt.Errorf(config.AmiErrorPasswordRequired)
	}

	if len(auth.Events) == 0 {
		auth.SetEvent(config.AmiManagerPerm)
	}

	a := NewCommand()

	if len(socket.UUID) <= 0 {
		uuid, err := GenUUID()
		if err != nil {
			return err
		}
		a.SetId(uuid)
		socket.SetUUID(uuid)
	} else {
		a.SetId(socket.UUID)
	}

	a.SetV(auth)
	a.SetAction(config.AmiActionLogin)
	response, err := a.Send(ctx, socket, a)

	if err != nil {
		return err
	}

	if !IsSuccess(response) {
		return fmt.Errorf(config.AmiErrorLoginFailedMessage, response.GetVal(config.AmiFieldMessage))
	}

	return nil
}

// Events gets events from current client connection
// It is mandatory set 'events' of ami.Login with "system,call,all,user", to received events.
func Events(ctx context.Context, socket AMISocket) (AMIResultRaw, error) {
	a := NewCommand()
	return a.Read(ctx, socket)
}

// Logoff
// Logoff logoff the current manager session.
func Logoff(ctx context.Context, socket AMISocket) error {
	a := NewCommand()
	a.SetAction(config.AmiActionLogoff)

	if len(socket.UUID) <= 0 {
		_uuid, err := GenUUID()
		if err != nil {
			return err
		}
		socket.SetUUID(_uuid)
	}

	a.SetId(socket.UUID)
	response, err := a.Send(ctx, socket, a)
	if err != nil {
		return err
	}

	if msg := response.GetVal(config.AmiResponseKey); msg != config.AmiFieldGoodbye {
		log.Printf("Logoff, response failed = %v", utils.ToJson(response))
		return fmt.Errorf(config.AmiErrorLogoutFailedMessage, response.GetVal(config.AmiFieldMessage))
	}

	return nil
}
