package post

import (
	"encoding/json"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/parser"
	uc "post-app/internal/usecase/post"
)

// DeletePostResponse выходные данные для удаления поста - ответ.
type DeletePostResponse struct {
	Message string `json:"message"`
}

// DeletePostHandler обработчик удаления поста.
func (h *Handler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ExtractIDFromRequest(r, "postID")
	if err != nil {
		http.Error(w, "incorrect post ID", http.StatusBadRequest)
		return
	}

	if err = h.deleteUC.Execute(r.Context(), uc.DeleteInputDTO{ID: id}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := DeletePostResponse{
		Message: "Пост успешно удален!",
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
