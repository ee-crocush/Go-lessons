package post

import (
	"context"
	"fmt"
	authordom "post-app/internal/domain/author"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
)

// GetAllOutputDTO выходные данные для получения всех постов.
type GetAllOutputDTO struct {
	Posts []AuthorWithPostsDTO `json:"posts"`
}

// GetAllContractUseCase определяет контракт для получения постов.
type GetAllContractUseCase interface {
	Execute(ctx context.Context) (GetAllOutputDTO, error)
}

// GetAllUseCase бизнес логика получения поста.
type GetAllUseCase struct {
	repo       dom.Finder
	authorRepo authordom.Finder
}

// NewGetAllUseCase конструктор бизнес логики получения поста.
func NewGetAllUseCase(repo dom.Finder, authorRepo authordom.Finder) *GetAllUseCase {
	return &GetAllUseCase{repo: repo, authorRepo: authorRepo}
}

// Execute выполняет бизнес логику.
func (uc *GetAllUseCase) Execute(ctx context.Context) (GetAllOutputDTO, error) {
	posts, err := uc.repo.FindAll(ctx)
	if err != nil {
		return GetAllOutputDTO{}, fmt.Errorf("Post.GetAllUseCase.Execute: %w", err)
	}

	var authorsIDS []vo.AuthorID
	for _, p := range posts {
		authorsIDS = append(authorsIDS, p.AuthorID())
	}

	authors, err := uc.authorRepo.FindByIDs(ctx, authorsIDS)
	if err != nil {
		return GetAllOutputDTO{}, fmt.Errorf("Post.GetAllUseCase.Execute: %w", err)
	}

	var authorsPosts []AuthorWithPostsDTO

	for _, a := range authors {
		for _, p := range posts {
			if p.AuthorID().Equal(a.ID()) {
				a.AddPost(p)
			}
		}

		authorPostDTO := MapAuthorWithPostsToDTO(a)
		authorsPosts = append(authorsPosts, authorPostDTO)
	}

	return GetAllOutputDTO{Posts: authorsPosts}, nil
}
