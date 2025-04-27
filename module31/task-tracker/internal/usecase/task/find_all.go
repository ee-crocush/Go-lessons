package task

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/task"
)

// FindAllOutputDTO выходные данные для Получения задач.
type FindAllOutputDTO struct {
	Tasks []TaskDTO `json:"tasks"`
}

// timeFormat формат времени
//const timeFormat = "2011-01-01 11:11:11"

// FindAllContractUseCase определяет контракт для Получения задач.
type FindAllContractUseCase interface {
	Execute(ctx context.Context) (FindAllOutputDTO, error)
}

// FindAllUseCase бизнес логика Получения задач.
type FindAllUseCase struct {
	repo dom.TaskFinder
}

// NewUseCase конструктор бизнес логики Получения задач.
func NewFindAllUseCase(repo dom.TaskFinder) *FindAllUseCase {
	return &FindAllUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *FindAllUseCase) Execute(ctx context.Context) (FindAllOutputDTO, error) {
	tasks, err := uc.repo.FindAll(ctx)
	if err != nil {
		return FindAllOutputDTO{}, fmt.Errorf("FindAllUseCase.Task.Execute: %w", err)
	}

	var result []TaskDTO
	for _, task := range tasks {
		result = append(result, toTaskDTO(task))
	}

	return FindAllOutputDTO{Tasks: result}, nil
}
