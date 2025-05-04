// Package mapper содержит мапперы для работы с БД.
package mapper

import (
	"fmt"
	dom "post-app/internal/domain/author"
	"post-app/internal/domain/vo"
)

// AuthorRow - структура для маппинга авторов из БД.
type AuthorRow struct {
	ID   int32  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// MapRowToAuthor - функция маппинга авторов из БД.
func MapRowToAuthor(row AuthorRow) (*dom.Author, error) {
	authorID, err := vo.NewAuthorID(row.ID)
	if err != nil {
		return nil, fmt.Errorf("AuthorRow.MapRowToAuthor: %w", err)
	}

	authorName, err := dom.NewAuthorName(row.Name)
	if err != nil {
		return nil, fmt.Errorf("AuthorRow.MapRowToAuthor: %w", err)
	}

	return dom.RehydrateAuthor(authorID, authorName), nil
}
