// Package health содержит реализацию эндпоинт для проверки состояния сервиса.
package health

import (
	"encoding/json"
	"net/http"
)

// Handler отвечает за healthcheck-эндпоинт.
type Handler struct{}

// NewHandler возвращает новый экземпляр Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// HealthCheckHandler обрабатывает запросы для проверки состояния сервиса.
func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status":  "ok",
		"message": "Service is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to send JSON response", http.StatusInternalServerError)
	}
}
