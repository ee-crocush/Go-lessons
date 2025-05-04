package post

import (
	"context"
	"fmt"
	authordom "post-app/internal/domain/author"
	dom "post-app/internal/domain/post"
	"post-app/internal/domain/vo"
)

// SaveInputDTO входные данные для Сохранения поста.
type SaveInputDTO struct {
	AuthorID int32  `json:"author_id"`
	ID       int32  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// SaveContractUseCase определяет контракт для Сохранения поста.
type SaveContractUseCase interface {
	Execute(ctx context.Context, in SaveInputDTO) error
}

// SaveUseCase бизнес логика Сохранения поста.
type SaveUseCase struct {
	repo       dom.Writer
	authorRepo authordom.Finder
}

// NewSaveUseCase конструктор бизнес логики Сохранения поста.
func NewSaveUseCase(repo dom.Writer, authorRepo authordom.Finder) *SaveUseCase {
	return &SaveUseCase{repo: repo, authorRepo: authorRepo}
}

// Execute выполняет бизнес логику.
func (uc *SaveUseCase) Execute(ctx context.Context, in SaveInputDTO) error {
	authorID, err := vo.NewAuthorID(in.ID)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	author, err := uc.authorRepo.FindByID(ctx, authorID)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	postID, err := dom.NewPostID(in.ID)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	postTile, err := dom.NewPostTitle(in.Title)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	postContent, err := dom.NewPostContent(in.Content)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	post := dom.RehydratePost(postID, author.ID(), postTile, postContent, nil)

	err = uc.repo.Save(ctx, post)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	return nil
}
