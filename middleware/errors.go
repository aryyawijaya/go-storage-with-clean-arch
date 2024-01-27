package middleware

import "errors"

var (
	ErrAuthorizationHeaderNotProvided = errors.New("authorization header is not provided")
	ErrAuthorizationHeaderFormat      = errors.New("invalid authorization header format")
	ErrInvalidAPIKey                  = errors.New("invalid API key")
)
