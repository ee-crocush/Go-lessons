package post

import (
	"context"
	"fmt"
	authordom "post-app/internal/domain/author"
	dom "post-app/internal/domain/post"
)

// GetByIDInputDTO входные данные для получения поста.
type GetByIDInputDTO struct {
	ID int32 `json:"id"`
}

// GetByIDOutputDTO выходные данные для получения поста.
type GetByIDOutputDTO struct {
	Author AuthorDTO `json:"author"`
	Post   PostDTO   `json:"post"`
}

// GetByIDContractUseCase определяет контракт для получения поста.
type GetByIDContractUseCase interface {
	Execute(ctx context.Context, in GetByIDInputDTO) (GetByIDOutputDTO, error)
}

// GetByIDUseCase бизнес логика получения поста.
type GetByIDUseCase struct {
	repo       dom.Finder
	authorRepo authordom.Finder
}

// NewGetByIDUseCase конструктор бизнес логики получения поста.
func NewGetByIDUseCase(repo dom.Finder, authorRepo authordom.Finder) *GetByIDUseCase {
	return &GetByIDUseCase{repo: repo, authorRepo: authorRepo}
}

// Execute выполняет бизнес логику.
func (uc *GetByIDUseCase) Execute(ctx context.Context, in GetByIDInputDTO) (GetByIDOutputDTO, error) {
	postID, err := dom.NewPostID(in.ID)
	if err != nil {
		return GetByIDOutputDTO{}, fmt.Errorf("Post.GetByIDUseCase.Execute: %w", err)
	}

	post, err := uc.repo.FindByID(ctx, postID)
	if err != nil {
		return GetByIDOutputDTO{}, fmt.Errorf("Post.GetByIDUseCase.Execute: %w", err)
	}

	author, err := uc.authorRepo.FindByID(ctx, post.AuthorID())
	if err != nil {
		return GetByIDOutputDTO{}, fmt.Errorf("Post.GetByIDUseCase.Execute: %w", err)
	}

	postDTO := MapPostToDTO(post)
	authorDTO := MapAuthorToDTO(author)

	return GetByIDOutputDTO{Author: authorDTO, Post: postDTO}, nil
}
