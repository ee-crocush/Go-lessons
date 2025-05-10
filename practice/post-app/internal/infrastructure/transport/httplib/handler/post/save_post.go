package post

import (
	"encoding/json"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/parser"
	uc "post-app/internal/usecase/post"
)

// SavePostRequest входные данные для сохранения поста из запроса.
type SavePostRequest struct {
	AuthorID int32  `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// SavePostResponse выходные данные для сохранения поста - ответ.
type SavePostResponse struct {
	Message string `json:"message"`
}

// SavePostHandler обработчик сохранения поста.
func (h *Handler) SavePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ExtractIDFromRequest(r, "postID")
	if err != nil {
		http.Error(w, "incorrect post ID", http.StatusBadRequest)
		return
	}

	var req SavePostRequest

	// Читаем и парсим JSON
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	in := uc.SaveInputDTO{
		AuthorID: req.AuthorID,
		ID:       id,
		Title:    req.Title,
		Content:  req.Content,
	}
	if err = h.saveUC.Execute(r.Context(), in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := SavePostResponse{Message: "Пост успешно сохранен!"}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
