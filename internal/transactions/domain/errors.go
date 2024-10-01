package domain

import "errors"

var (
	ErrorFileNotFound  = errors.New("transactions: file not found")
	ErrorInvalidHeader = errors.New("transactions: invalid header")
	ErrorInvalidID     = errors.New("transactions: invalid id")
	ErrorInvalidDate   = errors.New("transactions: invalid date format")
	ErrorInvalidAmount = errors.New("transactions: invalid amount format")
)
