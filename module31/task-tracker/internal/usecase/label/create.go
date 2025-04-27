// Package label содержит бизнес логику для меток
package label

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/label"
)

// CreateInputDTO входные данные для создания метки.
type CreateInputDTO struct {
	Name string `json:"name"`
}

// CreateOutputDTO выходные данные для создания метки.
// Так как мы не будем в рамказ учебного проекта использовать контракт на получения метки,
// то возвращать будем и Id и Name.
type CreateOutputDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateContractUseCase определяет контракт для создания метки.
type CreateContractUseCase interface {
	Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error)
}

// CreateUseCase бизнес логика создания метки.
type CreateUseCase struct {
	repo dom.LabelCreator
}

// NewCreateUseCase конструктор бизнес логики создания метки.
func NewCreateUseCase(repo dom.LabelCreator) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *CreateUseCase) Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error) {
	label, err := dom.NewLabel(in.Name)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.Label.Execute: %w", err)
	}

	id, err := uc.repo.Create(ctx, label)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.Label.Execute: %w", err)
	}

	return CreateOutputDTO{
		ID:   id.Value(),
		Name: label.Name().Value(),
	}, nil
}
