package task

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/task"
)

// DeleteInputDTO входные данные для удаления задачи.
type DeleteInputDTO struct {
	ID int `json:"id"`
}

// DeleteContractUseCase определяет контракт для удаления задачи.
type DeleteContractUseCase interface {
	Execute(ctx context.Context, in DeleteInputDTO) error
}

// DeleteUseCase бизнес логика удаления задачи.
type DeleteUseCase struct {
	repo dom.TaskDeleter
}

// NewDeleteUseCase конструктор бизнес логики удаления задачи.
func NewDeleteUseCase(repo dom.TaskDeleter) *DeleteUseCase {
	return &DeleteUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *DeleteUseCase) Execute(ctx context.Context, in DeleteInputDTO) error {
	taskID, err := dom.NewTaskID(in.ID)
	if err != nil {
		return fmt.Errorf("DeleteUseCase.Task.Execute: %w", err)
	}

	err = uc.repo.Delete(ctx, taskID)
	if err != nil {
		return fmt.Errorf("DeleteUseCase.Task.Execute: %w", err)
	}

	return nil
}
