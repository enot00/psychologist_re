package appErrors

import "errors"

var (
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrAuthorizationFailed  = errors.New("access denied")
)
