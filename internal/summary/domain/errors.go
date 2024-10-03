package domain

import "errors"

var (
	ErrEmailInvalidResponseCode = errors.New("email: invalid response code from email service")
)
