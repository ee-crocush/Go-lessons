// Package author содержит доменную область для сущности Автор.
package author

import "fmt"

// Author представляет автора поста.
type Author struct {
	id   AuthorID
	name AuthorName
}

// NewAuthor создает нового автора.
func NewAuthor(name string) (*Author, error) {
	authorName, err := NewAuthorName(name)
	if err != nil {
		return nil, fmt.Errorf("Author.NewAuthor: %w", err)
	}

	return &Author{
		name: authorName,
	}, nil
}

// ID возвращает идентификатор автора.
func (a *Author) ID() AuthorID { return a.id }

// Name возвращает название автора.
func (a *Author) Name() AuthorName { return a.name }

// RehydrateAuthor — вспомогательный конструктор для «восстановления» сущности из БД.
func RehydrateAuthor(id AuthorID, name AuthorName) *Author {
	return &Author{
		id:   id,
		name: name,
	}
}

// SetID устанавливает значение идентификатора.
func (a *Author) SetID(id AuthorID) { a.id = id }
