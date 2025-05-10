package author

import (
	"encoding/json"
	"net/http"
	uc "post-app/internal/usecase/author"
)

// CreateAuthorRequest входные данные для создания автора из запроса.
type CreateAuthorRequest struct {
	Name string `json:"name"`
}

// CreateAuthorDTO данные, которые вернем в ответе.
type CreateAuthorDTO struct {
	ID int32 `json:"id"`
}

// CreateAuthorResponse выходные данные для создания автора - ответ.
type CreateAuthorResponse struct {
	Data    CreateAuthorDTO `json:"data"`
	Message string          `json:"message"`
}

// CreateAuthorHandler обработчик создания автора.
func (h *Handler) CreateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateAuthorRequest

	// Читаем и парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	output, err := h.createUC.Execute(
		r.Context(), uc.CreateInputDTO{Name: req.Name},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := CreateAuthorResponse{
		Data:    CreateAuthorDTO{ID: output.ID},
		Message: "Автор успешно создан!",
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
