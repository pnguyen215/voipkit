package ami

import (
	"context"
	"fmt"
	"log"
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
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	response, err := callback.Send()

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
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	response, err := callback.Send()
	if err != nil {
		return err
	}
	log.Printf("Logoff, response = %v", JsonString(response))
	return err
}

// Ping action will ellicit a 'Pong' response.
// Used to keep the manager connection open.
func Ping(ctx context.Context, s AMISocket) error {
	c := NewCommand()
	c.SetAction(config.AmiActionPing)

	if len(s.UUID) <= 0 {
		_uuid, err := GenUUID()
		if err != nil {
			return err
		}
		s.SetUUID(_uuid)
	}

	c.SetId(s.UUID)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	response, err := callback.Send()
	if err != nil {
		return err
	}

	log.Printf("Ping, response = %v", JsonString(response))
	return err
}

// Command executes an Asterisk CLI Command.
func Command(ctx context.Context, s AMISocket, cmd string) (AMIResultRawLevel, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCommand)
	c.SetV(map[string]string{
		config.AmiActionCommand: cmd,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.SendLevel()
}

// CoreSettings shows PBX core settings (version etc).
func CoreSettings(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCoreSettings)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// CoreStatus shows PBX core status variables.
func CoreStatus(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCoreStatus)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ListCommands lists available manager commands.
// Returns the action name and synopsis for every action that is available to the user
func ListCommands(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionListCommands)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Challenge generates a challenge for MD5 authentication.
func Challenge(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionChallenge)
	c.SetV(map[string]string{
		config.AmiAuthTypeKey: "MD5",
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// CreateConfig creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory.
// This action is intended to be used before an UpdateConfig action.
func CreateConfig(ctx context.Context, s AMISocket, filename string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionCreateConfig)
	c.SetV(map[string]string{
		config.AmiFilenameKey: filename,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DataGet retrieves the data api tree.
func DataGet(ctx context.Context, s AMISocket, path, search, filter string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDataGet)
	c.SetV(map[string]string{
		config.AmiFieldPath:   path,
		config.AmiFieldSearch: search,
		config.AmiFieldFilter: filter,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// EventFlow control Event Flow.
// eventMask: Enable/Disable sending of events to this manager client.
func EventFlow(ctx context.Context, s AMISocket, eventMask string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionEvents)
	c.SetV(map[string]string{
		config.AmiFieldEventMask: eventMask,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
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
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
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
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// JabberSend sends a message to a Jabber Client
func JabberSend(ctx context.Context, s AMISocket, jabber, jid, message string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionJabberSend)
	c.SetV(map[string]string{
		config.AmiFieldJabber:  jabber,
		config.AmiFieldJID:     jid,
		config.AmiFieldMessage: message,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ListCategories lists categories in configuration file.
func ListCategories(ctx context.Context, s AMISocket, filename string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionListCategories)
	c.SetV(map[string]string{
		config.AmiFieldFilename: filename,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ModuleCheck checks if module is loaded.
// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
func ModuleCheck(ctx context.Context, s AMISocket, module string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionModuleCheck)
	c.SetV(map[string]string{
		config.AmiFieldModule: module,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ModuleLoad module management.
// Loads, unloads or reloads an Asterisk module in a running system.
func ModuleLoad(ctx context.Context, s AMISocket, module, loadType string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionModuleLoad)
	c.SetV(map[string]string{
		config.AmiFieldModule:   module,
		config.AmiFieldLoadType: loadType,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// Reload Sends a reload event.
func Reload(ctx context.Context, s AMISocket, module string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionReload)
	c.SetV(map[string]string{
		config.AmiFieldModule: module,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// ShowDialPlan shows dialplan contexts and extensions
// Be aware that showing the full dialplan may take a lot of capacity.
func ShowDialPlan(ctx context.Context, s AMISocket, extension, context string) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionShowDialPlan)
	c.SetV(map[string]string{
		config.AmiFieldExtension_: extension,
		config.AmiFieldContext:    context,
	})
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventListDialplan}, []string{config.AmiListenerEventShowDialPlanComplete})
	return callback.SendSuperLevel()
}

// Filter dynamically add filters for the current manager session.
func Filter(ctx context.Context, s AMISocket, operation, filter string) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionFilter)
	c.SetV(map[string]string{
		config.AmiFieldOperation: operation,
		config.AmiFieldFilter:    filter,
	})
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DeviceStateList list the current known device states.
func DeviceStateList(ctx context.Context, s AMISocket) ([]AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDeviceStateList)
	callback := NewAMICallbackService(ctx, s, c,
		[]string{config.AmiListenerEventDeviceStateChange}, []string{config.AmiListenerEventDeviceStateListComplete})
	return callback.SendSuperLevel()
}

// LoggerRotate reload and rotate the Asterisk logger.
func LoggerRotate(ctx context.Context, s AMISocket) (AMIResultRaw, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionLoggerRotate)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// UpdateConfig Updates a config file.
// Dynamically updates an Asterisk configuration file.
/*
Action: UpdateConfig
SrcFilename: voicemail2.conf
DstFilename: voicemail2.conf
Action-000000: Append
Cat-000000: default
Var-000000: 127
Value-000000: >5555, Jason Bourne97, ***@noCia.gov.do
Action-000001: Append
Cat-000001: default
Var-000001: 125
Value-000001: >55555, Jason Bourne76, ***@noCia.gov.do
Action-000002: Append
Cat-000002: default
Var-000002: 122
Value-000002: >5555, Jason Bourne74, ***@noCia.gov.do
Action-000003: Append
Cat-000003: default
Var-000003: 128
Value-000003: >5555, Jason Bourne48, ***@noCia.gov.do
Action-000004: Append
Cat-000004: default
Var-000004: 126
Value-000004: >55555, Jason Bourne18, ***@noCia.gov.do
ActionID: 495446608
*/
func UpdateConfig(ctx context.Context, s AMISocket, sourceFilename, destinationFilename string, reload bool, actions ...AMIUpdateConfigAction) (AMIResultRaw, error) {
	options := make(map[string]string)
	options[config.AmiFieldSourceFilename] = sourceFilename
	options[config.AmiFieldDestinationFilename] = destinationFilename
	if reload {
		options[config.AmiFieldReload] = "yes"
	}
	for i, a := range actions {
		uuid := fmt.Sprintf("%06d", i)
		options[fmt.Sprintf("%s%s", config.AmiFieldActionPrefix, uuid)] = a.Action
		options[fmt.Sprintf("%s%s", config.AmiFieldCategoryPrefix, uuid)] = a.Category
		if a.Var != "" {
			options[fmt.Sprintf("%s%s", config.AmiFieldVarPrefix, uuid)] = a.Var
			options[fmt.Sprintf("%s%s", config.AmiFieldValuePrefix, uuid)] = a.Value
		}
	}

	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionUpdateConfig)
	c.SetVCmd(options)
	callback := NewAMICallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
