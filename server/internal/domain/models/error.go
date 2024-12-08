// TODO: Improve error standardization and extensibility

package models

import (
	"encoding/json"
	"strings"
)

type CustomError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type MultiError struct {
	Errors []CustomError `json:"errors"`
	Code   string        `json:"code"`
}

const (
	NotFoundErrorCode     = "not_found"
	DuplicateErrorCode    = "duplicate"
	InternalErrorCode     = "internal_error"
	UnauthorizedErrorCode = "unauthorized"
	BadRequestErrorCode   = "bad_request"
)

var (
	ErrNotFound     = NewCustomError(NotFoundErrorCode, "Requested resource not found.")
	ErrDuplicate    = NewCustomError(DuplicateErrorCode, "Duplicate entry.")
	InternalError   = NewCustomError(InternalErrorCode, "Internal server error.")
	ErrUnauthorized = NewCustomError(UnauthorizedErrorCode, "Unauthorized operation.")
	ErrBadRequest   = NewCustomError(BadRequestErrorCode, "Bad request.")
)

func NewMultiError(code string, errors []CustomError) *MultiError {
	return &MultiError{
		errors,
		code,
	}
}

func (self *MultiError) Error() string {
	var msgs []string

	for _, err := range self.Errors {
		msgs = append(msgs, err.Error())
	}

	return strings.Join(msgs, ", ")
}

func NewCustomError(code string, message string) *CustomError {
	return &CustomError{
		message,
		code,
	}
}

func (self *CustomError) Error() string {
	return self.Message
}

func (self *CustomError) ToJSON() ([]byte, error) {
	return json.Marshal(self)
}

func (self *CustomError) ToHttpStatusCode() *CustomError {
	switch self.Code {
	case NotFoundErrorCode:
		return NewCustomError("404", self.Message)
	case DuplicateErrorCode:
		return NewCustomError("409", self.Message)
	case InternalErrorCode:
		return NewCustomError("500", self.Message)
	case UnauthorizedErrorCode:
		return NewCustomError("401", self.Message)
	case BadRequestErrorCode:
		return NewCustomError("400", self.Message)
	default:
		return self
	}
}
