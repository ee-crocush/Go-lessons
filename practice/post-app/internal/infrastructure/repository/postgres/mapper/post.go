package mapper

import (
	"fmt"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
)

// PostRow - структура для маппинга постов из БД.
type PostRow struct {
	ID        int32  `db:"id" json:"id"`
	AuthorID  int32  `db:"author_id" json:"author_id"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

// MapRowToPost - функция для маппинга поста из БД.
func MapRowToPost(row PostRow) (*dom.Post, error) {
	id, err := dom.NewPostID(row.ID)
	if err != nil {
		return nil, fmt.Errorf("PostRow.MapRowToPost: %w", err)
	}

	authorID, err := vo.NewAuthorID(row.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("PostRow.MapRowToPost: %w", err)
	}

	title, err := dom.NewPostTitle(row.Title)
	if err != nil {
		return nil, fmt.Errorf("PostRow.MapRowToPost: %w", err)
	}

	content, err := dom.NewPostContent(row.Content)
	if err != nil {
		return nil, fmt.Errorf("PostRow.MapRowToPost: %w", err)
	}

	createdAt := dom.FromUnixSeconds(row.CreatedAt)

	return dom.RehydratePost(id, authorID, title, content, createdAt), nil
}
