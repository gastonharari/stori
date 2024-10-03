package dtos

import "errors"

var (
	ErrMissingFile  = errors.New("missing required flag: --file")
	ErrMissingEmail = errors.New("missing required flag: --email")
)
