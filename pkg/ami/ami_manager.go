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
func Login(ctx context.Context, s AMISocket, auth *AMIAuth) error {
	if len(auth.Username) <= 0 {
		return fmt.Errorf(config.AmiErrorUsernameRequired)
	}

	if len(auth.Secret) <= 0 {
		return fmt.Errorf(config.AmiErrorPasswordRequired)
	}

	if len(auth.Events) == 0 {
		auth.SetEvent(config.AmiManagerPerm)
	}

	c := NewCommand()

	if len(s.UUID) <= 0 {
		uuid, err := GenUUID()
		if err != nil {
			return err
		}
		c.SetId(uuid)
		s.SetUUID(uuid)
	} else {
		c.SetId(s.UUID)
	}

	c.SetV(auth)
	c.SetAction(config.AmiActionLogin)
	response, err := c.Send(ctx, s, c)

	if len(response) == 0 {
		return fmt.Errorf(config.AmiErrorLoginFailed)
	}

	if err != nil {
		return fmt.Errorf(config.AmiErrorLoginFailedMessage, err.Error())
	}

	if IsFailure(response) {
		return fmt.Errorf(config.AmiErrorLoginFailedMessage, response.GetVal(config.AmiFieldMessage))
	}

	return nil
}

// Events gets events from current client connection
// It is mandatory set 'events' of ami.Login with "system,call,all,user", to received events.
func Events(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand()
	return c.Read(ctx, s)
}

// Logoff
// Logoff logoff the current manager session.
func Logoff(ctx context.Context, s AMISocket) error {
	c := NewCommand()
	c.SetAction(config.AmiActionLogoff)

	if len(s.UUID) <= 0 {
		_uuid, err := GenUUID()
		if err != nil {
			return err
		}
		s.SetUUID(_uuid)
	}

	c.SetId(s.UUID)
	response, err := c.Send(ctx, s, c)
	if err != nil {
		return err
	}
	log.Printf("Logoff, response = %v", utils.ToJson(response))
	return err
}

// Ping action will ellicit a 'Pong' response.
// Used to keep the manager connection open.
func Ping(ctx context.Context, socket AMISocket) error {
	c := NewCommand()
	c.SetAction(config.AmiActionPing)

	if len(socket.UUID) <= 0 {
		_uuid, err := GenUUID()
		if err != nil {
			return err
		}
		socket.SetUUID(_uuid)
	}

	c.SetId(socket.UUID)

	response, err := c.Send(ctx, socket, c)
	if err != nil {
		return err
	}

	log.Printf("Ping, response = %v", utils.ToJson(response))
	return err
}

// Command executes an Asterisk CLI Command.
func Command(ctx context.Context, s AMISocket, cmd string) (AMIResultRawLevel, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCommand)
	c.SetV(map[string]string{
		config.AmiActionCommand: cmd,
	})
	return c.SendLevel(ctx, s, c)
}

// CoreSettings shows PBX core settings (version etc).
func CoreSettings(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCoreSettings)
	return c.Send(ctx, s, c)
}

// CoreStatus shows PBX core status variables.
func CoreStatus(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCoreStatus)
	return c.Send(ctx, s, c)
}

// ListCommands lists available manager commands.
// Returns the action name and synopsis for every action that is available to the user
func ListCommands(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionListCommands)
	return c.Send(ctx, s, c)
}

// Challenge generates a challenge for MD5 authentication.
func Challenge(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionChallenge)
	c.SetV(map[string]string{
		config.AmiAuthTypeKey: "MD5",
	})
	return c.Send(ctx, s, c)
}

// CreateConfig creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory.
// This action is intended to be used before an UpdateConfig action.
func CreateConfig(ctx context.Context, s AMISocket, filename string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCreateConfig)
	c.SetV(map[string]string{
		config.AmiFilenameKey: filename,
	})
	return c.Send(ctx, s, c)
}

// DataGet retrieves the data api tree.
func DataGet(ctx context.Context, s AMISocket, path, search, filter string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDataGet)
	c.SetV(map[string]string{
		config.AmiFieldPath:   path,
		config.AmiFieldSearch: search,
		config.AmiFieldFilter: filter,
	})
	return c.Send(ctx, s, c)
}

// EventFlow control Event Flow.
// eventMask: Enable/Disable sending of events to this manager client.
func EventFlow(ctx context.Context, s AMISocket, eventMask string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionEvents)
	c.SetV(map[string]string{
		config.AmiFieldEventMask: eventMask,
	})
	return c.Send(ctx, s, c)
}

// GetConfig retrieves configuration.
// This action will dump the contents of a configuration file by category and contents or optionally by specified category only.
func GetConfig(ctx context.Context, s AMISocket, filename, category, filter string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionGetConfig)
	c.SetV(map[string]string{
		config.AmiFieldFilename: filename,
		config.AmiFieldFilter:   filter,
		config.AmiFieldCategory: category,
	})
	return c.Send(ctx, s, c)
}

// GetConfigJson retrieves configuration (JSON format).
// This action will dump the contents of a configuration file by category and contents in JSON format.
// This only makes sense to be used using raw man over the HTTP interface.
func GetConfigJson(ctx context.Context, s AMISocket, filename, category, filter string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionGetConfigJson)
	c.SetV(map[string]string{
		config.AmiFieldFilename: filename,
		config.AmiFieldFilter:   filter,
		config.AmiFieldCategory: category,
	})
	return c.Send(ctx, s, c)
}

// JabberSend sends a message to a Jabber Client
func JabberSend(ctx context.Context, s AMISocket, jabber, jid, message string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionJabberSend)
	c.SetV(map[string]string{
		config.AmiFieldJabber:  jabber,
		config.AmiFieldJID:     jid,
		config.AmiFieldMessage: message,
	})
	return c.Send(ctx, s, c)
}
