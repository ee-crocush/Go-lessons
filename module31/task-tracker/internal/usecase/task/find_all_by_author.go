package task

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/task"
	"github.com/ee-crocush/task-tracker/internal/domain/user"
)

// FindAllByAuthorIDDTO входные данные для Получения задач по автору.
type FindAllByAuthorIDDTO struct {
	AuthorID int `json:"author_id"`
}

// FindAllByAuthorIDOutputDTO выходные данные для Получения задач по автору.
type FindAllByAuthorIDOutputDTO struct {
	Tasks []TaskDTO `json:"tasks"`
}

// FindAllByAuthorIDContractUseCase определяет контракт для Получения задач.
type FindAllByAuthorIDContractUseCase interface {
	Execute(ctx context.Context, in FindAllByAuthorIDDTO) (FindAllByAuthorIDOutputDTO, error)
}

// FindAllByAuthorIDUseCase бизнес логика Получения задач по автору.
type FindAllByAuthorIDUseCase struct {
	repo dom.TaskFinder
}

// NewFindAllByAuthorIDUseCase конструктор бизнес логики Получения задач по автору.
func NewFindAllByAuthorIDUseCase(repo dom.TaskFinder) *FindAllByAuthorIDUseCase {
	return &FindAllByAuthorIDUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *FindAllByAuthorIDUseCase) Execute(ctx context.Context, in FindAllByAuthorIDDTO) (
	FindAllByAuthorIDOutputDTO, error,
) {
	authorID, err := user.NewUserID(in.AuthorID)
	if err != nil {
		return FindAllByAuthorIDOutputDTO{}, fmt.Errorf("FindAllByAuthorIDUseCase.Task.Execute: %w", err)
	}

	tasks, err := uc.repo.FindAllByAuthorID(ctx, authorID)
	if err != nil {
		return FindAllByAuthorIDOutputDTO{}, fmt.Errorf("FindAllByAuthorIDUseCase.Task.Execute: %w", err)
	}

	var result []TaskDTO
	for _, task := range tasks {
		result = append(result, toTaskDTO(task))
	}

	return FindAllByAuthorIDOutputDTO{Tasks: result}, nil
}
