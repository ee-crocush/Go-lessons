package post

import (
	"encoding/json"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/parser"
	uc "post-app/internal/usecase/post"
)

// GetPostByAuthorIDResponse выходные данные для получения постов по ID автора - ответ.
type GetPostByAuthorIDResponse struct {
	Data AuthorWithPostsDTO `json:"data"`
}

// GetPostsByAuthorIDHandler обработчик получения постов.
func (h *Handler) GetPostsByAuthorIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ExtractIDFromRequest(r, "authorID")
	if err != nil {
		http.Error(w, "incorrect authorID", http.StatusBadRequest)
		return
	}

	posts, err := h.getByAuthorID.Execute(r.Context(), uc.GetByAuthorIDInputDTO{AuthorID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
