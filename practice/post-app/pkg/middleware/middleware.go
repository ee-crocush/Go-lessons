package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"runtime/debug"
	"time"
)

type contextKey string

const requestIDKey contextKey = "request_id"

// RequestIDMiddleware генерирует и добавляет request_id в контекст.
func RequestIDMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				requestID := r.Header.Get("X-Request-ID")
				if requestID == "" {
					requestID = uuid.New().String()
				}

				// Добавляем request_id в контекст
				ctx := context.WithValue(r.Context(), requestIDKey, requestID)

				// Добавляем request_id в заголовок ответа
				w.Header().Set("X-Request-ID", requestID)

				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}

// responseWriter оборачивает http.ResponseWriter для захвата статуса ответа.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware логгирует запросы и их выполнение.
func LoggingMiddleware(log *zerolog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()

				// Получаем request_id из контекста
				requestID, _ := r.Context().Value(requestIDKey).(string)

				// Логируем начало обработки
				log.Info().
					Str("request_id", requestID).
					Str("method", r.Method).
					Str("url", r.URL.Path).
					Msg("Request started")

				// Обертка для записи статуса ответа
				ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

				// Выполнение запроса
				next.ServeHTTP(ww, r)

				// Логируем завершение запроса
				duration := time.Since(start)

				// Если статус >= 400 — логируем как ошибку
				event := log.Info()
				if ww.statusCode >= 400 {
					event = log.Error()
					// Извлекаем сообщение об ошибке из заголовка ответа
					if errMsg := ww.Header().Get("X-Error-Message"); errMsg != "" {
						event.Str("error", errMsg)
					}
				}

				event.
					Str("request_id", requestID).
					Str("method", r.Method).
					Str("url", r.URL.Path).
					Int("status_code", ww.statusCode).
					Str("duration", duration.String()).
					Msg("Request completed")
			},
		)
	}
}

// ErrorHandlerMiddleware обрабатывает панику и логгирует её.
func ErrorHandlerMiddleware(log *zerolog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				defer func() {
					if rec := recover(); rec != nil {
						requestID, _ := r.Context().Value(requestIDKey).(string)

						log.Error().
							Str("request_id", requestID).
							Str("method", r.Method).
							Str("url", r.URL.Path).
							Str("stack_trace", fmt.Sprintf("%+v", rec)).
							Bytes("stack", debug.Stack()).
							Msg("Recovered from panic")

						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					}
				}()
				next.ServeHTTP(w, r)
			},
		)
	}
}
