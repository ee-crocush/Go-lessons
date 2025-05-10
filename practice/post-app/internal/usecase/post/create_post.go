// Package post содержит бизнес-логику работы с постами.
package post

import (
	"context"
	"fmt"
	authordom "post-app/internal/domain/author"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
)

// CreateInputDTO входные данные для создания поста.
type CreateInputDTO struct {
	AuthorID int32  `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// CreateContractUseCase определяет контракт для создания поста.
type CreateContractUseCase interface {
	Execute(ctx context.Context, in CreateInputDTO) error
}

// CreateUseCase бизнес логика создания поста.
type CreateUseCase struct {
	repo       dom.Creator
	authorRepo authordom.Finder
}

// NewCreateUseCase конструктор бизнес логики создания поста.
func NewCreateUseCase(repo dom.Creator, authorRepo authordom.Finder) *CreateUseCase {
	return &CreateUseCase{repo: repo, authorRepo: authorRepo}
}

// Execute выполняет бизнес логику.
func (uc *CreateUseCase) Execute(ctx context.Context, in CreateInputDTO) error {
	authorID, err := vo.NewAuthorID(in.AuthorID)
	if err != nil {
		return fmt.Errorf("Post.CreateUseCase.Execute: %w", err)
	}

	author, err := uc.authorRepo.FindByID(ctx, authorID)
	if err != nil {
		return fmt.Errorf("Post.CreateUseCase.Execute: %w", err)
	}

	post, err := dom.NewPost(author.ID(), in.Title, in.Content)
	if err != nil {
		return fmt.Errorf("Post.CreateUseCase.Execute: %w", err)
	}

	if err = uc.repo.Create(ctx, post); err != nil {
		return fmt.Errorf("Post.CreateUseCase.Execute: %w", err)
	}

	return nil
}
