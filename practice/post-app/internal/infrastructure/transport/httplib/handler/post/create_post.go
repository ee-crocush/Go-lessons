package post

import (
	"encoding/json"
	"net/http"
	uc "post-app/internal/usecase/post"
)

// CreatePostRequest входные данные для создания поста из запроса.
type CreatePostRequest struct {
	AuthorID int32  `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// CreatePostResponse выходные данные для создания поста - ответ.
type CreatePostResponse struct {
	Message string `json:"message"`
}

// CreatePostHandler обработчик создания поста.
func (h *Handler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest

	// Читаем и парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	in := uc.CreateInputDTO{
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Content:  req.Content,
	}
	if err := h.createUC.Execute(r.Context(), in); err != nil {
		http.Error(w, "failed to create post", http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := CreatePostResponse{
		Message: "Пост успешно создан!",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
