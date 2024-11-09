// TODO: Improve error standardization and extensibility

package models

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

const (
	NotFoundErrorCode = iota
	DuplicateErrorCode
	InternalErrorCode
	UnauthorizedErrorCode
	BadRequestErrorCode
)

var (
	ErrNotFound     = NewCustomError(NotFoundErrorCode, "Requested resource not found.")
	ErrDuplicate    = NewCustomError(DuplicateErrorCode, "Duplicate entry.")
	InternalError   = NewCustomError(InternalErrorCode, "Internal server error.")
	ErrUnauthorized = NewCustomError(UnauthorizedErrorCode, "Unauthorized operation.")
	ErrBadRequest = NewCustomError(UnauthorizedErrorCode, "Bad request.")
)

func NewCustomError(statuscode int, message string) *CustomError {
	return &CustomError{
		message,
		statuscode,
	}
}

func (self *CustomError) Error() string {
	return self.Message
}

func (self *CustomError) ToJSON() ([]byte, error) {
	return json.Marshal(self)
}

func (self *CustomError) ToHttpStatusCode() *CustomError {
	switch self.StatusCode {
	case NotFoundErrorCode:
		return NewCustomError(http.StatusNotFound, self.Message)
	case DuplicateErrorCode:
		return NewCustomError(http.StatusConflict, self.Message)
	case InternalErrorCode:
		return NewCustomError(http.StatusInternalServerError, self.Message)
	case UnauthorizedErrorCode:
		return NewCustomError(http.StatusUnauthorized, self.Message)
  case BadRequestErrorCode:
		return NewCustomError(http.StatusUnauthorized, self.Message)
	default:
		return self
	}
}
