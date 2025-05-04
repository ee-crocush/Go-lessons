package vo

import "errors"

var (
	// ErrInvalidAuthorID представляет ошибку невалидного идентификатора автора.
	ErrInvalidAuthorID = errors.New("invalid Author ID")
)
