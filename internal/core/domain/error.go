package domain

import "errors"

var (
	ErrAuthRequired = errors.New("Username and password are required")
	ErrInternal = errors.New("Internal server error")
)
