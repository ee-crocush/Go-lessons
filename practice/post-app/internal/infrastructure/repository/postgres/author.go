package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	dom "post-app/internal/domain/author"
	"post-app/internal/infrastructure/repository/postgres/mapper"
)

var _ dom.Repository = (*AuthorRepository)(nil)

// AuthorRepository представляет собой репозиторий для работы с авторами в PostgreSQL.
type AuthorRepository struct {
	pool *pgxpool.Pool
}

// NewAuthorRepository создает новый экземпляр AuthorRepository.
func NewAuthorRepository(pool *pgxpool.Pool) *AuthorRepository {
	return &AuthorRepository{pool: pool}
}

// Create сохраняет нового автора в базе данных.
func (r *AuthorRepository) Create(ctx context.Context, author *dom.Author) error {
	const query = `
		INSERT INTO authors (name)
		VALUES ($1)
	`
	_, err := r.pool.Exec(ctx, query, author.Name().Value())
	if err != nil {
		return fmt.Errorf("AuthorRepository.Create: %w", err)
	}

	return nil
}

// FindByID находит автора по его идентификатору.
func (r *AuthorRepository) FindByID(ctx context.Context, id dom.AuthorID) (*dom.Author, error) {
	var row mapper.AuthorRow

	const query = `SELECT id, name FROM authors WHERE id=$1 LIMIT 1`

	err := r.pool.QueryRow(ctx, query, id.Value()).Scan(&row.ID, &row.Name)
	if err != nil {
		return nil, fmt.Errorf("AuthorRepository.FindByID: %w", err)
	}

	return mapper.MapRowToAuthor(row)
}

// Save сохраняет изменения в существующем авторе.
func (r *AuthorRepository) Save(ctx context.Context, author *dom.Author) error {
	const query = `
		UPDATE authors SET
			name=$2
		WHERE id=$1
	`
	cmd, err := r.pool.Exec(ctx, query, author.ID().Value(), author.Name().Value())
	if err != nil {
		return fmt.Errorf("AuthorRepository.Save: %w", err)
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("AuthorRepository.Save: %w", pgx.ErrNoRows)
	}

	return nil
}
