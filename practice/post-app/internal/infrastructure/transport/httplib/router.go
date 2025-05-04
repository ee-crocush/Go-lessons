// Package httplib содержит все эндпоинты.
package httplib

import (
	"github.com/gorilla/mux"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/handler/author"
	"post-app/internal/infrastructure/transport/httplib/handler/health"
)

// NewRouter возвращает роутер с обработчиками.
func NewRouter(authorHandler *author.Handler) http.Handler {
	r := mux.NewRouter()

	// Healthcheck
	healthHandler := health.NewHandler()
	r.HandleFunc("/health", healthHandler.HealthCheckHandler).Methods(http.MethodGet)
	// Авторы
	r.HandleFunc("/authors", authorHandler.CreateAuthorHandler).Methods(http.MethodPost)
	r.HandleFunc("/authors/{id}", authorHandler.GetAuthorHandler).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id}", authorHandler.SaveAuthorHandler).Methods(http.MethodPut)

	return r
}
