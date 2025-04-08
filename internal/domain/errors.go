package domain

import "errors"

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErroDuplicatedAPIKey = errors.New("duplicated API key")
	ErrInvoiceNotFound   = errors.New("invoice not found")
	ErrUnauthorized      = errors.New("unauthorized")
)
