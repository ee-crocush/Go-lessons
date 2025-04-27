package postgres

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/label"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ dom.LabelRepository = (*LabelPGRepository)(nil)

// LabelPGRepository представляет репозиторий меток в PostgreSQL.
type LabelPGRepository struct {
	pool *pgxpool.Pool
}

// NewLabelPGRepository создает новый экземпляр LabelPGRepository
func NewLabelPGRepository(pool *pgxpool.Pool) *LabelPGRepository {
	return &LabelPGRepository{pool: pool}
}

// Create сохраняет новую метку в базе данных.
func (r *LabelPGRepository) Create(ctx context.Context, label *dom.Label) (dom.LabelID, error) {
	const query = `
		INSERT INTO labels (name)
		VALUES ($1)
		RETURNING id
	`

	var id int
	err := r.pool.QueryRow(ctx, query, label.Name().Value()).Scan(&id)
	if err != nil {
		return dom.LabelID{}, fmt.Errorf("LabelPGRepository.Create: %w", err)
	}

	// Устанавливаем ID созданной метки.
	labelID, err := dom.NewLabelID(id)
	if err != nil {
		return dom.LabelID{}, fmt.Errorf("LabelPGRepository.Create: %w", err)
	}

	return labelID, nil
}
