package utils

import "errors"

var (
	ErrorUnauthorized = errors.New("Unauthorized")
	ErrorForbidden    = errors.New("Forbidden")
	ErrorBadRequest   = errors.New("Malformed request")
)
