package author

import "errors"

var (
	// ErrEmptyAuthorName представляет ошибку пустого имени автора.
	ErrEmptyAuthorName = errors.New("empty Author name")
	// ErrAuthorNotFound представляет ошибку ненайденного автора.
	ErrAuthorNotFound = errors.New("author not found")
)
