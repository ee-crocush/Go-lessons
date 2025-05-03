package post

import "errors"

var (
	// ErrInvalidPostID представляет ошибку невалидного идентификатора поста.
	ErrInvalidPostID = errors.New("invalid Post ID")
	// ErrEmptyPostTitle представляет ошибку незаполненного заголовка поста.
	ErrEmptyPostTitle = errors.New("empty Post title")
	// ErrEmptyPostContent представляет ошибку незаполненного содержимого поста.
	ErrEmptyPostContent = errors.New("empty Post content")
	// ErrPostNotFound представляет ошибку ненайденного поста.
	ErrPostNotFound = errors.New("post not found")
)
