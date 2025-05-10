package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
	"post-app/internal/infrastructure/repository/postgres/mapper"
)

var _ dom.Repository = (*PostRepository)(nil)

// PostRepository представляет собой репозиторий для работы с постами в PostgreSQL.
type PostRepository struct {
	pool *pgxpool.Pool
}

// NewPostRepository создает новый экземпляр репозитория PostRepository.
func NewPostRepository(pool *pgxpool.Pool) *PostRepository {
	return &PostRepository{pool: pool}
}

// Create сохраняет новый пост в базе данных.
func (r *PostRepository) Create(ctx context.Context, post *dom.Post) error {
	const query = `
		INSERT INTO posts (author_id, title, content, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.pool.Exec(
		ctx, query, post.AuthorID().Value(), post.Title().Value(), post.Content().Value(),
		post.CreatedAt().Time().UTC().Unix(),
	)
	if err != nil {
		return fmt.Errorf("PostRepository.Create: %w", err)
	}

	return nil
}

// FindByID находит пост в базе данных по его идентификатору.
func (r *PostRepository) FindByID(ctx context.Context, id dom.PostID) (*dom.Post, error) {
	var row mapper.PostRow

	const query = `SELECT id, author_id, title, content, created_at FROM posts WHERE id=$1 LIMIT 1`

	err := r.pool.QueryRow(ctx, query, id.Value()).Scan(
		&row.ID, &row.AuthorID, &row.Title, &row.Content, &row.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.FindByID: %w", err)
	}

	return mapper.MapRowToPost(row)
}

// FindByAuthorID находит все посты автора в базе данных по его идентификатору.
func (r *PostRepository) FindByAuthorID(ctx context.Context, authorID vo.AuthorID) ([]*dom.Post, error) {
	const query = `SELECT id, author_id, title, content, created_at FROM posts WHERE author_id=$1`

	return r.fetchPosts(ctx, query, authorID.Value())
}

// FindAll находит все посты в базе данных.
func (r *PostRepository) FindAll(ctx context.Context) ([]*dom.Post, error) {
	const query = `SELECT id, author_id, title, content, created_at FROM posts`

	return r.fetchPosts(ctx, query)
}

func (r *PostRepository) Save(ctx context.Context, post *dom.Post) error {
	const query = `
		UPDATE posts SET
			author_id=$2, title=$3, content=$4
		WHERE id=$1
	`
	cmd, err := r.pool.Exec(
		ctx, query, post.ID().Value(), post.AuthorID().Value(), post.Title().Value(), post.Content().Value(),
	)
	if err != nil {
		return fmt.Errorf("PostRepository.Save: %w", err)
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("PostRepository.Save: %w", pgx.ErrNoRows)
	}

	return nil
}

func (r *PostRepository) DeleteByID(ctx context.Context, id dom.PostID) error {
	const query = `DELETE FROM posts WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, id.Value())
	if err != nil {
		return fmt.Errorf("PostRepository.Delete: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("PostRepository.Delete: %w", pgx.ErrNoRows)
	}

	return nil
}

// fetchPosts выполняет запрос и маппит строки в []*dom.Post.
func (r *PostRepository) fetchPosts(ctx context.Context, query string, args ...interface{}) ([]*dom.Post, error) {
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.fetchPosts: %w", err)
	}
	defer rows.Close()

	var posts []*dom.Post

	for rows.Next() {
		var row mapper.PostRow

		if err = rows.Scan(&row.ID, &row.AuthorID, &row.Title, &row.Content, &row.CreatedAt); err != nil {
			return nil, fmt.Errorf("PostRepository.fetchPosts: %w", err)
		}

		post, err := mapper.MapRowToPost(row)
		if err != nil {
			return nil, fmt.Errorf("PostRepository.fetchPosts: %w", err)
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("PostRepository.fetchPosts: %w", err)
	}

	return posts, nil
}
