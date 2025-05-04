package post

import (
	"encoding/json"
	"net/http"
	uc "post-app/internal/usecase/post"
)

// GetPostByAuthorIDRequest входные данные для получения постов по ID автора - запрос.
type GetPostByAuthorIDRequest struct {
	AuthorID int32 `json:"author_id"`
}

// GetPostByAuthorIDResponse выходные данные для получения постов по ID автора - ответ.
type GetPostByAuthorIDResponse struct {
	Data AuthorWithPostsDTO `json:"data"`
}

// GetPostsByAuthorIDHandler обработчик получения постов.
func (h *Handler) GetPostsByAuthorIDHandler(w http.ResponseWriter, r *http.Request) {
	var req GetPostByAuthorIDRequest

	// Читаем и парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	posts, err := h.getByAuthorID.Execute(r.Context(), uc.GetByAuthorIDInputDTO{AuthorID: req.AuthorID})
	if err != nil {
		http.Error(w, "failed to get posts by author id", http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := MapGetByAuthorIDUseCaseToRequest(posts)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
