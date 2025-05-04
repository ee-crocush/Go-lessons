// Package post содержит реализацию обработчиков запросов по постам.
package post

import (
	"context"
	uc "post-app/internal/usecase/post"
)

// CreateExecutor интерфейс для создания поста.
type CreateExecutor interface {
	Execute(ctx context.Context, in uc.CreateInputDTO) error
}

// GetByIDExecutor интерфейс для получения поста по ID.
type GetByIDExecutor interface {
	Execute(ctx context.Context, in uc.GetByIDInputDTO) (uc.GetByIDOutputDTO, error)
}

// GetByAuthorIDExecutor интерфейс для получения постов автора по ID автора.
type GetByAuthorIDExecutor interface {
	Execute(ctx context.Context, in uc.GetByAuthorIDInputDTO) (uc.GetByAuthorIDOutputDTO, error)
}

// GetAllExecutor интерфейс для получения всех постов.
type GetAllExecutor interface {
	Execute(ctx context.Context) (uc.GetAllOutputDTO, error)
}

// SaveExecutor интерфейс для сохранения поста.
type SaveExecutor interface {
	Execute(ctx context.Context, in uc.SaveInputDTO) error
}

// DeleteExecutor интерфейс для удаления поста.
type DeleteExecutor interface {
	Execute(ctx context.Context, in uc.DeleteInputDTO) error
}

// Handler структура для обработки запросов.
type Handler struct {
	createUC      CreateExecutor
	getByIDUC     GetByIDExecutor
	getByAuthorID GetByAuthorIDExecutor
	getAll        GetAllExecutor
	saveUC        SaveExecutor
	deleteUC      DeleteExecutor
}

// NewHandler конструктор для создания нового экземпляра Handler.
func NewHandler(
	createUC CreateExecutor, getByIDUC GetByIDExecutor, getByAuthorID GetByAuthorIDExecutor, getAll GetAllExecutor,
	saveUC SaveExecutor, deleteUC DeleteExecutor,
) *Handler {
	return &Handler{
		createUC:      createUC,
		getByIDUC:     getByIDUC,
		getByAuthorID: getByAuthorID,
		getAll:        getAll,
		saveUC:        saveUC,
		deleteUC:      deleteUC,
	}
}
