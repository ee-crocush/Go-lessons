// Package author содержит доменную область для сущности Автор.
package author

import (
	"fmt"
	"post-app/internal/domain/post"
	"post-app/internal/domain/vo"
)

// Author представляет автора поста.
type Author struct {
	id    vo.AuthorID
	name  AuthorName
	posts []*post.Post
}

// NewAuthor создает нового автора.
func NewAuthor(name string) (*Author, error) {
	authorName, err := NewAuthorName(name)
	if err != nil {
		return nil, fmt.Errorf("Author.NewAuthor: %w", err)
	}

	return &Author{
		name:  authorName,
		posts: make([]*post.Post, 0),
	}, nil
}

// ID возвращает идентификатор автора.
func (a *Author) ID() vo.AuthorID { return a.id }

// Name возвращает название автора.
func (a *Author) Name() AuthorName { return a.name }

func (a *Author) Posts() []*post.Post { return a.posts }

// RehydrateAuthor — вспомогательный конструктор для «восстановления» сущности из БД.
func RehydrateAuthor(id vo.AuthorID, name AuthorName) *Author {
	return &Author{
		id:   id,
		name: name,
	}
}

// SetID устанавливает значение идентификатора.
func (a *Author) SetID(id vo.AuthorID) { a.id = id }

// AddPost добавляет новый пост к списку постов автора.
func (a *Author) AddPost(post *post.Post) {
	// Проверяем, существует ли уже такой пост
	for _, existingPost := range a.posts {
		if existingPost.ID() == post.ID() {
			return
		}
	}
	a.posts = append(a.posts, post)
}
