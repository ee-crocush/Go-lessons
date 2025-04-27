package task

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/task"
)

// FindByIDInputDTO входные данные для Получения задачи по ID.
type FindByIDInputDTO struct {
	ID int `json:"author_id"`
}

// FindByIDContractUseCase определяет контракт для Получения задачи.
type FindByIDContractUseCase interface {
	Execute(ctx context.Context, in FindByIDInputDTO) (TaskDTO, error)
}

// FindByIDUseCase бизнес логика Получения задачи по ID.
type FindByIDUseCase struct {
	repo dom.TaskFinder
}

// NewFindByIDUseCase конструктор бизнес логики Получения задачи по ID.
func NewFindByIDUseCase(repo dom.TaskFinder) *FindByIDUseCase {
	return &FindByIDUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *FindByIDUseCase) Execute(ctx context.Context, in FindByIDInputDTO) (
	TaskDTO, error,
) {
	taskID, err := dom.NewTaskID(in.ID)
	if err != nil {
		return TaskDTO{}, fmt.Errorf("FindByIDUseCase.Task.Execute: %w", err)
	}

	task, err := uc.repo.FindByID(ctx, taskID)
	if err != nil {
		return TaskDTO{}, fmt.Errorf("FindByIDUseCase.Task.Execute: %w", err)
	}

	return toTaskDTO(task), nil
}
