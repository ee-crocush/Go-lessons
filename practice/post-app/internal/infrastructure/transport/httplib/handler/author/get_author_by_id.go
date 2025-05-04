package author

import (
	"encoding/json"
	"net/http"
	uc "post-app/internal/usecase/author"
)

// GetAuthorRequest входные данные для получения автора из запроса.
type GetAuthorRequest struct {
	ID int32 `json:"id"`
}

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
	var req GetAuthorRequest

	// Читаем и парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	output, err := h.getUC.Execute(
		r.Context(), uc.GetInputDTO{ID: req.ID},
	)
	if err != nil {
		http.Error(w, "failed to get author", http.StatusBadRequest)
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
