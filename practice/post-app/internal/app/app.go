package app

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"post-app/internal/infrastructure/config"
	"post-app/internal/infrastructure/server"
	author_handler "post-app/internal/infrastructure/transport/httplib/handler/author"
	post_handler "post-app/internal/infrastructure/transport/httplib/handler/post"
	author_uc "post-app/internal/usecase/author"
	post_uc "post-app/internal/usecase/post"
)

// Run запускает HTTP сервер и инициализирует все необходимые компоненты.
func Run(log *zerolog.Logger) error {
	configPath := "./configs/config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("File config not found")
	}

	// Создаем репозитории
	authorRepo, postRepo, mongoClient, pgPool, err := createRepository(cfg, log)

	// Если работали с монгой - закрываем
	if mongoClient != nil {
		defer func() {
			if err = mongoClient.Disconnect(context.Background()); err != nil {
				log.Error().Err(err).Msg("Failed to disconnect MongoDB")
			} else {
				log.Info().Msg("MongoDB disconnected successfully")
			}
		}()
	}

	// Если с postgres
	if pgPool != nil {
		defer pgPool.Close()
	}

	// UseCase для авторов
	authorCreateUC := author_uc.NewCreateUseCase(authorRepo)
	authorGetUC := author_uc.NewGetUseCase(authorRepo)
	authorSaveUC := author_uc.NewSaveUseCase(authorRepo)

	// UseCase для постов
	postCreateUC := post_uc.NewCreateUseCase(postRepo, authorRepo)
	postGetByIDUC := post_uc.NewGetByIDUseCase(postRepo, authorRepo)
	postGetByAuthorUC := post_uc.NewGetByAuthorIDUseCase(postRepo, authorRepo)
	postGetAllUC := post_uc.NewGetAllUseCase(postRepo, authorRepo)
	postSaveUC := post_uc.NewSaveUseCase(postRepo)
	postDeleteUC := post_uc.NewDeleteUseCase(postRepo)

	// Создаем хендлеры
	authorHandler := author_handler.NewHandler(authorCreateUC, authorGetUC, authorSaveUC)
	postHandler := post_handler.NewHandler(
		postCreateUC, postGetByIDUC, postGetByAuthorUC, postGetAllUC, postSaveUC, postDeleteUC,
	)

	servers, err := server.CreateServers(cfg, authorHandler, postHandler)
	if err != nil {
		return fmt.Errorf("failed to create servers: %w", err)
	}

	return server.StartAll(servers...)
}
