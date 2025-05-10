package post

import (
	"encoding/json"
	"net/http"
)

// GetAllPostsResponse выходные данные для получения постов - ответ.
type GetAllPostsResponse struct {
	Data []AuthorWithPostsDTO `json:"data"`
}

// GetAllPostsHandler обработчик получения постов.
func (h *Handler) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := h.getAll.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := MapGetAllUseCaseToRequest(posts)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
