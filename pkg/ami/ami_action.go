package ami

import (
	"fmt"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAction() *AMIAction {
	a := &AMIAction{}
	return a
}

func NewRevokeAction(cmd string, timeout int) *AMIAction {
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

// Revoke run cli on asterisk server
func (c *AMIAction) Revoke(a *AMI, d *AMIDictionary, e *AMIMessage, deadlock bool) (*AMIResponse, error) {
	D().Info("[>] Ami revoke action (state mutex opened lock~unlock) >>> '%v'", e.String())
	var response AMIResponse
	var _err error

	if strings.EqualFold(c.Name, "") {
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "name")
		response.IsSuccess = false
		_err = fmt.Errorf(response.ErrorMessage)
		return &response, _err
	}

	a.Action(e)
	all := a.AllEvents()

	if deadlock {
		defer a.Close()
	}

	for {
		select {
		case message := <-all:
			message.SetDateTimeLayout(e.DateTimeLayout)
			message.SetPhonePrefix(e.PhonePrefix)
			message.SetRegion(e.Region)
			message.AddFieldDateReceivedAt()
			if message.IsResponse() {
				response.Event = message
				response.IsSuccess = true
				response.RawJson = message.JsonTranslator(d)
				goto on_success
			}
		case err := <-a.Error():
			a.Close()
			_err = err
			response.Event = nil
			response.IsSuccess = false
			response.ErrorMessage = _err.Error()
			goto on_failed
		}
	}

on_success:
	return &response, nil
on_failed:
	return &response, _err
}

// Run for run cli asterisk server
func (c *AMIAction) Run(a *AMI) (*AMIResponse, error) {
	action := NewActionWith(config.AmiActionCommand)
	action.AddField(config.AmiActionCommand, c.Name)
	return c.Revoke(a, NewDictionary(), action, false)
}

// RunDictionary
func (c *AMIAction) RunDictionary(a *AMI, dictionaries map[string]string) (*AMIResponse, error) {
	action := NewActionWith(config.AmiActionCommand)
	action.AddField(config.AmiActionCommand, c.Name)
	d := NewDictionary()
	d.AddKeysTranslator(dictionaries)
	return c.Revoke(a, d, action, false)
}

// RunScript with script action
func (c *AMIAction) RunScript(a *AMI, script map[string]string) (*AMIResponse, error) {
	action := NewActionWith(c.Name)
	action.AddFields(script)
	return c.Revoke(a, NewDictionary(), action, false)
}

// WithRun with script action
func (c *AMIAction) WithRun(a *AMI, script, dictionaries map[string]string) (*AMIResponse, error) {
	action := NewActionWith(c.Name)
	action.AddFields(script)
	d := NewDictionary()
	d.AddKeysTranslator(dictionaries)
	return c.Revoke(a, d, action, false)
}
