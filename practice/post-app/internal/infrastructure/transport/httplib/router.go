// Package httplib содержит все эндпоинты.
package httplib

import (
	"github.com/gorilla/mux"
	"net/http"
	"post-app/internal/infrastructure/transport/httplib/handler/author"
	"post-app/internal/infrastructure/transport/httplib/handler/health"
	"post-app/internal/infrastructure/transport/httplib/handler/post"
	"post-app/pkg/logger"
	"post-app/pkg/middleware"
)

// SetupRoutes возвращает роутер с обработчиками.
func SetupRoutes(authorHandler *author.Handler, postHandler *post.Handler) http.Handler {
	r := mux.NewRouter()

	log := logger.GetLogger()
	// Подключаем middleware
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.LoggingMiddleware(log))
	r.Use(middleware.ErrorHandlerMiddleware(log))

	// Healthcheck
	healthHandler := health.NewHandler()
	r.HandleFunc("/health", healthHandler.HealthCheckHandler).Methods(http.MethodGet)
	// Авторы
	r.HandleFunc("/authors", authorHandler.CreateAuthorHandler).Methods(http.MethodPost)
	r.HandleFunc("/authors/{id}", authorHandler.GetAuthorHandler).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id}", authorHandler.SaveAuthorHandler).Methods(http.MethodPut)

	// Посты
	r.HandleFunc("/posts", postHandler.CreatePostHandler).Methods(http.MethodPost)
	r.HandleFunc("/posts", postHandler.GetAllPostsHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", postHandler.GetPostByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts/by_author/{id}", postHandler.GetPostsByAuthorIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", postHandler.SavePostHandler).Methods(http.MethodPut)
	r.HandleFunc("/posts/{id}", postHandler.DeletePostHandler).Methods(http.MethodDelete)

	return r
}
