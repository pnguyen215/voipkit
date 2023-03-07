package ami

import (
	"fmt"
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func NewCli() *AMICommand {
	return &AMICommand{}
}

func NewCliWith(cmd string, timeout int) *AMICommand {
	cli := NewCli()
	cli.ActionCmd = cmd
	cli.Timeout = timeout
	return cli
}

// RevokeCommand run cli on asterisk server
func (c *AMICommand) RevokeCommand(a *AMI, d *AMIDictionary, e *AMIMessage, deadlock bool) (*AMIResponse, error) {
	var response AMIResponse
	var _err error

	if strings.EqualFold(c.ActionCmd, "") {
		response.ErrorMessage = fmt.Sprintf(config.AmiCliErrorFieldRequired, "cmd")
		response.IsSuccess = false
		_err = fmt.Errorf(response.ErrorMessage)
		return &response, _err
	}

	// trace log
	log.Printf("Ami run cli ::: '%v', timeout = %v", e.String(), c.Timeout)
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

// Command with script cli
func (c *AMICommand) Command(a *AMI, scripts map[string]string) (*AMIResponse, error) {
	action := NewActionWith(c.ActionCmd)
	action.AddFields(scripts)
	return c.RevokeCommand(a, NewDictionary(), action, false)
}

// RunCmd for run cli asterisk server
func (c *AMICommand) RunCmd(a *AMI) (*AMIResponse, error) {
	action := NewActionWith(config.AmiCliCommand)
	action.AddField(config.AmiCliCommand, c.ActionCmd)
	return c.RevokeCommand(a, NewDictionary(), action, false)
}

// RunCmdDictionaries
func (c *AMICommand) RunCmdDictionaries(a *AMI, dictionaries map[string]string) (*AMIResponse, error) {
	action := NewActionWith(config.AmiCliCommand)
	action.AddField(config.AmiCliCommand, c.ActionCmd)
	d := NewDictionary()
	d.AddKeysTranslator(dictionaries)
	return c.RevokeCommand(a, d, action, false)
}
