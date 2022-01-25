package common

import (
	"fmt"
	"net/http"
)

type ErrorCustom struct {
	StatusCode int    `json:"-"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
}

func NewErrorResponse(root error, msg, log string) *ErrorCustom {
	return &ErrorCustom{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log string) *ErrorCustom {
	return &ErrorCustom{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
	}
}

func (e *ErrorCustom) RootError() error {
	// check RootErr có phải là kiểu ErrorCustom không.
	// nếu ok thì lấy error gốc
	// type assertions
	fmt.Println("====>check", e.RootErr)
	if err, ok := e.RootErr.(*ErrorCustom); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *ErrorCustom) Error() string {
	return e.RootError().Error()
}

func ErrInvalidRequest(err error) *ErrorCustom {
	return NewErrorResponse(err, "invalid request", err.Error())
}

func ErrInternal(err error) *ErrorCustom {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "server error!", err.Error())
}
