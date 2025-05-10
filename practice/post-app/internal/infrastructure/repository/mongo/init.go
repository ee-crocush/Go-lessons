// Package mongo содержит реализацию репозиториев для работы с MongoDB.
package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"post-app/internal/infrastructure/config"
)

// Init инициализирует и возвращает клиент MongoDB и выбранную базу.
func Init(cfg config.Config) (*mongo.Client, *mongo.Database, error) {
	uri := cfg.MongoDB.URI().String()

	clientOpts := options.Client().
		ApplyURI(uri).
		SetAuth(
			options.Credential{
				Username:   cfg.MongoDB.User,
				Password:   cfg.MongoDB.Password,
				AuthSource: cfg.MongoDB.AuthSource,
			},
		).
		SetConnectTimeout(cfg.MongoDB.ConnectTimeout)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoDB.ConnectTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("Init.MongoDB.Connect: %w", err)
	}

	//defer client.Disconnect(context.Background())

	// Проверка соединения
	if err = client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("Init.MongoDB.Ping: %w", err)
	}

	db := client.Database(cfg.MongoDB.Database)

	return client, db, nil
}
