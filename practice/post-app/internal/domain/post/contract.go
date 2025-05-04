package post

import (
	"context"
	"post-app/internal/domain/vo"
)

// Creator представляет контракт для создания поста.
type Creator interface {
	Create(ctx context.Context, post *Post) error
}

// Finder представляет контракт для получения поста.
type Finder interface {
	FindByID(ctx context.Context, id PostID) (*Post, error)
	FindByAuthorID(ctx context.Context, authorID vo.AuthorID) ([]*Post, error)
	FindAll(ctx context.Context) ([]*Post, error)
}

// Writer представляет контракт для сохранения/обновления поста.
type Writer interface {
	Save(ctx context.Context, post *Post) error
}

// Deleter представляет контракт для удаления поста.
type Deleter interface {
	DeleteByID(ctx context.Context, id PostID) error
}
