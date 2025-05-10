package mapper

import (
	"fmt"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
)

// PostDocument - структура для маппинга постов из Mongo.
type PostDocument struct {
	ID        int32  `bson:"_id,omitempty"`
	PostID    int32  `bson:"author_id"`
	Title     string `bson:"title"`
	Content   string `bson:"content"`
	CreatedAt int64  `bson:"created_at"`
}

// MapDocToPost - функция для маппинга поста из Mongo.
func MapDocToPost(doc PostDocument) (*dom.Post, error) {
	id, err := dom.NewPostID(doc.ID)
	if err != nil {
		return nil, fmt.Errorf("PostDocument.MapDocToPost: %w", err)
	}

	authorID, err := vo.NewAuthorID(doc.PostID)
	if err != nil {
		return nil, fmt.Errorf("PostDocument.MapDocToPost: %w", err)
	}

	title, err := dom.NewPostTitle(doc.Title)
	if err != nil {
		return nil, fmt.Errorf("PostDocument.MapDocToPost: %w", err)
	}

	content, err := dom.NewPostContent(doc.Content)
	if err != nil {
		return nil, fmt.Errorf("PostDocument.MapDocToPost: %w", err)
	}

	createdAt := dom.FromUnixSeconds(doc.CreatedAt)

	return dom.RehydratePost(id, authorID, title, content, createdAt), nil
}

// FromPostToDoc маппинг доменной модели поста в MongoDB-документ.
func FromPostToDoc(a *dom.Post) *PostDocument {
	return &PostDocument{
		ID:        a.ID().Value(),
		PostID:    a.AuthorID().Value(),
		Title:     a.Title().Value(),
		Content:   a.Content().Value(),
		CreatedAt: a.CreatedAt().Time().Unix(),
	}
}
