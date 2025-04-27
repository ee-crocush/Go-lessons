// Package user содержит бизнес логику пользователей.
package user

import (
	"context"
	"fmt"
	dom "github.com/ee-crocush/task-tracker/internal/domain/user"
)

// CreateInputDTO входные данные для создания пользователя.
type CreateInputDTO struct {
	Name string `json:"name"`
}

// CreateOutputDTO выходные данные для создания пользователя.
type CreateOutputDTO struct {
	ID int `json:"id"`
}

// CreateContractUseCase определяет контракт для создания пользователя.
type CreateContractUseCase interface {
	Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error)
}

// CreateUseCase бизнес логика создания пользователя.
type CreateUseCase struct {
	repo dom.UserCreator
}

// NewCreateUseCase конструктор бизнес логики создания пользователя.
func NewCreateUseCase(repo dom.UserCreator) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *CreateUseCase) Execute(ctx context.Context, in CreateInputDTO) (CreateOutputDTO, error) {
	user, err := dom.NewUser(in.Name)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.User.Execute: %w", err)
	}

	id, err := uc.repo.Create(ctx, user)
	if err != nil {
		return CreateOutputDTO{}, fmt.Errorf("CreateUseCase.User.Execute: %w", err)
	}

	return CreateOutputDTO{
		ID: id.Value(),
	}, nil
}
