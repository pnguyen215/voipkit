package ami

import (
	"context"
	"fmt"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
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

// Login provides the login manager.
func Login(ctx context.Context, socket AMISocket, auth *AMIAuth) error {
	if len(auth.Username) <= 0 {
		return fmt.Errorf("(Ami Auth). username was missing")
	}

	if len(auth.Secret) <= 0 {
		return fmt.Errorf("(Ami Auth). secret was missing")
	}

	if len(auth.Events) == 0 {
		auth.Events = config.AmiManagerPerm
	}

	a := NewCommand()

	if len(auth.ID) <= 0 {
		uuid, err := GenUUID()
		if err != nil {
			return err
		}
		a.SetId(uuid)
	} else {
		a.SetId(auth.ID)
	}

	a.SetV(auth)
	a.SetAction(config.AmiActionLogin)
	response, err := a.Send(ctx, socket, a)

	if err != nil {
		return err
	}

	if !IsSuccess(response) {
		return fmt.Errorf("(Ami Auth). login failed with reason is >> %v", response.GetVal(config.AmiFieldMessage))
	}

	return nil
}

// Events gets events from current client connection
// It is mandatory set 'events' of ami.Login with "system,call,all,user", to received events.
func Events(ctx context.Context, socket AMISocket) (AMIResultRaw, error) {
	a := NewCommand()
	return a.Read(ctx, socket)
}

// Logoff logoff the current manager session.
func Logoff(ctx context.Context, socket AMISocket, uuid string) error {
	a := NewCommand()
	a.SetAction(config.AmiActionLogoff)

	if len(uuid) <= 0 {
		_uuid, err := GenUUID()
		if err != nil {
			return err
		}
		uuid = _uuid
	}
	a.SetId(uuid)
	response, err := a.Send(ctx, socket, a)
	if err != nil {
		return err
	}

	if msg := response.GetVal(config.AmiResponseKey); msg != config.AmiFieldGoodbye {
		return fmt.Errorf("Ami Logout failed: %v", response.GetVal(config.AmiFieldMessage))
	}

	return nil
}
