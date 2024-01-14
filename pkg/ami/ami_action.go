package ami

import (
	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAction() *AMIAction {
	a := &AMIAction{}
	return a
}

func NewRevoke(cmd string, timeout int) *AMIAction {
	cli := NewAction()
	cli.Name = cmd
	cli.Timeout = timeout
	return cli
}

func (a *AMIAction) SetName(action string) *AMIAction {
	a.Name = action
	return a
}

func (a *AMIAction) SetTimeout(timeout int) *AMIAction {
	if timeout >= config.AmiMinTimeoutInMsForCall && timeout <= config.AmiMaxTimeoutInMsForCall {
		a.Timeout = timeout
	}
	return a
}

// func (c *AMIAction) Revoke(a *AMI, d *AMIDictionary, e *AMIMessage, deadlock bool) (*AMIResponse, error) {
// 	D().Info("Ami revoking action (state mutex opened lock~unlock): '%v'", e.String())
// 	var response AMIResponse
// 	var _err error
// 	if IsStringEmpty(c.Name) {
// 		response.Message = fmt.Sprintf(config.AmiErrorFieldRequired, "name")
// 		response.IsSuccess = false
// 		_err = fmt.Errorf(response.Message)
// 		return &response, _err
// 	}
// 	a.Action(e)
// 	all := a.AllEvents()
// 	if deadlock {
// 		defer a.Close()
// 	}
// 	for {
// 		select {
// 		case message := <-all:
// 			message.SetTimeFormat(e.TimeFormat)
// 			message.SetPhonePrefix(e.PhonePrefix)
// 			message.SetRegion(e.Region)
// 			message.AddFieldDateReceivedAt()
// 			if message.IsResponse() {
// 				response.event = message
// 				response.IsSuccess = true
// 				response.Json = message.JsonTranslator(d)
// 				goto on_success
// 			}
// 		case err := <-a.Error():
// 			a.Close()
// 			_err = err
// 			response.event = nil
// 			response.IsSuccess = false
// 			response.Message = _err.Error()
// 			goto on_failed
// 		}
// 	}

// on_success:
// 	return &response, nil
// on_failed:
// 	return &response, _err
// }

// func (c *AMIAction) Run(a *AMI) (*AMIResponse, error) {
// 	action := WithMessage(config.AmiActionCommand)
// 	action.AddField(config.AmiActionCommand, c.Name)
// 	return c.Revoke(a, NewDictionary(), action, false)
// }

// func (c *AMIAction) WithRunX(a *AMI, dictionaries map[string]string) (*AMIResponse, error) {
// 	action := WithMessage(config.AmiActionCommand)
// 	action.AddField(config.AmiActionCommand, c.Name)
// 	d := NewDictionary()
// 	d.AddKeysTranslator(dictionaries)
// 	return c.Revoke(a, d, action, false)
// }

// func (c *AMIAction) WithRunV(a *AMI, command map[string]string) (*AMIResponse, error) {
// 	action := WithMessage(c.Name)
// 	action.AddFields(command)
// 	return c.Revoke(a, NewDictionary(), action, false)
// }

// func (c *AMIAction) WithRunXV(a *AMI, command, dictionaries map[string]string) (*AMIResponse, error) {
// 	action := WithMessage(c.Name)
// 	action.AddFields(command)
// 	d := NewDictionary()
// 	d.AddKeysTranslator(dictionaries)
// 	return c.Revoke(a, d, action, false)
// }
