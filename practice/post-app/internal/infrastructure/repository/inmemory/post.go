package inmemory

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
	"sync"
)

var _ dom.Repository = (*PostRepository)(nil)

// PostRepository представляет БД постов в памяти.
type PostRepository struct {
	mu     sync.RWMutex
	lastID int32
	posts  map[dom.PostID]*dom.Post
}

// NewPostRepository возвращает новый in-memory репозиторий постов.
func NewPostRepository() *PostRepository {
	return &PostRepository{
		posts: make(map[dom.PostID]*dom.Post),
	}
}

// Create создает новый пост.
func (r *PostRepository) Create(ctx context.Context, p *dom.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	newID, err := dom.NewPostID(r.lastID)
	if err != nil {
		return fmt.Errorf("PostRepository.Create: %w", err)
	}

	p.SetID(newID)
	r.posts[newID] = p

	return nil
}

// FindByID находит пост по его ID.
func (r *PostRepository) FindByID(ctx context.Context, id dom.PostID) (*dom.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	post, ok := r.posts[id]
	if !ok {
		return nil, fmt.Errorf("PostRepository.FindByID: %v", dom.ErrPostNotFound)
	}

	return post, nil
}

// FindByAuthorID находит пост по автору.
func (r *PostRepository) FindByAuthorID(ctx context.Context, authorID vo.AuthorID) ([]*dom.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*dom.Post
	for _, p := range r.posts {
		if p.AuthorID() == authorID {
			result = append(result, p)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("PostRepository.FindByAuthorID: %v", dom.ErrPostNotFound)
	}

	return result, nil
}

// FindAll находит все посты.
func (r *PostRepository) FindAll(ctx context.Context) ([]*dom.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*dom.Post
	for _, p := range r.posts {
		result = append(result, p)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("PostRepository.FindAll: %v", dom.ErrPostNotFound)
	}

	return result, nil
}

// Save обновляет пост.
func (r *PostRepository) Save(ctx context.Context, post *dom.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	postID := post.ID()
	if _, exists := r.posts[postID]; !exists {
		return fmt.Errorf("PostRepository.Save: %w", dom.ErrPostNotFound)
	}

	r.posts[postID] = post

	return nil
}

// DeleteByID удаляет пост по его ID.
func (r *PostRepository) DeleteByID(ctx context.Context, id dom.PostID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.posts[id]; !exists {
		return fmt.Errorf("PostRepository.DeleteByID: %w", dom.ErrPostNotFound)
	}

	delete(r.posts, id)

	return nil
}
