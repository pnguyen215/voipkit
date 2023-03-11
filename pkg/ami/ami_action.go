package ami

import (
	"fmt"
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func NewAction() *AMIAction {
	return &AMIAction{}
}

func NewRevokeAction(cmd string, timeout int) *AMIAction {
	cli := NewAction()
	cli.ActionCmd = cmd
	cli.Timeout = timeout
	return cli
}

// RevokeAction run cli on asterisk server
func (c *AMIAction) RevokeAction(a *AMI, d *AMIDictionary, e *AMIMessage, deadlock bool) (*AMIResponse, error) {
	var response AMIResponse
	var _err error

	if strings.EqualFold(c.ActionCmd, "") {
		response.ErrorMessage = fmt.Sprintf(config.AmiErrorFieldRequired, "action_cmd")
		response.IsSuccess = false
		_err = fmt.Errorf(response.ErrorMessage)
		return &response, _err
	}

	// trace log
	log.Printf(" [>] Ami run cli ::: '%v' \n timeout = %v", e.String(), c.Timeout)
	// call asterisk server cli
	a.Action(e)

	// listen all action feedback
	all := a.AllEvents()

	if deadlock {
		defer a.Close()
	}

	for {
		select {
		case message := <-all:
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

// RunAction for run cli asterisk server
func (c *AMIAction) RunAction(a *AMI) (*AMIResponse, error) {
	action := NewActionWith(config.AmiActionCommand)
	action.AddField(config.AmiActionCommand, c.ActionCmd)
	return c.RevokeAction(a, NewDictionary(), action, false)
}

// RunActionDict
func (c *AMIAction) RunActionDict(a *AMI, dictionaries map[string]string) (*AMIResponse, error) {
	action := NewActionWith(config.AmiActionCommand)
	action.AddField(config.AmiActionCommand, c.ActionCmd)
	d := NewDictionary()
	d.AddKeysTranslator(dictionaries)
	return c.RevokeAction(a, d, action, false)
}

// RunActionScript with script action
func (c *AMIAction) RunActionScript(a *AMI, script map[string]string) (*AMIResponse, error) {
	action := NewActionWith(c.ActionCmd)
	action.AddFields(script)
	return c.RevokeAction(a, NewDictionary(), action, false)
}

// RunActionScriptDict with script action
func (c *AMIAction) RunActionScriptDict(a *AMI, script, dictionaries map[string]string) (*AMIResponse, error) {
	action := NewActionWith(c.ActionCmd)
	action.AddFields(script)
	d := NewDictionary()
	d.AddKeysTranslator(dictionaries)
	return c.RevokeAction(a, d, action, false)
}
