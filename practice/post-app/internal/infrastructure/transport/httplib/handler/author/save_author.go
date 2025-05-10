package author

import (
	"encoding/json"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/parser"
	uc "post-app/internal/usecase/author"
)

// SaveAuthorRequest входные данные для сохранения автора из запроса.
type SaveAuthorRequest struct {
	Name string `json:"name"`
}

// SaveAuthorDTO данные, которые вернем в ответе.
type SaveAuthorDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// SaveAuthorResponse выходные данные для сохранения автора - ответ.
type SaveAuthorResponse struct {
	Message string `json:"message"`
}

// SaveAuthorHandler обработчик сохранения автора.
func (h *Handler) SaveAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var req SaveAuthorRequest

	id, err := parser.ExtractIDFromRequest(r, "authorID")
	if err != nil {
		http.Error(w, "incorrect author ID", http.StatusBadRequest)
		return
	}

	// Читаем и парсим JSON
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "wrong request body", http.StatusBadRequest)
		return
	}

	// Выполняем use-case с извлеченным ID и именем из запроса
	if err = h.saveUC.Execute(r.Context(), uc.SaveInputDTO{ID: id, Name: req.Name}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := SaveAuthorResponse{Message: "Автор успешно сохранен!"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
