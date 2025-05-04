package post

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/post"
)

// DeleteInputDTO входные данные для удаления поста.
type DeleteInputDTO struct {
	ID int32 `json:"id"`
}

// DeleteContractUseCase определяет контракт для удаления поста.
type DeleteContractUseCase interface {
	Execute(ctx context.Context, in DeleteInputDTO) error
}

// DeleteUseCase бизнес логика удаления поста.
type DeleteUseCase struct {
	repo dom.Deleter
}

// NewDeleteUseCase конструктор бизнес логики удаления поста.
func NewDeleteUseCase(repo dom.Deleter) *DeleteUseCase {
	return &DeleteUseCase{repo: repo}
}

// Execute выполняет бизнес логику.
func (uc *DeleteUseCase) Execute(ctx context.Context, in DeleteInputDTO) error {
	postID, err := dom.NewPostID(in.ID)
	if err != nil {
		return fmt.Errorf("Post.DeleteUseCase.Execute: %w", err)
	}

	if err = uc.repo.DeleteByID(ctx, postID); err != nil {
		return fmt.Errorf("Post.DeleteUseCase.Execute: %w", err)
	}

	return nil
}
