package post

import (
	"encoding/json"
	"net/http"
	uc "post-app/internal/usecase/post"
)

// GetPostByIDRequest входные данные для получения поста по его ID из запроса.
type GetPostByIDRequest struct {
	ID int32 `json:"id"`
}

// GetPostByIDResponse выходные данные для получения поста по его ID - ответ.
type GetPostByIDResponse struct {
	Author AuthorDTO `json:"author"`
	Post   PostDTO   `json:"post"`
}

// GetPostByIDHandler обработчик получения поста по его ID.
func (h *Handler) GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	var req GetPostByIDRequest

	// Читаем и парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	post, err := h.getByIDUC.Execute(r.Context(), uc.GetByIDInputDTO{ID: req.ID})
	if err != nil {
		http.Error(w, "failed to get post", http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := MapGetByIDUseCaseToRequest(post)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
