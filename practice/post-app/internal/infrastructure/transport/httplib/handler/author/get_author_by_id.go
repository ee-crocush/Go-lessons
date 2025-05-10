package author

import (
	"encoding/json"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/parser"
	uc "post-app/internal/usecase/author"
)

// GetAuthorDTO данные, которые вернем в ответе.
type GetAuthorDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// GetAuthorResponse выходные данные для получения автора - ответ.
type GetAuthorResponse struct {
	Data GetAuthorDTO `json:"data"`
}

// GetAuthorHandler обработчик получения автора.
func (h *Handler) GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ExtractIDFromRequest(r, "authorID")
	if err != nil {
		http.Error(w, "incorrect author ID", http.StatusBadRequest)
		return
	}

	output, err := h.getUC.Execute(
		r.Context(), uc.GetInputDTO{ID: id},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := GetAuthorResponse{
		Data: GetAuthorDTO{ID: output.ID, Name: output.Name},
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
