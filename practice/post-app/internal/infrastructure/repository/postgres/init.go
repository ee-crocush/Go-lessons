// Package postgres содержит реализацию репозиториев для работы с PostgreSQL.
package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"post-app/internal/infrastructure/config"
)

// Init инициализирует и возвращает новый экземпляр пула соединений PostgreSQL.
func Init(cfg config.Config) (*pgxpool.Pool, error) {
	dbURL := cfg.DB.DSN()

	// Добавляем параметры подключения
	queryParams := dbURL.Query()
	queryParams.Set("sslmode", cfg.DB.SSLMode)
	queryParams.Set("pool_max_conns", cfg.DB.PoolMaxConns)
	queryParams.Set("pool_min_conns", cfg.DB.PoolMinConns)
	queryParams.Set("pool_max_conn_lifetime", cfg.DB.PoolMaxConnLifetime)
	queryParams.Set("pool_max_conn_idle_time", cfg.DB.PoolMaxConnIdletime)
	queryParams.Set("connect_timeout", fmt.Sprintf("%.0f", cfg.DB.ConnectTimeout.Seconds()))
	dbURL.RawQuery = queryParams.Encode()

	// Создаём пул соединений
	pool, err := NewPGXPool(dbURL.String())
	if err != nil {
		return nil, fmt.Errorf("Init.Postgres.NewPGXPool: %w", err)
	}

	// Проверяем соединение
	ctx, cancel := context.WithTimeout(context.Background(), cfg.DB.ConnectTimeout)
	defer cancel()

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Init.Postgres.Ping: %w", err)
	}

	return pool, nil
}

// NewPGXPool создаёт пул соединений к PostgreSQL с помощью pgxpool.
func NewPGXPool(dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect: %w", err)
	}

	return pool, nil
}
