package main

import (
	"context"
	"encoding/json"
	"github.com/ee-crocush/task-tracker/internal/infrastructure/config"
	"github.com/ee-crocush/task-tracker/internal/infrastructure/repository/postgres"
	uctask "github.com/ee-crocush/task-tracker/internal/usecase/task"
	"github.com/rs/zerolog/log"
)

func main() {
	configPath := "./configs/config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("File config not found")
	}

	dbPool, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")
	}
	defer dbPool.Close()

	log.Info().
		Str("host", cfg.DB.Host).
		Int("port", cfg.DB.Port).
		Str("database", cfg.DB.Name).
		Msg("Database connected successfully!")

	// Типа имитация ручек http

	// Создаем пару меток
	label1, err := CreateLabel(context.Background(), dbPool)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create label")
	}
	log.Info().Int("labelID", label1.ID).Msg("Label created successfully!")

	label2, err := CreateLabel(context.Background(), dbPool)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create label")
	}
	log.Info().Int("labelID", label2.ID).Msg("Label created successfully!")

	// Создаем нового пользователя автора задачи и получаем его ID.
	authorID, err := CreateUser(context.Background(), dbPool)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create user")
	}
	log.Info().Int("authorID", authorID).Msg("User created successfully!")

	// Создаем нового пользователя проверяющего задачу и получаем его ID.
	assignedID, err := CreateUser(context.Background(), dbPool)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create user")
	}
	log.Info().Int("assignedID", assignedID).Msg("User created successfully!")

	// Создадим несколько задач
	var labels1 = []uctask.LabelDTO{
		{ID: label1.ID, Name: label1.Name},
		{ID: label2.ID, Name: label2.Name},
	}
	taskID1, err := CreateTask(context.Background(), dbPool, authorID, assignedID, labels1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create task")
	}
	log.Info().Int("taskID", taskID1).Msg("Task created successfully!")

	var labels2 = []uctask.LabelDTO{
		{ID: label1.ID, Name: label1.Name},
	}
	taskID2, err := CreateTask(context.Background(), dbPool, authorID, authorID, labels2)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create task")
	}
	log.Info().Int("taskID", taskID2).Msg("Task created successfully!")

	var labels3 = []uctask.LabelDTO{
		{ID: label2.ID, Name: label2.Name},
	}
	taskID3, err := CreateTask(context.Background(), dbPool, assignedID, assignedID, labels3)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create task")
	}
	log.Info().Int("taskID", taskID3).Msg("Task created successfully!")

	// Получим все задачи
	tasksAll, err := GetAllTasks(context.Background(), dbPool)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get all tasks")
	}

	jsonBytes, err := json.Marshal(tasksAll)
	if err != nil {
		log.Error().Err(err).Msg("Error marshaling JSON")
		return
	}

	log.Info().RawJSON("output", jsonBytes).Msg("Tasks retrieved all successfully!")

	// Получим все задачи автора authorID
	tasks, err := GetTasksByAuthor(context.Background(), dbPool, authorID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get tasks by author")
	}

	jsonBytes, err = json.Marshal(tasks)
	if err != nil {
		log.Error().Err(err).Msg("Error marshaling JSON")
		return
	}

	log.Info().RawJSON("output", jsonBytes).Msg("Tasks retrieved by author successfully!")
}
