// Package mapper содержит мапперы для работы с MongoDB.
package mapper

import (
	"fmt"
	dom "post-app/internal/domain/author"
	"post-app/internal/domain/vo"
)

// AuthorDocument - структура для маппинга авторов из Mongo.
type AuthorDocument struct {
	ID   int32  `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

// MapDocToAuthor - функция маппинга авторов из БД.
func MapDocToAuthor(doc AuthorDocument) (*dom.Author, error) {
	authorID, err := vo.NewAuthorID(doc.ID)
	if err != nil {
		return nil, fmt.Errorf("AuthorDocument.MapDocToAuthor: %w", err)
	}

	authorName, err := dom.NewAuthorName(doc.Name)
	if err != nil {
		return nil, fmt.Errorf("AuthorDocument.MapDocToAuthor: %w", err)
	}

	return dom.RehydrateAuthor(authorID, authorName), nil
}

// FromAuthorToDoc маппинг доменной модели автора в MongoDB-документ.
func FromAuthorToDoc(a *dom.Author) *AuthorDocument {
	return &AuthorDocument{
		ID:   a.ID().Value(),
		Name: a.Name().Value(),
	}
}
