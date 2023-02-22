package ari

import "fmt"

type asteriskCXError struct {
	s string
	e string
}

func AsteriskErrorNew(ctx string) *asteriskCXError {
	return &asteriskCXError{s: ctx}
}

func (e *asteriskCXError) AsteriskErrorWith(msg string, args ...interface{}) *asteriskCXError {
	txt := fmt.Sprintf(msg, args...)
	e.e = fmt.Sprintf(": %s", txt)
	return e
}

func (e *asteriskCXError) Error() string {
	return fmt.Sprintf("asterisk error: %s%s", e.s, e.e)
}
