package author

import "context"

// Creator представляет контракт для создания автора.
type Creator interface {
	Create(ctx context.Context, author *Author) error
}

// Finder представляет контракт для получения автора.
type Finder interface {
	FindByID(ctx context.Context, id AuthorID) (*Author, error)
}

// Writer представляет контракт для сохранения/обновления автора.
type Writer interface {
	Save(ctx context.Context, author *Author) error
}
