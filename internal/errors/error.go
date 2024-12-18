package errors

import "errors"

var (
	NotMatchRouterError = errors.New("no match router")
	InternalServerError = errors.New("internal server error")
	NoSupportPathType   = errors.New("no support path type")
)
