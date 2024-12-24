package custom_errors

import "errors"

var (
	ErrInvalidUsernameLength     = errors.New("invalid username length")
	ErrInvalidUsernameCharacters = errors.New("invalid username characters")
	ErrInvalidPasswordLength     = errors.New("invalid password length")
	ErrInvalidBalanceValue       = errors.New("invalid balance value")
	ErrUsernameIsNotUnique       = errors.New("username is not unique")
	ErrAccountIdNotFound         = errors.New("account id not found")
	ErrCanNotDeleteAccount       = errors.New("can not delete account with active rents")

	ErrTransportIdNotFound        = errors.New("transport id not found")
	ErrInvalidTransportType       = errors.New("invalid transport type")
	ErrInvalidTransportProperties = errors.New("invalid transport properties")
	ErrCanNotDeleteTransport      = errors.New("can not delete transport with active rents")

	ErrInvalidPaginationParams   = errors.New("invalid pagination params")
	ErrAccessDenied              = errors.New("access denied")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")

	ErrInvalidTokenType      = errors.New("token is of invalid type")
	ErrTokenNotFound         = errors.New("token not found")
	ErrInvalidIdType         = errors.New("id is of invalid type")
	ErrIdNotFound            = errors.New("temporary shit")
	ErrEmptyAuthHeader       = errors.New("empty auth header")
	ErrInvalidToken          = errors.New("invalid token")
	ErrInvalidParams         = errors.New("invalid params")
	ErrTransportNotAvailable = errors.New("transport is not available")
	ErrAlreadyStopped        = errors.New("rent has been already stopped")
	ErrCanNotRent            = errors.New("can not rent your own transport")
)
