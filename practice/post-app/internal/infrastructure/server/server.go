// Package server запускает сервер.
package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"post-app/internal/infrastructure/config"
	"post-app/internal/infrastructure/server/mux"
	author_handler "post-app/internal/infrastructure/transport/httplib/handler/author"
	post_handler "post-app/internal/infrastructure/transport/httplib/handler/post"
	"sync"
	"syscall"
	"time"
)

// GracefulServer — интерфейс для серверов, которые поддерживают
// блокирующий запуск и корректное завершение (shutdown).
type GracefulServer interface {
	Start() error
	Shutdown(ctx context.Context) error
}

// CreateServers создаёт список серверов на основе конфигурации.
func CreateServers(cfg config.Config, ah *author_handler.Handler, ph *post_handler.Handler) ([]GracefulServer, error) {
	var servers []GracefulServer

	if cfg.HTTP.Enabled {
		servers = append(servers, mux.NewMuxServer(cfg, ah, ph))
	}

	if len(servers) == 0 {
		return nil, fmt.Errorf("no servers are enabled in the configuration")
	}

	return servers, nil
}

// StartAll запускает все сервера, слушает ошибки и сигналы, делает graceful shutdown
func StartAll(servers ...GracefulServer) error {
	errChan := make(chan error, len(servers))
	var wg sync.WaitGroup

	for _, srv := range servers {
		wg.Add(1)
		go func(s GracefulServer) {
			defer wg.Done()
			if err := s.Start(); err != nil {
				errChan <- err
			}
		}(srv)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		shutdownAll(servers)
		wg.Wait()
		return err
	case <-sigChan:
		shutdownAll(servers)
		wg.Wait()
		return nil
	}
}

// shutdownAll вызывает Shutdown у всех серверов с общим таймаутом
func shutdownAll(servers []GracefulServer) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(servers))

	for _, srv := range servers {
		go func(s GracefulServer) {
			defer wg.Done()
			if err := s.Shutdown(ctx); err != nil {
				fmt.Printf("Error shutting down server: %v\n", err)
			}
		}(srv)
	}

	wg.Wait()
}
