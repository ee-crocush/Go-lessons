package task

import "errors"

var (
	ErrInvalidLength  = errors.New("invalid length")
	ErrDuplicateLabel = errors.New("duplicate label")
)
