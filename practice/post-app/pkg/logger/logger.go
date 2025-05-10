package logger

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

var log zerolog.Logger
var once sync.Once

// InitLogger инициализирует глобальный логгер.
func InitLogger(serviceName string) {
	once.Do(
		func() {
			var output zerolog.ConsoleWriter

			output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
			zerolog.SetGlobalLevel(zerolog.DebugLevel)

			log = zerolog.New(output).
				With().
				Timestamp().
				Str("service", serviceName).
				Logger()
		},
	)
}

// GetLogger возвращает глобальный логгер.
func GetLogger() *zerolog.Logger {
	if log.GetLevel() == zerolog.NoLevel {
		panic("Logger is not initialized. Call InitLogger first.")
	}

	return &log
}

func WithContext(ctx context.Context) zerolog.Logger {
	logCtx := log.With()

	if requestID, ok := ctx.Value("request_id").(string); ok && requestID != "" {
		logCtx = logCtx.Str("request_id", requestID)
	}
	if correlationID, ok := ctx.Value("correlation_id").(string); ok && correlationID != "" {
		logCtx = logCtx.Str("correlation_id", correlationID)
	}
	if userID, ok := ctx.Value("user_id").(string); ok && userID != "" {
		logCtx = logCtx.Str("user_id", userID)
	}

	return logCtx.Logger()
}

func GetLoggerWithContext(ctx context.Context) *zerolog.Logger {
	l := WithContext(ctx)
	newLogger := l.With().Logger()

	return &newLogger
}
