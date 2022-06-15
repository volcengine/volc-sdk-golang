package tls

import (
	"encoding/json"
)

type Error struct {
	HTTPCode  int32  `json:"httpCode"`
	Code      string `json:"errorCode"`
	Message   string `json:"errorMessage"`
	RequestID string `json:"requestID"`
}

func NewClientError(err error) *Error {
	if err == nil {
		return nil
	}
	if clientError, ok := err.(*Error); ok {
		return clientError
	}
	clientError := new(Error)
	clientError.HTTPCode = -1
	clientError.Code = "ClientError"
	clientError.Message = err.Error()
	return clientError
}

func (e Error) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e Error) Error() string {
	return e.String()
}

type BadResponseError struct {
	RespBody   string
	RespHeader map[string][]string
	HTTPCode   int
}

func (e BadResponseError) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e BadResponseError) Error() string {
	return e.String()
}

func NewBadResponseError(body string, header map[string][]string, httpCode int) *BadResponseError {
	return &BadResponseError{
		RespBody:   body,
		RespHeader: header,
		HTTPCode:   httpCode,
	}
}
