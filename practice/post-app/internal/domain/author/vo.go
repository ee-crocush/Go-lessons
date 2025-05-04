package author

// AuthorName представляет собой имя автора поста.
type AuthorName struct {
	name string
}

// NewAuthorName создает новое имя автора поста.
func NewAuthorName(name string) (AuthorName, error) {
	if len(name) > 0 {
		return AuthorName{name}, nil
	}

	return AuthorName{}, ErrEmptyAuthorName
}

// Value возвращает значение наименования автора поста.
func (n AuthorName) Value() string { return n.name }
