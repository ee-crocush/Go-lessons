package app

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	author_dom "post-app/internal/domain/author"
	post_dom "post-app/internal/domain/post"
	"post-app/internal/infrastructure/config"
	mem_repo "post-app/internal/infrastructure/repository/inmemory"
	mongo_repo "post-app/internal/infrastructure/repository/mongo"
	pg_repo "post-app/internal/infrastructure/repository/postgres"
)

// createRepository создает репозитории в зависимости от типа базы данных, указанного в конфигурации.
// Возвращает интерфейсы репозиториев для авторов и постов, а также ошибку, если таковая произошла.
func createRepository(cfg config.Config, log *zerolog.Logger) (
	author_dom.Repository, post_dom.Repository, *mongo.Client, *pgxpool.Pool, error,
) {
	switch cfg.App.DBType {
	case "postgres":
		aRepo, pRepo, pool, err := createPGRepository(cfg, log)
		return aRepo, pRepo, nil, pool, err
	case "mongo":
		aRepo, pRepo, client, err := createMongoRepository(cfg, log)
		return aRepo, pRepo, client, nil, err
	case "in-memory":
		aRepo, pRepo, err := createInMemoryRepository()
		return aRepo, pRepo, nil, nil, err
	default:
		err := fmt.Errorf("unsupported DB type: %s", cfg.App.DBType)
		log.Error().Err(err).Msg("Invalid database type")

		return nil, nil, nil, nil, err
	}
}

// createPGRepository Создает репозиторий Postgres
func createPGRepository(cfg config.Config, log *zerolog.Logger) (
	*pg_repo.AuthorRepository, *pg_repo.PostRepository, *pgxpool.Pool, error,
) {
	pool, err := pg_repo.Init(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")

		return nil, nil, nil, err
	}
	//defer pool.Close()

	log.Info().
		Str("host", cfg.DB.Host).
		Int("port", cfg.DB.Port).
		Str("database", cfg.DB.Name).
		Msg("Database connected successfully!")

	authorRepo := pg_repo.NewAuthorRepository(pool)
	postRepo := pg_repo.NewPostRepository(pool)

	return authorRepo, postRepo, pool, nil
}

// createMongoRepository создает репозитории Mongo
func createMongoRepository(cfg config.Config, log *zerolog.Logger) (
	*mongo_repo.AuthorRepository, *mongo_repo.PostRepository, *mongo.Client, error,
) {
	client, db, err := mongo_repo.Init(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("MongoDB connection failed")

		return nil, nil, nil, err
	}
	//defer client.Disconnect(context.Background())

	log.Info().
		Str("host", cfg.MongoDB.Host).
		Int("port", cfg.MongoDB.Port).
		Str("database", cfg.MongoDB.Database).
		Msg("MongoDB  connected successfully!")

	authorRepo := mongo_repo.NewAuthorRepository(db, cfg.MongoDB.ConnectTimeout)
	postRepo := mongo_repo.NewPostRepository(db, cfg.MongoDB.ConnectTimeout)

	return authorRepo, postRepo, client, nil
}

// createInMemoryRepository создает репозитории InMemory
func createInMemoryRepository() (*mem_repo.AuthorRepository, *mem_repo.PostRepository, error) {
	authorRepo := mem_repo.NewAuthorRepository()
	postRepo := mem_repo.NewPostRepository()

	return authorRepo, postRepo, nil
}
