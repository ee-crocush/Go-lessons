// Package post содержит доменную область для сущности Пост.
package post

import (
	"fmt"
	"post-app/internal/domain/vo"
)

// Post представляет пост.
type Post struct {
	id        PostID
	authorID  vo.AuthorID
	title     PostTitle
	content   PostContent
	createdAt Timestamp
}

// NewPost создает новый пост.
func NewPost(authorID vo.AuthorID, title, content string) (*Post, error) {
	postTitle, err := NewPostTitle(title)
	if err != nil {
		return nil, fmt.Errorf("NewPost.NewPostTitle: %w", err)
	}

	postContent, err := NewPostContent(content)
	if err != nil {
		return nil, fmt.Errorf("NewPost.NewPostTitle: %w", err)
	}

	return &Post{
		authorID:  authorID,
		title:     postTitle,
		content:   postContent,
		createdAt: NewTimestamp(),
	}, nil
}

// ID возвращает идентификатор поста.
func (p *Post) ID() PostID { return p.id }

// AuthorID возвращает идентификатор автора поста.
func (p *Post) AuthorID() vo.AuthorID { return p.authorID }

// Title возвращает заголовок поста.
func (p *Post) Title() PostTitle { return p.title }

// Content возвращает содержимое поста.
func (p *Post) Content() PostContent { return p.content }

// CreatedAt возвращает дату создания поста.
func (p *Post) CreatedAt() Timestamp { return p.createdAt }

// RehydratePost — вспомогательный конструктор для «восстановления» сущности из БД.
func RehydratePost(
	id PostID, authorID vo.AuthorID, title PostTitle, content PostContent, createdAt Timestamp,
) *Post {
	return &Post{
		id:        id,
		authorID:  authorID,
		title:     title,
		content:   content,
		createdAt: createdAt,
	}
}

// SetID устанавливает идентификатор поста.
func (p *Post) SetID(id PostID) { p.id = id }
