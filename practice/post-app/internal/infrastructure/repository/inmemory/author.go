// Package inmemory содержит реализацию репозитория авторов в памяти.
package inmemory

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/author"
	"post-app/internal/domain/vo"
	"sync"
)

var _ dom.Repository = (*AuthorRepository)(nil)

// AuthorRepository представляет БД авторов в памяти.
type AuthorRepository struct {
	mu      sync.RWMutex
	lastID  int32
	authors map[vo.AuthorID]*dom.Author
}

// NewAuthorRepository создает новый репозиторий авторов.
func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{
		authors: make(map[vo.AuthorID]*dom.Author),
	}
}

// Create сохраняет нового автора.
func (r *AuthorRepository) Create(ctx context.Context, author *dom.Author) (vo.AuthorID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++

	authorID, err := vo.NewAuthorID(r.lastID)
	if err != nil {
		return vo.AuthorID{}, fmt.Errorf("AuthorRepository.Create: %v", err)
	}

	author.SetID(authorID)
	r.authors[authorID] = author

	return authorID, nil
}

// FindByID находит автора по его ID.
func (r *AuthorRepository) FindByID(ctx context.Context, id vo.AuthorID) (*dom.Author, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	author, ok := r.authors[id]
	if !ok {
		return nil, fmt.Errorf("PostRepository.FindByID: %v", dom.ErrAuthorNotFound)
	}

	return author, nil
}

// FindByIDs находит авторов по их ID.
func (r *AuthorRepository) FindByIDs(ctx context.Context, ids []vo.AuthorID) ([]*dom.Author, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Используем map для исключения дубликатов
	uniqueAuthors := make(map[vo.AuthorID]*dom.Author)

	for _, id := range ids {
		if author, exists := r.authors[id]; exists {
			uniqueAuthors[id] = author
		}
	}

	var result []*dom.Author
	for _, author := range uniqueAuthors {
		result = append(result, author)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("PostRepository.FindByAuthorID: %v", dom.ErrAuthorNotFound)
	}
	return result, nil
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
