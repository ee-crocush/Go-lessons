package author

// AuthorID идентификатор автора поста.
type AuthorID struct {
	value int32
}

// NewAuthorID создает новый идентификатор автора поста.
func NewAuthorID(id int32) (AuthorID, error) {
	if id < 1 {
		return AuthorID{}, ErrInvalidAuthorID
	}
	return AuthorID{value: id}, nil
}

// Value возвращает значение идентификатора автора поста.
func (l AuthorID) Value() int32 { return l.value }

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
