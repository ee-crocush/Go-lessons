// Package postgres содержит реализацию репозиториев для работы с PostgreSQL.
package postgres

import (
	"context"
	"fmt"
	"github.com/ee-crocush/task-tracker/internal/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// InitDB инициализирует и возвращает новый экземпляр пула соединений PostgreSQL.
func InitDB(config *config.Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.DB.DSN())
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect: %w", err)
	}
	return pool, nil
}
