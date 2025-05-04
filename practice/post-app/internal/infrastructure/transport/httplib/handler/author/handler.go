// Package author содержит реализацию обработчиков запросов по авторам.
package author

import (
	"context"
	uc "post-app/internal/usecase/author"
)

// CreateExecutor интерфейс для создания автора.
type CreateExecutor interface {
	Execute(ctx context.Context, in uc.CreateInputDTO) (uc.CreateOutputDTO, error)
}

// GetExecutor интерфейс для получения автора.
type GetExecutor interface {
	Execute(ctx context.Context, in uc.GetInputDTO) (uc.GetOutputDTO, error)
}

// SaveExecutor интерфейс для сохранения автора.
type SaveExecutor interface {
	Execute(ctx context.Context, in uc.SaveInputDTO) error
}

// Handler структура для обработки запросов.
type Handler struct {
	createUC CreateExecutor
	getUC    GetExecutor
	saveUC   SaveExecutor
}

// NewHandler конструктор для создания нового экземпляра Handler.
func NewHandler(createUC CreateExecutor, getUC GetExecutor, saveUC SaveExecutor) *Handler {
	return &Handler{
		createUC: createUC,
		getUC:    getUC,
		saveUC:   saveUC,
	}
}
