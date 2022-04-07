package errors

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

type MyError struct {
	Code    int
	Message string
}

func NewError(message string, code int) *MyError {
	return &MyError{
		Code:    code,
		Message: message,
	}
}

func (err *MyError) Error() string {
	return err.Message
}

func (err *MyError) GetGrpcCode() codes.Code {
	switch err.Code {
	case http.StatusConflict:
		return codes.AlreadyExists
	}

	return codes.Internal
}
