package author

import "errors"

var (
	// ErrInvalidAuthorID представляет ошибку невалидного идентификатора автора.
	ErrInvalidAuthorID = errors.New("invalid Author ID")
	// ErrEmptyAuthorName представляет ошибку пустого имени автора.
	ErrEmptyAuthorName = errors.New("empty Author name")
	// ErrAuthorNotFound представляет ошибку ненайденного автора.
	ErrAuthorNotFound = errors.New("author not found")
)
