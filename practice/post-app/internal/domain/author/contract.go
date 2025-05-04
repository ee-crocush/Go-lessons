package author

import (
	"context"
	"post-app/internal/domain/vo"
)

// Creator представляет контракт для создания автора.
type Creator interface {
	Create(ctx context.Context, author *Author) (vo.AuthorID, error)
}

// Finder представляет контракт для получения автора.
type Finder interface {
	FindByID(ctx context.Context, id vo.AuthorID) (*Author, error)
	FindByIDs(ctx context.Context, ids []vo.AuthorID) ([]*Author, error)
}

// Writer представляет контракт для сохранения/обновления автора.
type Writer interface {
	Save(ctx context.Context, author *Author) error
}
