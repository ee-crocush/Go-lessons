package postgres

import (
	"context"
	"fmt"
	"github.com/ee-crocush/task-tracker/internal/domain/label"
	"github.com/ee-crocush/task-tracker/internal/domain/task"
	"github.com/ee-crocush/task-tracker/internal/domain/user"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ task.TaskRepository = (*TaskPGRepository)(nil)

// TaskPGRepository представляет репозиторий задач в PostgreSQL.
type TaskPGRepository struct {
	pool *pgxpool.Pool
}

// NewTaskPGRepository создает новый экземпляр TaskPGRepository.
func NewTaskPGRepository(pool *pgxpool.Pool) *TaskPGRepository {
	return &TaskPGRepository{
		pool: pool,
	}
}

// Create сохраняет новую задачу в базе данных.
func (r *TaskPGRepository) Create(ctx context.Context, t *task.Task) (task.TaskID, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return task.TaskID{}, fmt.Errorf("TaskPGRepository.Create: begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	const query = `
    INSERT INTO tasks (author_id, assigned_id, title, content, opened)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
    `

	var id int
	err = tx.QueryRow(
		ctx, query,
		t.AuthorID().Value(),
		t.AssignedID().Value(),
		t.Title().Value(),
		t.Content().Value(),
		t.OpenedAt().Time().Unix(),
	).Scan(&id)
	if err != nil {
		return task.TaskID{}, fmt.Errorf("TaskPGRepository.Create: %w", err)
	}

	// Вставка меток по одной в цикле
	labels := t.Labels()
	if len(labels) > 0 {
		const insertLabelQuery = `INSERT INTO tasks_labels (task_id, label_id) VALUES ($1, $2)`

		for _, l := range labels {
			_, err = tx.Exec(ctx, insertLabelQuery, id, l.ID().Value())
			if err != nil {
				return task.TaskID{}, fmt.Errorf("TaskPGRepository.Create: insert label: %w", err)
			}
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return task.TaskID{}, fmt.Errorf("TaskPGRepository.Create: commit tx: %w", err)
	}

	return task.NewTaskID(id)
}

func (r *TaskPGRepository) MassCreate(ctx context.Context, tasks *[]task.Task) ([]task.TaskID, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("TaskPGRepository.MassCreate: begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	taskIDs := make([]task.TaskID, 0, len(*tasks))

	const query = `
    INSERT INTO tasks (author_id, assigned_id, title, content, opened)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
    `

	const insertLabelQuery = `INSERT INTO tasks_labels (task_id, label_id) VALUES ($1, $2)`

	for _, t := range *tasks {
		var id int
		err = tx.QueryRow(
			ctx, query,
			t.AuthorID().Value(),
			t.AssignedID().Value(),
			t.Title().Value(),
			t.Content().Value(),
			t.OpenedAt().Time().Unix(),
		).Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("TaskPGRepository.MassCreate: insert task: %w", err)
		}

		// Вставка меток для задачи
		labels := t.Labels()
		if len(labels) > 0 {
			for _, l := range labels {
				_, err = tx.Exec(ctx, insertLabelQuery, id, l.ID().Value())
				if err != nil {
					return nil, fmt.Errorf("TaskPGRepository.MassCreate: insert label: %w", err)
				}
			}
		}

		taskID, err := task.NewTaskID(id)
		if err != nil {
			return nil, fmt.Errorf("TaskPGRepository.MassCreate: %w", err)
		}
		taskIDs = append(taskIDs, taskID)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("TaskPGRepository.MassCreate: %w", err)
	}

	return taskIDs, nil
}

// FindAll получает все задачи.
func (r *TaskPGRepository) FindAll(ctx context.Context) ([]*task.Task, error) {
	rows, err := r.pool.Query(
		ctx, `
		SELECT id, title, content, author_id, assigned_id, opened
		FROM tasks
	`,
	)
	if err != nil {
		return nil, fmt.Errorf("TaskPGRepository.FindAll: %w", err)
	}
	defer rows.Close()

	var result []*task.Task

	for rows.Next() {
		var row dbTaskRow

		if err = rows.Scan(
			&row.ID, &row.Title, &row.Content, &row.AuthorID, &row.AssignedID, &row.OpenedAt,
			&row.ClosedAt,
		); err != nil {
			return nil, fmt.Errorf("TaskPGRepository.FindAll: %w", err)
		}

		t, err := MapRowToTask(row)
		if err != nil {
			return nil, fmt.Errorf("TaskPGRepository.FindAll: %w", err)
		}

		result = append(result, t)
	}

	return result, nil
}

// FindByID получает задачу по ее ID.
func (r *TaskPGRepository) FindByID(ctx context.Context, id task.TaskID) (*task.Task, error) {
	var row dbTaskRow

	const query = `
		SELECT id, author_id, assigned_id, title, content, opened, closed
		FROM tasks
		WHERE id = $1
	`
	err := r.pool.QueryRow(ctx, query, id.Value()).Scan(
		&row.ID, &row.AuthorID, &row.AssignedID, &row.Title,
		&row.Content, &row.OpenedAt, &row.ClosedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("TaskPGRepository.FindByID: %w", err)
	}

	return MapRowToTask(row)
}

// FindAllByAuthorID получает все задачи по автору.
func (r *TaskPGRepository) FindAllByAuthorID(ctx context.Context, authorID user.UserID) ([]*task.Task, error) {
	const query = `
		SELECT id, author_id, assigned_id, title, content, opened, closed
		FROM tasks
		WHERE author_id = $1
	`
	rows, err := r.pool.Query(ctx, query, authorID.Value())
	if err != nil {
		return nil, fmt.Errorf("TaskPGRepository.FindAllByAuthorID: %w", err)
	}
	defer rows.Close()

	var result []*task.Task

	for rows.Next() {
		var row dbTaskRow

		if err = rows.Scan(
			&row.ID, &row.AuthorID, &row.AssignedID, &row.Title, &row.Content, &row.OpenedAt,
			&row.ClosedAt,
		); err != nil {
			return nil, fmt.Errorf("TaskPGRepository.FindAllByAuthorID: %w", err)
		}

		t, err := MapRowToTask(row)
		if err != nil {
			return nil, fmt.Errorf("TaskPGRepository.FindAllByAuthorID: %w", err)
		}

		result = append(result, t)
	}

	return result, nil
}

// FindAllByLabelID получает все задачи по метке.
func (r *TaskPGRepository) FindAllByLabelID(ctx context.Context, labelID label.LabelID) ([]*task.Task, error) {
	const query = `
		SELECT t.id, t.author_id, t.assigned_id, t.title, t.content, t.opened, t.closed
		FROM tasks t
		JOIN tasks_labels tl ON tl.task_id = t.id
		WHERE tl.label_id = $1
	`

	rows, err := r.pool.Query(ctx, query, labelID.Value())
	if err != nil {
		return nil, fmt.Errorf("TaskPGRepository.FindAllByLabelID: %w", err)
	}
	defer rows.Close()

	var result []*task.Task
	for rows.Next() {
		var row dbTaskRow

		if err = rows.Scan(
			&row.ID, &row.Title, &row.Content, &row.AuthorID, &row.AssignedID, &row.OpenedAt,
			&row.ClosedAt,
		); err != nil {
			return nil, fmt.Errorf("TaskPGRepository.FindAllByAuthorID: %w", err)
		}

		t, err := MapRowToTask(row)
		if err != nil {
			return nil, fmt.Errorf("TaskPGRepository.FindAllByAuthorID: %w", err)
		}

		result = append(result, t)
	}

	return result, nil
}

// Update обновляет задачу в базе данных.
func (r *TaskPGRepository) Update(ctx context.Context, task *task.Task) error {
	const query = `
		UPDATE tasks SET
			title=$1, content=$2, author_id=$3, assigned_id=$4, closed=$5
		WHERE id = $6
	`

	cmdTag, err := r.pool.Exec(
		ctx, query,
		task.Title().Value(),
		task.Content().Value(),
		task.AuthorID().Value(),
		task.AssignedID().Value(),
		task.ClosedAt().Time().Unix(),
		task.ID().Value(),
	)
	if err != nil {
		return fmt.Errorf("TaskPGRepository.Update: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("TaskPGRepository.Update: %w", pgx.ErrNoRows)
	}

	return nil
}

func (r *TaskPGRepository) Delete(ctx context.Context, id task.TaskID) error {
	const query = `DELETE FROM tasks WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, id.Value())
	if err != nil {
		return fmt.Errorf("TaskPGRepository.Delete: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("TaskPGRepository.Delete: %w", pgx.ErrNoRows)
	}

	return nil
}
