package xerror

import "errors"

var (
	ErrRowNotFound = errors.New("row not found")
)

var (
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
