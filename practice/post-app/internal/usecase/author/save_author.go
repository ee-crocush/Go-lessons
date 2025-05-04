package author

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/author"
	"post-app/internal/domain/vo"
)

// SaveInputDTO входные данные для Сохранения автора.
type SaveInputDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// SaveContractUseCase определяет контракт для Сохранения автора.
type SaveContractUseCase interface {
	Execute(ctx context.Context, in SaveInputDTO) error
}

// SaveUseCase бизнес логика Сохранения автора.
type SaveUseCase struct {
	repo dom.Writer
}

// NewSaveUseCase конструктор бизнес логики Сохранения автора.
func NewSaveUseCase(repo dom.Writer) *SaveUseCase {
	return &SaveUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *SaveUseCase) Execute(ctx context.Context, in SaveInputDTO) error {
	authorID, err := vo.NewAuthorID(in.ID)
	if err != nil {
		return fmt.Errorf("Author.SaveUseCase.Execute: %w", err)
	}

	authorName, err := dom.NewAuthorName(in.Name)
	if err != nil {
		return fmt.Errorf("Author.SaveUseCase.Execute: %w", err)
	}

	author := dom.RehydrateAuthor(authorID, authorName)

	err = uc.repo.Save(ctx, author)
	if err != nil {
		return fmt.Errorf("Author.SaveUseCase.Execute: %w", err)
	}

	return nil
}
