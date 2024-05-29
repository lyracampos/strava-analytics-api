package domain

import "errors"

var (
	ErrUserNotAuthorized                  = errors.New("unauthorized user")
	ErrRequestAuthorizationHeaderExpected = errors.New("the request expected an authorization header")
	ErrRequestAuthorizationHeaderInvalid  = errors.New("authorization reader invalid, expecetd Bearer")
)
