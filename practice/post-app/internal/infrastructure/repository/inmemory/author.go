// Package inmemory содержит реализацию репозитория авторов в памяти.
package inmemory

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/author"
	"sync"
)

var _ dom.Repository = (*AuthorRepository)(nil)

// AuthorRepository представляет БД авторов в памяти.
// Предполагам, что авторов будет не будет много. Поэтому не используем map для хранения данных
type AuthorRepository struct {
	mu     sync.RWMutex
	lastID int32
	items  []*dom.Author
}

// NewAuthorRepository создает новый репозиторий авторов.
func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{
		items: make([]*dom.Author, 0),
	}
}

// Create сохраняет нового автора.
func (r *AuthorRepository) Create(ctx context.Context, author *dom.Author) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++

	authorID, err := dom.NewAuthorID(r.lastID)
	if err != nil {
		return fmt.Errorf("AuthorRepository.Create: %v", err)
	}

	author.SetID(authorID)
	r.items = append(r.items, author)

	return nil
}

// FindByID находит автора по его ID.
func (r *AuthorRepository) FindByID(ctx context.Context, id dom.AuthorID) (*dom.Author, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.lastID++

	for _, a := range r.items {
		if a.ID() == id {
			return a, nil
		}
	}

	return nil, fmt.Errorf("AuthorRepository.FindByID: %v", dom.ErrAuthorNotFound)
}

// Save сохраняет изменения в существующем авторе.
func (r *AuthorRepository) Save(ctx context.Context, author *dom.Author) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, a := range r.items {
		if a.ID() == author.ID() {
			r.items[i] = author
			return nil
		}
	}

	return fmt.Errorf("AuthorRepository.FindByID: %v", dom.ErrAuthorNotFound)
}
