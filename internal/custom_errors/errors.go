package custom_errors

import "errors"

var (
	ErrUsernameExists            = errors.New("username already exists")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrInvalidTokenType          = errors.New("token is of invalid type")
	ErrTokenNotFound             = errors.New("token not found")
	ErrInvalidIdType             = errors.New("id is of invalid type")
	ErrIdNotFound                = errors.New("id not found")
	ErrEmptyAuthHeader           = errors.New("empty auth header")
	ErrInvalidToken              = errors.New("invalid token")
	ErrAccessDenied              = errors.New("access denied")
	ErrInvalidParams             = errors.New("invalid params")
)
