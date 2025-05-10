package post

import (
	"context"
	"fmt"
	dom "post-app/internal/domain/post"
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

// postRepository - интерфейс агрегатор
type postRepository interface {
	dom.Finder
	dom.Writer
}

// SaveUseCase бизнес логика Сохранения поста.
type SaveUseCase struct {
	readRepo  dom.Finder
	writeRepo dom.Writer
}

// NewSaveUseCase конструктор бизнес логики Сохранения поста.
func NewSaveUseCase(repo postRepository) *SaveUseCase {
	return &SaveUseCase{readRepo: repo, writeRepo: repo}
}

// Execute выполняет бизнес логику.
func (uc *SaveUseCase) Execute(ctx context.Context, in SaveInputDTO) error {
	postID, err := dom.NewPostID(in.ID)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	findedPost, err := uc.readRepo.FindByID(ctx, postID)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	postTitle, err := dom.NewPostTitle(in.Title)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	postContent, err := dom.NewPostContent(in.Content)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	post := dom.RehydratePost(postID, findedPost.AuthorID(), postTitle, postContent, findedPost.CreatedAt())

	err = uc.writeRepo.Save(ctx, post)
	if err != nil {
		return fmt.Errorf("Post.SaveUseCase.Execute: %w", err)
	}

	return nil
}
