// Package vo содержит value_objects.
package vo

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

// Equal сравнивает два идентификатора авторов постов.
func (l AuthorID) Equal(other AuthorID) bool { return l.Value() == other.Value() }
