package api

import (
	"fmt"
)

// NOTE: the returned value cannot be `nil`
func NewClientSDKRequestError(raw string) *Error {
	return &Error{
		Code:    "ClientSDKRequestError",
		CodeN:   1709701,
		Message: fmt.Sprintf("MaaS SDK request error: %s", raw),
	}
}

func (m *Error) Error() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s: %s, code=%d", m.Code, m.Message, m.CodeN)
}
