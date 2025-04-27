package task

import (
	"context"
	"github.com/ee-crocush/task-tracker/internal/domain/label"
	"github.com/ee-crocush/task-tracker/internal/domain/user"
)

// TaskCreator представляет контракт для создания задачи.
type TaskCreator interface {
	Create(ctx context.Context, task *Task) (TaskID, error)
	MassCreate(ctx context.Context, tasks *[]Task) ([]TaskID, error)
}

// TaskFinder представляет контракт для поиска задач.
// Не будет усложнять, просто ищем по id.
type TaskFinder interface {
	FindAll(ctx context.Context) ([]*Task, error)
	FindByID(ctx context.Context, id TaskID) (*Task, error)
	FindAllByAuthorID(ctx context.Context, authorID user.UserID) ([]*Task, error)
	FindAllByLabelID(ctx context.Context, labelID label.LabelID) ([]*Task, error)
}

// TaskUpdater представляет контракт для обновления задачи.
type TaskUpdater interface {
	Update(ctx context.Context, task *Task) error
}

// TaskDeleter представляет контракт для удаления задачи.
type TaskDeleter interface {
	Delete(ctx context.Context, taskID TaskID) error
}
