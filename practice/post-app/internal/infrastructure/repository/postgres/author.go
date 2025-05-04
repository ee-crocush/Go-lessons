package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	dom "post-app/internal/domain/author"
	"post-app/internal/domain/vo"
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
func (r *AuthorRepository) Create(ctx context.Context, author *dom.Author) (vo.AuthorID, error) {
	var id int32

	const query = `
		INSERT INTO authors (name)
		VALUES ($1)
		RETURNING id;
	`
	err := r.pool.QueryRow(ctx, query, author.Name().Value()).Scan(&id)
	if err != nil {
		return vo.AuthorID{}, fmt.Errorf("AuthorRepository.Create: %w", err)
	}

	authorID, err := vo.NewAuthorID(id)
	if err != nil {
		return vo.AuthorID{}, fmt.Errorf("AuthorRepository.Create: %w", err)
	}

	return authorID, nil
}

// FindByID находит автора по его идентификатору.
func (r *AuthorRepository) FindByID(ctx context.Context, id vo.AuthorID) (*dom.Author, error) {
	var row mapper.AuthorRow

	const query = `SELECT id, name FROM authors WHERE id=$1 LIMIT 1`

	err := r.pool.QueryRow(ctx, query, id.Value()).Scan(&row.ID, &row.Name)
	if err != nil {
		return nil, fmt.Errorf("AuthorRepository.FindByID: %w", err)
	}

	return mapper.MapRowToAuthor(row)
}

// FindByIDs находит авторов по их идентификаторам.
func (r *AuthorRepository) FindByIDs(ctx context.Context, ids []vo.AuthorID) ([]*dom.Author, error) {
	authorsIDS := make([]int32, len(ids))
	for i, id := range ids {
		authorsIDS[i] = id.Value()
	}

	const query = `SELECT id, name FROM authors WHERE id = ANY($1)`

	rows, err := r.pool.Query(ctx, query, authorsIDS)
	if err != nil {
		return nil, fmt.Errorf("AuthorRepository.FindByIDs: %w", err)
	}
	defer rows.Close()

	var authors []*dom.Author

	for rows.Next() {
		var row mapper.AuthorRow

		if err = rows.Scan(&row.ID, &row.Name); err != nil {
			return nil, fmt.Errorf("AuthorRepository.FindByIDs: %w", err)
		}

		author, err := mapper.MapRowToAuthor(row)
		if err != nil {
			return nil, fmt.Errorf("AuthorRepository.FindByIDs: %w", err)
		}

		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("AuthorRepository.FindByIDs: %w", err)
	}

	return authors, nil
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
