package post

import (
	"encoding/json"
	"net/http"
	uc "post-app/internal/usecase/post"
)

// DeletePostRequest входные данные для удаления поста из запроса.
type DeletePostRequest struct {
	ID int32 `json:"id"`
}

// DeletePostResponse выходные данные для удаления поста - ответ.
type DeletePostResponse struct {
	Message string `json:"message"`
}

// DeletePostHandler обработчик удаления поста.
func (h *Handler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var req DeletePostRequest

	// Читаем и парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	if err := h.deleteUC.Execute(r.Context(), uc.DeleteInputDTO{ID: req.ID}); err != nil {
		http.Error(w, "failed to create post", http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := DeletePostResponse{
		Message: "Пост успешно удален!",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
