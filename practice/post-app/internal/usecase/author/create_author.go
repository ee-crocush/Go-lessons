// Package author содержит бизнес-логику работы с авторами.
package author

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/author"
)

// CreateInputDTO входные данные для создания автора.
type CreateInputDTO struct {
	Name string `json:"name"`
}

// CreateOutputDTO выходные данные для создания автора.
type CreateOutputDTO struct {
	ID int32 `json:"id"`
}

// CreateContractUseCase определяет контракт для создания автора.
type CreateContractUseCase interface {
	Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error)
}

// CreateUseCase бизнес логика создания автора.
type CreateUseCase struct {
	repo dom.Creator
}

// NewCreateUseCase конструктор бизнес логики создания автора.
func NewCreateUseCase(repo dom.Creator) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *CreateUseCase) Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error) {
	author, err := dom.NewAuthor(in.Name)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("Author.CreateUseCase.Execute: %w", err)
	}

	authorID, err := uc.repo.Create(ctx, author)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("Author.CreateUseCase.Execute: %w", err)
	}

	return CreateOutputDTO{ID: authorID.Value()}, nil
}
