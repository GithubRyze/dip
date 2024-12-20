package errors

import "errors"

var (
	NotMatchRouterError = errors.New("no match router")
	InternalServerError = errors.New("internal server error")
	NoSupportPathType   = errors.New("no support path type")
	NotOpenedError      = errors.New("not opened")
	NotInEffectiveError = errors.New("not in effective time")
	ExpirationError     = errors.New("expiration time")
	IpForbiddenError    = errors.New("ip forbidden")
	IpNotAllowedError   = errors.New("ip not allowed")
)
