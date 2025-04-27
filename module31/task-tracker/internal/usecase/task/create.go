// Package task содержит бизнес логику для задач.
package task

import (
	"context"
	"fmt"
	"github.com/ee-crocush/task-tracker/internal/domain/label"
	dom "github.com/ee-crocush/task-tracker/internal/domain/task"
)

// CreateInputDTO входные данные для создания задачи.
type CreateInputDTO struct {
	AuthorID   int        `json:"author_id"`
	AssignedID int        `json:"assigned_id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Labels     []LabelDTO `json:"labels"`
}

// CreateOutputDTO выходные данные для создания задачи.
type CreateOutputDTO struct {
	ID int `json:"id"`
}

// CreateContractUseCase определяет контракт для создания задачи.
type CreateContractUseCase interface {
	Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error)
}

// CreateUseCase бизнес логика создания задачи.
type CreateUseCase struct {
	repo dom.TaskCreator
}

// NewCreateUseCase конструктор бизнес логики создания задачи.
func NewCreateUseCase(repo dom.TaskCreator) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *CreateUseCase) Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error) {
	openedAt := dom.NewTimestamp()

	task, err := dom.NewTask(in.AuthorID, in.AssignedID, in.Title, in.Content, openedAt.Time().Unix(), nil)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.Task.Execute: %w", err)
	}
	for _, l := range in.Labels {
		newlabel, err := label.NewLabel(l.Name)
		if err != nil {
			return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.Task.Execute: %w", err)
		}

		labelID, err := label.NewLabelID(l.ID)
		if err != nil {
			return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.Task.Execute: %w", err)
		}
		newlabel.SetID(labelID)

		err = task.AddLabel(newlabel)
		if err != nil {
			return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.Task.Execute: %w", err)
		}
	}

	id, err := uc.repo.Create(ctx, task)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.Task.Execute: %w", err)
	}

	return CreateOutputDTO{
		ID: id.Value(),
	}, nil
}
