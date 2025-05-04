package post

import (
	"context"
	"fmt"
	authordom "post-app/internal/domain/author"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
)

// GetByAuthorIDInputDTO входные данные для получения поста по ID автора.
type GetByAuthorIDInputDTO struct {
	AuthorID int32 `json:"author_id"`
}

// GetByAuthorIDOutputDTO выходные данные для получения поста по ID автора.
type GetByAuthorIDOutputDTO struct {
	Author AuthorDTO `json:"author"`
	Posts  []PostDTO `json:"posts"`
}

// GetByAuthorIDContractUseCase определяет контракт для получения поста.
type GetByAuthorIDContractUseCase interface {
	Execute(ctx context.Context, in GetByAuthorIDInputDTO) (GetByAuthorIDOutputDTO, error)
}

// GetByAuthorIDUseCase бизнес логика получения поста.
type GetByAuthorIDUseCase struct {
	repo       dom.Finder
	authorRepo authordom.Finder
}

// NewGetByAuthorIDUseCase конструктор бизнес логики получения постов по ID автора.
func NewGetByAuthorIDUseCase(repo dom.Finder, authorRepo authordom.Finder) *GetByAuthorIDUseCase {
	return &GetByAuthorIDUseCase{repo: repo, authorRepo: authorRepo}
}

// Execute выполняет бизнес логику.
func (uc *GetByAuthorIDUseCase) Execute(ctx context.Context, in GetByAuthorIDInputDTO) (GetByAuthorIDOutputDTO, error) {
	authorID, err := vo.NewAuthorID(in.AuthorID)
	if err != nil {
		return GetByAuthorIDOutputDTO{}, fmt.Errorf("Post.GetByAuthorIDUseCase.Execute: %w", err)
	}

	author, err := uc.authorRepo.FindByID(ctx, authorID)
	if err != nil {
		return GetByAuthorIDOutputDTO{}, fmt.Errorf("Post.GetByAuthorIDUseCase.Execute: %w", err)
	}

	authorDTO := MapAuthorToDTO(author)

	posts, err := uc.repo.FindByAuthorID(ctx, authorID)
	if err != nil {
		return GetByAuthorIDOutputDTO{}, fmt.Errorf("Post.GetByAuthorIDUseCase.Execute: %w", err)
	}

	// Мапим полученные посты в DTO.
	var postsDTO []PostDTO
	for _, p := range posts {
		postDTO := MapPostToDTO(p)
		postsDTO = append(postsDTO, postDTO)
	}

	return GetByAuthorIDOutputDTO{Author: authorDTO, Posts: postsDTO}, nil
}
