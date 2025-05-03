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
// Предполагаем, что авторов будет не будет много. Поэтому не используем map для хранения данных
type AuthorRepository struct {
	mu      sync.RWMutex
	lastID  int32
	authors []*dom.Author
}

// NewAuthorRepository создает новый репозиторий авторов.
func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{
		authors: make([]*dom.Author, 0),
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
	r.authors = append(r.authors, author)

	return nil
}

// FindByID находит автора по его ID.
func (r *AuthorRepository) FindByID(ctx context.Context, id dom.AuthorID) (*dom.Author, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.lastID++

	for _, a := range r.authors {
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

	for i, a := range r.authors {
		if a.ID() == author.ID() {
			r.authors[i] = author
			return nil
		}
	}

	return fmt.Errorf("AuthorRepository.FindByID: %v", dom.ErrAuthorNotFound)
}
