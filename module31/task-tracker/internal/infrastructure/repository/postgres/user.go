package postgres

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ dom.UserRepository = (*UserPGRepository)(nil)

// UserPGRepository представляет репозиторий пользователя в PostgreSQL.
type UserPGRepository struct {
	pool *pgxpool.Pool
}

// NewUserPGRepository создает новый экземпляр UserPGRepository
func NewUserPGRepository(pool *pgxpool.Pool) *UserPGRepository {
	return &UserPGRepository{pool: pool}
}

// Create сохраняет нового пользователя в базе данных.
func (r *UserPGRepository) Create(ctx context.Context, user *dom.User) (dom.UserID, error) {
	const query = `
		INSERT INTO users (name)
		VALUES ($1)
		RETURNING id
	`

	var id int
	err := r.pool.QueryRow(ctx, query, user.Name().Value()).Scan(&id)
	if err != nil {
		return dom.UserID{}, fmt.Errorf("UserPGRepository.Create: %w", err)
	}

	// Устанавливаем ID созданного пользователя.
	userID, err := dom.NewUserID(id)
	if err != nil {
		return dom.UserID{}, fmt.Errorf("UserPGRepository.Create: %w", err)
	}

	return userID, nil
}
