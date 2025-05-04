package author

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/author"
	"post-app/internal/domain/vo"
)

// GetInputDTO входные данные для получения автора.
type GetInputDTO struct {
	ID int32 `json:"id"`
}

// GetOutputDTO выходные данные для получения автора.
type GetOutputDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// GetContractUseCase определяет контракт для получения автора.
type GetContractUseCase interface {
	Execute(ctx context.Context, in GetInputDTO) (GetOutputDTO, error)
}

// GetUseCase бизнес логика получения автора.
type GetUseCase struct {
	repo dom.Finder
}

// NewGetUseCase конструктор бизнес логики получения автора.
func NewGetUseCase(repo dom.Finder) *GetUseCase {
	return &GetUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *GetUseCase) Execute(ctx context.Context, in GetInputDTO) (GetOutputDTO, error) {
	authorID, err := vo.NewAuthorID(in.ID)
	if err != nil {
		return GetOutputDTO{}, fmt.Errorf("Author.GetUseCase.Execute: %w", err)
	}

	author, err := uc.repo.FindByID(ctx, authorID)
	if err != nil {
		return GetOutputDTO{}, fmt.Errorf("Author.GetUseCase.Execute: %w", err)
	}

	return GetOutputDTO{ID: author.ID().Value(), Name: author.Name().Value()}, nil
}
