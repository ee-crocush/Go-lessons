package mux

import (
	"context"
	"fmt"
	"net/http"
	"post-app/internal/infrastructure/config"
	"post-app/internal/infrastructure/transport/httplib"
	author_handler "post-app/internal/infrastructure/transport/httplib/handler/author"
	post_handler "post-app/internal/infrastructure/transport/httplib/handler/post"
	"post-app/pkg/logger"
	"time"
)

// MuxServer Mux сервер
type MuxServer struct {
	cfg        config.Config
	httpServer *http.Server
}

func NewMuxServer(
	cfg config.Config, ah *author_handler.Handler, ph *post_handler.Handler,
) *MuxServer {
	router := httplib.SetupRoutes(ah, ph)

	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.App.WriteTimeout) * time.Second,
	}

	return &MuxServer{
		cfg:        cfg,
		httpServer: httpServer,
	}
}

func (s *MuxServer) Start() error {
	log := logger.GetLogger()
	log.Info().
		Str("host", s.httpServer.Addr).
		Msg("Mux HTTP server started")

	return s.httpServer.ListenAndServe()
}

func (s *MuxServer) Shutdown(ctx context.Context) error {
	log := logger.GetLogger()
	log.Info().Msg("Shutting down Mux HTTP server...")

	return s.httpServer.Shutdown(ctx)
}
