package post

import (
	"encoding/json"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/parser"
	uc "post-app/internal/usecase/post"
)

// GetPostByIDResponse выходные данные для получения поста по его ID - ответ.
type GetPostByIDResponse struct {
	Author AuthorDTO `json:"author"`
	Post   PostDTO   `json:"post"`
}

// GetPostByIDHandler обработчик получения поста по его ID.
func (h *Handler) GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ExtractIDFromRequest(r, "postID")
	if err != nil {
		http.Error(w, "incorrect post ID", http.StatusBadRequest)
		return
	}

	output, err := h.getByIDUC.Execute(r.Context(), uc.GetByIDInputDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := MapGetByIDUseCaseToRequest(output)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
