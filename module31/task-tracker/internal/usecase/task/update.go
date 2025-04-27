package task

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/task"
)

// UpdateInputDTO входные данные для обновления задачи.
type UpdateInputDTO struct {
	TaskDTO
}

// UpdateContractUseCase определяет контракт для обновления задачи.
type UpdateContractUseCase interface {
	Execute(ctx context.Context, in UpdateInputDTO) error
}

// UpdateUseCase бизнес логика обновления задачи.
type UpdateUseCase struct {
	repo dom.TaskUpdater
}

// NewUpdateUseCase конструктор бизнес логики обновления задачи.
func NewUpdateUseCase(repo dom.TaskUpdater) *UpdateUseCase {
	return &UpdateUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *UpdateUseCase) Execute(ctx context.Context, in UpdateInputDTO) error {
	taskID, err := dom.NewTaskID(in.ID)
	if err != nil {
		return fmt.Errorf("UpdateUseCase.Task.Execute: %w", err)
	}

	task, err := dom.NewTask(in.AuthorID, in.AssignedID, in.Title, in.Content, in.OpenedAt, in.ClosedAt)
	if err != nil {
		return fmt.Errorf("UpdateUseCase.Task.Execute: %w", err)
	}

	task.SetID(taskID)

	err = uc.repo.Update(ctx, task)
	if err != nil {
		return fmt.Errorf("UpdateUseCase.Task.Execute: %w", err)
	}

	return nil
}
