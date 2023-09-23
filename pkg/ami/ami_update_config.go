package ami

import (
	"fmt"
	"strings"
)

func NewAMIUpdateConfigAction() *AMIUpdateConfigAction {
	a := &AMIUpdateConfigAction{}
	return a
}

func (a *AMIUpdateConfigAction) SetAction(value string) *AMIUpdateConfigAction {
	a.Action = value
	return a
}

func (a *AMIUpdateConfigAction) SetCategory(value string) *AMIUpdateConfigAction {
	a.Category = value
	return a
}

func (a *AMIUpdateConfigAction) SetVar(value string) *AMIUpdateConfigAction {
	a.Var = value
	return a
}

func (a *AMIUpdateConfigAction) SetValue(value string) *AMIUpdateConfigAction {
	a.Value = value
	return a
}

func (a *AMIUpdateConfigAction) SetVars(delimiter string, variables ...string) *AMIUpdateConfigAction {
	if strings.EqualFold(delimiter, "") {
		delimiter = ";"
	}
	vars := strings.Join(variables, delimiter)
	a.SetVar(vars)
	return a
}

func (a *AMIUpdateConfigAction) SetVarsMap(delimiter string, variables map[string]interface{}) *AMIUpdateConfigAction {
	if len(variables) <= 0 {
		return a
	}
	_vars := make([]string, len(variables))
	for k, v := range variables {
		str := fmt.Sprintf("%s=%v", k, JsonString(v))
		_vars = append(_vars, str)
	}
	a.SetVars(delimiter, _vars...)
	return a
}
