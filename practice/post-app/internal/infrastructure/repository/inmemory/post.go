package inmemory

import (
	"context"
	dom "post-app/internal/domain/post"
	"sync"
)

var _ dom.Repository = (*PostRepository)(nil)

// PostRepository представляет БД постов в памяти.
type PostRepository struct {
	mu     sync.RWMutex
	lastID int32
	items  map[dom.PostID]*dom.Post
}

// NewPostRepository возвращает новый in-memory репозиторий постов.
func NewPostRepository() *PostRepository {
	return &PostRepository{
		items: make(map[dom.PostID]*dom.Post),
	}
}

// Create создает новый пост.
func (r *PostRepository) Create(ctx context.Context, p *post.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	newID := post.NewPostID(r.lastID)
	p.SetID(newID)
	r.posts[newID] = p
	return nil
}
