package user

import "errors"

var (
	ErrInvalidNameLength = errors.New("invalid user name length")
)
